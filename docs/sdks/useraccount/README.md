# UserAccount

## Overview

User authentication including multi-step MFA, password reset, OTP login, and token management

### Available Operations

* [InitAuth](#initauth) - Initialize authentication session
* [Authenticate](#authenticate) - Authenticate user with credentials
* [RefreshToken](#refreshtoken) - Refresh access token
* [GetPasswordStatus](#getpasswordstatus) - Get password status
* [SendPasswordSetupEmail](#sendpasswordsetupemail) - Email a link to set a first password
* [ResetPassword](#resetpassword) - Reset password

## InitAuth

Start a server-side authentication session and discover which sign-in methods are
configured for the organization. This is the first step in the multi-step login flow.

**Request body (optional)**

- You may omit the body, send an empty JSON object `{}`, or send `{ "email": "..." }`.
- `email` in the body is optional and kept for legacy reasons; omitting it does not prevent
  initialization. The web client typically calls this endpoint without a body and sends
  `email` on `/authenticate` instead.
- When provided, `email` is stored on the session for correlation with subsequent steps.

**Flow:**

1. Call this endpoint (optional JSON body as above).
2. Receive a session token in the `x-session-token` response header.
3. Send that token on subsequent `/authenticate` requests (`x-session-token` header).
4. Use `allowedMethods` and `authProviders` from the response to render the login UI.

**Session token**

- Returned as header `x-session-token`.
- Required for `/authenticate` (and related steps) until it expires.

**Multi-factor authentication**

If the organization has MFA, complete multiple authentication steps; each step may
return the next step's allowed methods.


### Example Usage

<!-- UsageSnippet language="go" operationID="initAuth" method="post" path="/userAccount/initAuth" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.UserAccount.InitAuth(ctx, &components.InitAuthRequest{
        Email: pipeshub.Pointer("user@example.com"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InitAuthResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.InitAuthRequest](../../models/components/initauthrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.InitAuthResponse](../../models/operations/initauthresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.ErrorResponse          | 400                              | application/json                 |
| apierrors.RateLimitErrorResponse | 429                              | application/json                 |
| apierrors.ErrorResponse          | 500                              | application/json                 |
| apierrors.APIError               | 4XX, 5XX                         | \*/\*                            |

## Authenticate

Authenticate a user using the specified method and credentials.
Requires a valid session token from `/initAuth`.

**Credential Formats by Method:**

- `password`: `{ "credentials": { "password": "your-password" } }`
- `otp`: `{ "credentials": { "otp": "123456" } }` (6-digit code, valid for 10 minutes)
- `google`: `{ "credentials": "google-id-token-string" }`. The object form
  `{ "credentials": { "credential": "..." } }` passes the first token check, but by code
  reading the second check is handed the whole object, so it ends in `500`
- `microsoft`: `{ "credentials": { "idToken": "..." } }` (only `idToken` is read)

SAML (`samlSso`), generic OAuth (`oauth`) and Azure AD (`azureAd`) sign-in were removed in 1.0;
this endpoint answers `400` for them. Azure AD accounts sign in with `microsoft`.

**Multi-Step Response:**

If organization uses MFA, successful authentication returns:
- `status: "success"` with `nextStep` and `allowedMethods` for next step

**Fully Authenticated Response:**

After completing all steps:
- `message: "Fully authenticated"` with `accessToken` (1hr) and `refreshToken` (7d)

**Security:**

- Account locks for 24 hours after 10 consecutive failed attempts, and the owner is
  sent an email saying so. While it is locked, sign-in is refused with the same answer
  as a wrong password or code, even when the password or code is right
- CAPTCHA may be required if enabled (pass `cf-turnstile-response`)
- An email with no account gets the same status and message as a real account given
  a wrong password (`400`) or a wrong, missing or expired sign-in code (`401`)
- The session token is single-use for the final step: once tokens are returned the
  session is gone, and sending it again answers `401`

**Order of checks:** the session token is checked before the body, so a request
without a valid `x-session-token` answers `401` whatever its body holds.


### Example Usage

<!-- UsageSnippet language="go" operationID="authenticate" method="post" path="/userAccount/authenticate" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.UserAccount.Authenticate(ctx, "<value>", components.AuthenticateRequest{
        Method: components.MethodMicrosoft,
        Credentials: components.CreateCredentialsPasswordCredentials(
            components.PasswordCredentials{
                Password: "o_5N_tt72qMx3WV",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthenticateResponse != nil {
        switch res.AuthenticateResponse.Type {
            case components.AuthenticateResponseTypeAuthenticateMultiStepResponse:
                // res.AuthenticateResponse.AuthenticateMultiStepResponse is populated
            case components.AuthenticateResponseTypeAuthenticateFinalResponse:
                // res.AuthenticateResponse.AuthenticateFinalResponse is populated
            default:
                // Unknown type - use res.AuthenticateResponse.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `xSessionToken`                                                                  | `string`                                                                         | :heavy_check_mark:                                                               | Session token received from `/initAuth` endpoint                                 |
| `body`                                                                           | [components.AuthenticateRequest](../../models/components/authenticaterequest.md) | :heavy_check_mark:                                                               | Request payload                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.AuthenticateResponse](../../models/operations/authenticateresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.ErrorResponse          | 400, 401                         | application/json                 |
| apierrors.RateLimitErrorResponse | 429                              | application/json                 |
| apierrors.ErrorResponse          | 500                              | application/json                 |
| apierrors.APIError               | 4XX, 5XX                         | \*/\*                            |

## RefreshToken

Get a new access token using a valid refresh token.

**Usage:**

- Pass the refresh token as a Bearer token in the Authorization header
- Returns a new access token and basic user information
- The request body and query string are not read

**Token Lifetimes:**

- Access token: 24 hours (configurable via `ACCESS_TOKEN_EXPIRY` environment variable)
- Refresh token: 30 days (configurable via `REFRESH_TOKEN_EXPIRY` environment variable)

**Best Practices:**

- Call this endpoint before the access token expires
- Store the new access token and continue using it for authenticated requests
- If refresh fails with 401, redirect user to login flow


### Example Usage

<!-- UsageSnippet language="go" operationID="refreshToken" method="post" path="/userAccount/refresh/token" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.UserAccount.RefreshToken(ctx, operations.RefreshTokenSecurity{
        ScopedToken: os.Getenv("PIPESHUB_SCOPED_TOKEN"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RefreshTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `security`                                                                         | [operations.RefreshTokenSecurity](../../models/operations/refreshtokensecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.RefreshTokenResponse](../../models/operations/refreshtokenresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetPasswordStatus

Returns whether the authenticated user has a password set. Users who signed in only through SSO do not.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPasswordStatus" method="get" path="/userAccount/password/status" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.UserAccount.GetPasswordStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPasswordStatusResponse](../../models/operations/getpasswordstatusresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SendPasswordSetupEmail

Sends the authenticated user a single-use link, to the email on their account, for setting a first password.
For users who have no password yet, for example those who signed in with Google. Users who already have one change it with `POST /userAccount/password/reset`.


### Example Usage

<!-- UsageSnippet language="go" operationID="sendPasswordSetupEmail" method="post" path="/userAccount/password/setup/email" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.UserAccount.SendPasswordSetupEmail(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SendPasswordSetupEmailResponse](../../models/operations/sendpasswordsetupemailresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404, 429      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ResetPassword

Reset the password for the currently authenticated user.

**Overview:**

Allows a logged-in user to change their password by providing the current password and a new password.

**Order of checks:** the access token is read before the body, so a request without
an `Authorization` header answers `400` whatever its body holds. Body fields other
than the three below are dropped by the validator, and the query string is not read.

**Effect on sessions:** the change ends refresh tokens issued before it. The response
carries a new access token.


### Example Usage

<!-- UsageSnippet language="go" operationID="resetPassword" method="post" path="/userAccount/password/reset" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.UserAccount.ResetPassword(ctx, operations.ResetPasswordRequest{
        CurrentPassword: "fR5Alu28cPCa984",
        NewPassword: "vcFGz9GLaOB88kV",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthenticatedPasswordResetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ResetPasswordRequest](../../models/operations/resetpasswordrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ResetPasswordResponse](../../models/operations/resetpasswordresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.ErrorResponse          | 400, 401, 404                    | application/json                 |
| apierrors.RateLimitErrorResponse | 429                              | application/json                 |
| apierrors.ErrorResponse          | 500                              | application/json                 |
| apierrors.APIError               | 4XX, 5XX                         | \*/\*                            |