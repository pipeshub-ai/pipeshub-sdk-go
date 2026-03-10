# OauthProvider

## Overview

### Available Operations

* [Authorize](#authorize) - Initiate OAuth authorization flow
* [AuthorizeConsent](#authorizeconsent) - Submit authorization consent
* [ExchangeToken](#exchangetoken) - Exchange authorization code for tokens
* [RevokeToken](#revoketoken) - Revoke an access or refresh token
* [IntrospectToken](#introspecttoken) - Introspect a token

## Authorize

OAuth 2.0 Authorization Endpoint (RFC 6749 Section 4.1.1).
<br><br>
Initiates the authorization code flow. Users are redirected here by OAuth clients
to authorize access to their account.
<br><br>
<b>Flow:</b><br>
1. Client redirects user to this endpoint with required parameters<br>
2. If not logged in, user is redirected to PipesHub login<br>
3. User sees consent page with requested scopes<br>
4. User grants or denies consent<br>
5. User is redirected back to client with authorization code
<br><br>
<b>PKCE Support (RFC 7636):</b><br>
- Required for public clients (SPA, mobile apps)<br>
- Recommended for confidential clients<br>
- Use S256 method (SHA256 hash of code_verifier)
<br><br>
<b>Security:</b><br>
- Always use HTTPS in production<br>
- State parameter provides CSRF protection<br>
- Redirect URI must match registered URIs exactly


### Example Usage

<!-- UsageSnippet language="go" operationID="oauthAuthorize" method="get" path="/oauth2/authorize" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.OauthProvider.Authorize(ctx, operations.OauthAuthorizeRequest{
        ResponseType: operations.ResponseTypeCode,
        ClientID: "<id>",
        RedirectURI: "https://coordinated-dime.biz/",
        Scope: "openid profile email read:records",
        State: "Delaware",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.OauthAuthorizeRequest](../../models/operations/oauthauthorizerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.OauthAuthorizeResponse](../../models/operations/oauthauthorizeresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.OAuthErrorResponse | 400                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## AuthorizeConsent

Submit user's consent decision for OAuth authorization.
<br><br>
Called after user reviews the consent page and makes a decision.
This endpoint generates an authorization code if consent is granted.
<br><br>
<b>Responses:</b><br>
- Consent granted: Redirects to client with authorization code<br>
- Consent denied: Redirects to client with `access_denied` error


### Example Usage

<!-- UsageSnippet language="go" operationID="oauthAuthorizeConsent" method="post" path="/oauth2/authorize" -->
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

    res, err := s.OauthProvider.AuthorizeConsent(ctx, components.OAuthConsentRequest{
        ClientID: "<id>",
        RedirectURI: "https://excellent-license.com",
        Scope: "<value>",
        State: "South Dakota",
        Consent: components.ConsentDenied,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.OAuthConsentRequest](../../models/components/oauthconsentrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.OauthAuthorizeConsentResponse](../../models/operations/oauthauthorizeconsentresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.OAuthErrorResponse | 400                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## ExchangeToken

OAuth 2.0 Token Endpoint (RFC 6749 Section 4.1.3).
<br><br>
Exchanges an authorization code, client credentials, or refresh token for access tokens.
<br><br>
<b>Grant Types:</b><br>
- `authorization_code`: Exchange auth code for tokens (user-based)<br>
- `client_credentials`: Get tokens for machine-to-machine auth<br>
- `refresh_token`: Get new access token using refresh token
<br><br>
<b>Client Authentication:</b><br>
Can be provided via:<br>
- HTTP Basic auth: `Authorization: Basic base64(client_id:client_secret)`<br>
- Request body: `client_id` and `client_secret` parameters
<br><br>
<b>PKCE Verification:</b><br>
If authorization used PKCE, the `code_verifier` must be provided and will be
verified against the stored code challenge.


### Example Usage

<!-- UsageSnippet language="go" operationID="oauthToken" method="post" path="/oauth2/token" -->
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

    res, err := s.OauthProvider.ExchangeToken(ctx, components.OAuthTokenRequest{
        GrantType: components.GrantTypeClientCredentials,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.OAuthTokenRequest](../../models/components/oauthtokenrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.OauthTokenResponse](../../models/operations/oauthtokenresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.OAuthErrorResponse | 400, 401                     | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## RevokeToken

OAuth 2.0 Token Revocation Endpoint (RFC 7009).
<br><br>
Revokes an access token or refresh token, preventing further use.
Revoking a refresh token also invalidates associated access tokens.
<br><br>
<b>Use Cases:</b><br>
- User logs out of third-party app<br>
- User revokes app access from account settings<br>
- Security incident response
<br><br>
<b>Note:</b> Returns 200 OK even if token was already revoked or invalid
(per RFC 7009, to prevent token enumeration).


### Example Usage

<!-- UsageSnippet language="go" operationID="oauthRevoke" method="post" path="/oauth2/revoke" -->
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

    res, err := s.OauthProvider.RevokeToken(ctx, components.OAuthRevokeRequest{
        Token: "<value>",
        ClientID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.OAuthRevokeRequest](../../models/components/oauthrevokerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.OauthRevokeResponse](../../models/operations/oauthrevokeresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.OAuthErrorResponse | 401                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## IntrospectToken

OAuth 2.0 Token Introspection Endpoint (RFC 7662).
<br><br>
Check if a token is active and retrieve its metadata.
<br><br>
<b>Use Cases:</b><br>
- Resource servers validating tokens<br>
- Debugging token issues<br>
- Checking token scopes before processing requests
<br><br>
<b>Response:</b><br>
- Active token: Returns `active: true` with token metadata<br>
- Invalid/expired/revoked token: Returns only `active: false`


### Example Usage: active

<!-- UsageSnippet language="go" operationID="oauthIntrospect" method="post" path="/oauth2/introspect" example="active" -->
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

    res, err := s.OauthProvider.IntrospectToken(ctx, components.OAuthIntrospectRequest{
        Token: "<value>",
        ClientID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthIntrospectResponse != nil {
        // handle response
    }
}
```
### Example Usage: inactive

<!-- UsageSnippet language="go" operationID="oauthIntrospect" method="post" path="/oauth2/introspect" example="inactive" -->
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

    res, err := s.OauthProvider.IntrospectToken(ctx, components.OAuthIntrospectRequest{
        Token: "<value>",
        ClientID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthIntrospectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.OAuthIntrospectRequest](../../models/components/oauthintrospectrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.OauthIntrospectResponse](../../models/operations/oauthintrospectresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.OAuthErrorResponse | 401                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |