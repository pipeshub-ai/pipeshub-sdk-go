# UserAccount

## Overview

### Available Operations

* [InitAuth](#initauth) - Initialize authentication session
* [Authenticate](#authenticate) - Authenticate user with credentials
* [GenerateLoginOtp](#generateloginotp) - Generate and send OTP for login
* [RequestPasswordReset](#requestpasswordreset) - Request password reset email
* [ResetPasswordWithToken](#resetpasswordwithtoken) - Reset password with email token
* [RefreshToken](#refreshtoken) - Refresh access token
* [Logout](#logout) - Logout current session
* [ResetPassword](#resetpassword) - Reset password

## InitAuth

Initialize an authentication session for a user by email address.
This is the first step in the multi-step authentication flow.
<br><br>
<b>Flow:</b><br>
1. Call this endpoint with the user's email<br>
2. Receive a session token in the <code>x-session-token</code> response header<br>
3. Use the session token in subsequent <code>/authenticate</code> calls<br>
4. The response includes <code>allowedMethods</code> for the first authentication step
<br><br>
<b>Session Token:</b><br>
- Stored in response header <code>x-session-token</code><br>
- Required for all subsequent authentication requests<br>
- Expires after a configured timeout period
<br><br>
<b>Multi-Factor Authentication:</b><br>
If the organization has MFA configured, you'll need to complete multiple
authentication steps. Each step completion returns the next step's allowed methods.


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

    res, err := s.UserAccount.InitAuth(ctx, components.InitAuthRequest{
        Email: "user@example.com",
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

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.AuthError | 400, 403, 404       | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |

## Authenticate

Authenticate a user using the specified method and credentials.
Requires a valid session token from <code>/initAuth</code>.
<br><br>
<b>Credential Formats by Method:</b><br>
- <code>password</code>: <code>{ "credentials": { "password": "your-password" } }</code><br>
- <code>otp</code>: <code>{ "credentials": { "otp": "123456" } }</code> (6-digit code, valid for 10 minutes)<br>
- <code>google</code>: <code>{ "credentials": "google-id-token-string" }</code><br>
- <code>microsoft</code>: <code>{ "credentials": { "accessToken": "...", "idToken": "..." } }</code><br>
- <code>azureAd</code>: <code>{ "credentials": { "accessToken": "...", "idToken": "..." } }</code><br>
- <code>oauth</code>: <code>{ "credentials": { "accessToken": "...", "idToken": "..." } }</code><br>
- <code>samlSso</code>: Handled via redirect flow (use <code>/saml/signIn</code> instead)
<br><br>
<b>Multi-Step Response:</b><br>
If organization uses MFA, successful authentication returns:<br>
- <code>status: "success"</code> with <code>nextStep</code> and <code>allowedMethods</code> for next step
<br><br>
<b>Fully Authenticated Response:</b><br>
After completing all steps:<br>
- <code>message: "Fully authenticated"</code> with <code>accessToken</code> (1hr) and <code>refreshToken</code> (7d)
<br><br>
<b>Security:</b><br>
- Account locks after 5 consecutive failed attempts<br>
- CAPTCHA may be required if enabled (pass <code>cf-turnstile-response</code>)


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
        Method: components.MethodOauth,
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
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `xSessionToken`                                                                  | *string*                                                                         | :heavy_check_mark:                                                               | Session token received from `/initAuth` endpoint                                 |
| `body`                                                                           | [components.AuthenticateRequest](../../models/components/authenticaterequest.md) | :heavy_check_mark:                                                               | Request payload                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.AuthenticateResponse](../../models/operations/authenticateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.AuthError     | 400, 401, 403, 404, 410 | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GenerateLoginOtp

Generate and send a 6-digit one-time password (OTP) to the user's email.
Use this endpoint before authenticating with the <code>otp</code> method.
<br><br>
<b>OTP Details:</b><br>
- 6 digits numeric code<br>
- Valid for <b>10 minutes</b> after generation<br>
- Sent to user's registered email address
<br><br>
<b>Rate Limiting:</b><br>
- Multiple OTP requests may be rate-limited<br>
- Wait for the current OTP to expire before requesting a new one
<br><br>
<b>CAPTCHA:</b><br>
If Cloudflare Turnstile is enabled, include <code>cf-turnstile-response</code> in the request body.


### Example Usage

<!-- UsageSnippet language="go" operationID="generateLoginOtp" method="post" path="/userAccount/login/otp/generate" -->
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

    res, err := s.UserAccount.GenerateLoginOtp(ctx, components.OtpGenerateRequest{
        Email: "Garland.Sipes42@hotmail.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.OtpGenerateRequest](../../models/components/otpgeneraterequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GenerateLoginOtpResponse](../../models/operations/generateloginotpresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.AuthError | 400, 403, 404       | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |

## RequestPasswordReset

Send a password reset link to the user's email.
The link contains a time-limited token that can be used to reset the password.
<br><br>
<b>Note:</b> This endpoint always returns 200 even if the email doesn't exist (to prevent email enumeration).


### Example Usage

<!-- UsageSnippet language="go" operationID="forgotPassword" method="post" path="/userAccount/password/forgot" -->
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

    res, err := s.UserAccount.RequestPasswordReset(ctx, components.ForgotPasswordRequest{
        Email: "Barton.Gutkowski68@yahoo.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.ForgotPasswordRequest](../../models/components/forgotpasswordrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ForgotPasswordResponse](../../models/operations/forgotpasswordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ResetPasswordWithToken

Reset password using a token received via email from the forgot password flow.
<br><br>
<b>Password Requirements:</b><br>
- Minimum 8 characters<br>
- At least 1 uppercase letter<br>
- At least 1 lowercase letter<br>
- At least 1 number<br>
- At least 1 special character (#?!@$%^&*-)
<br><br>
<b>Security Notes:</b><br>
- Token is single-use and expires after a set time<br>
- A new access token is returned upon successful reset


### Example Usage

<!-- UsageSnippet language="go" operationID="resetPasswordWithToken" method="post" path="/userAccount/password/reset/token" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.UserAccount.ResetPasswordWithToken(ctx, components.TokenPasswordResetRequest{
        Password: "H9GEHoL829GXj06",
    }, operations.ResetPasswordWithTokenSecurity{
        ScopedToken: os.Getenv("PIPESHUB_SCOPED_TOKEN"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PasswordResetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [components.TokenPasswordResetRequest](../../models/components/tokenpasswordresetrequest.md)           | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.ResetPasswordWithTokenSecurity](../../models/operations/resetpasswordwithtokensecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ResetPasswordWithTokenResponse](../../models/operations/resetpasswordwithtokenresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.AuthError | 400                 | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |

## RefreshToken

Get a new access token using a valid refresh token.
<br><br>
<b>Usage:</b><br>
- Pass the refresh token as a Bearer token in the Authorization header<br>
- Returns a new access token (1 hour expiry) and basic user information
<br><br>
<b>Token Lifetimes:</b><br>
- Access token: 1 hour<br>
- Refresh token: 7 days
<br><br>
<b>Best Practices:</b><br>
- Call this endpoint before the access token expires<br>
- Store the new access token and continue using it for authenticated requests<br>
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

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.AuthError | 401, 404            | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |

## Logout

Log out the current user session and invalidate tokens.
<br><br>
<b>Effects:</b><br>
- Invalidates the current access token<br>
- Clears server-side session data<br>
- Client should also clear stored tokens locally
<br><br>
<b>Note:</b> This endpoint requires the access token, not the refresh token.


### Example Usage

<!-- UsageSnippet language="go" operationID="logout" method="post" path="/userAccount/logout/manual" -->
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

    res, err := s.UserAccount.Logout(ctx)
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

**[*operations.LogoutResponse](../../models/operations/logoutresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ResetPassword

Reset the password for the currently authenticated user.<br><br>
<b>Overview:</b><br>
Allows a logged-in user to change their password by providing the current password and a new password.


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
    if res.Object != nil {
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

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| apierrors.BadRequestError | 400                       | application/json          |
| apierrors.APIError        | 4XX, 5XX                  | \*/\*                     |