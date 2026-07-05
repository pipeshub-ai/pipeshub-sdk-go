# OAuthProvider

## Overview

PipesHub OAuth 2.0 Authorization Server implementing RFC 6749, RFC 7636 (PKCE), and OpenID Connect.

**Supported Grant Types:**
- `authorization_code` - Standard OAuth flow with PKCE support
- `client_credentials` - Machine-to-machine authentication
- `refresh_token` - Token refresh for long-lived access

**Security Features:**
- PKCE (Proof Key for Code Exchange) for public clients
- State parameter for CSRF protection
- Configurable token lifetimes
- Token revocation and introspection

**OpenID Connect:**
- ID tokens with standard claims
- UserInfo endpoint for profile data
- Discovery endpoint for automatic configuration

**Machine tokens (`client_credentials`) — gateway and downstream identity:**
Access tokens may encode **`userId === client_id`**. The **Node.js API gateway** resolves the effective user to the OAuth **app creator**: first using the JWT **`createdBy`** claim when present, otherwise by loading the OAuth app by **`client_id`** from the registry. After verification it sets the authenticated session to that creator.

**Python services:** Validate `Authorization: Bearer` as today and use the JWT payload’s **`userId`** as-is for scopes and user-scoped logic (which may still equal **`client_id`** for machine tokens).

**Operational note:** Prefer tokens whose JWT already carries the creator as **`userId`**; use **`POST /oauth-clients/{appId}/revoke-all-tokens`** and obtain new tokens from **`POST /oauth2/token`** when rotating integrations.


### Available Operations

* [OauthToken](#oauthtoken) - Exchange authorization code for tokens
* [OauthRevoke](#oauthrevoke) - Revoke an access or refresh token
* [OauthIntrospect](#oauthintrospect) - Introspect a token

## OauthToken

OAuth 2.0 Token Endpoint (RFC 6749 Section 4.1.3).

Exchanges an authorization code, client credentials, or refresh token for access tokens.

**Grant Types:**
- `authorization_code`: Exchange auth code for tokens (user-based)
- `client_credentials`: Get tokens for machine-to-machine auth
- `refresh_token`: Get new access token using refresh token

For **`client_credentials`**, access tokens represent the **OAuth app creator** (the user who registered the client). The JWT may encode **`userId === client_id`**; the **Node API gateway** resolves the creator (**`createdBy`** claim or OAuth app lookup) — see **OAuth Provider** tag.

**Client Authentication:**
Can be provided via:
- HTTP Basic auth: `Authorization: Basic base64(client_id:client_secret)`
- Request body: `client_id` and `client_secret` parameters

**PKCE Verification:**
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

    res, err := s.OAuthProvider.OauthToken(ctx, components.OAuthTokenRequest{
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

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ErrorResponse                       | 400                                           | application/json                              |
| apierrors.OAuthErrorResponse                  | 401                                           | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## OauthRevoke

OAuth 2.0 Token Revocation Endpoint (RFC 7009).

Revokes an access token or refresh token, preventing further use.
Revoking a refresh token also invalidates associated access tokens.

**Use Cases:**
- User logs out of third-party app
- User revokes app access from account settings
- Security incident response

**Note:** Returns 200 OK even if token was already revoked or invalid
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

    res, err := s.OAuthProvider.OauthRevoke(ctx, components.OAuthRevokeRequest{
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

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.OAuthErrorResponse                  | 401                                           | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## OauthIntrospect

OAuth 2.0 Token Introspection Endpoint (RFC 7662).

Check if a token is active and retrieve its metadata.

**Use Cases:**
- Resource servers validating tokens
- Debugging token issues
- Checking token scopes before processing requests

**Response:**
- Active token: Returns `active: true` with token metadata
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

    res, err := s.OAuthProvider.OauthIntrospect(ctx, components.OAuthIntrospectRequest{
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

    res, err := s.OAuthProvider.OauthIntrospect(ctx, components.OAuthIntrospectRequest{
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

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.OAuthErrorResponse                  | 401                                           | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |