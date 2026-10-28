# OAuthProvider

## Overview

PipesHub OAuth 2.0 Authorization Server implementing RFC 6749, RFC 7636 (PKCE), and OpenID Connect.

**Supported Grant Types:**
- `authorization_code` - Standard OAuth flow with PKCE support
- `client_credentials` - Machine-to-machine authentication
- `refresh_token` - Token refresh for long-lived access
- `urn:ietf:params:oauth:grant-type:device_code` - Device authorization grant (RFC 8628) for devices and CLIs without a browser; started with `POST /oauth2/device_authorization`

**Dynamic Client Registration (RFC 7591):** `POST /oauth2/register`, only when the instance enables it.

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
Access tokens may encode **`userId === client_id`**. The **Node.js API gateway** resolves the effective user to the application's **token identity**: the service account it is pointed at (**`tokenIdentityUserId`**), or its creator for an application registered before a service account was required and never bound. It loads the application by **`client_id`** on every such token and takes the identity from that record.

The JWT's **`createdBy`** claim is a cross-check rather than a source: despite its name it carries the identity, not who registered the application, and a token whose claim disagrees with the record is refused with **401** ("This token was issued for an identity the application no longer acts as"), and so is a token whose app was deleted ("OAuth app not found or revoked"). Pointing an application at a different service account also revokes the tokens it has already issued, so a stale claim should not arrive in the first place.

**Python services:** Validate `Authorization: Bearer` as today and use the JWT payload’s **`userId`** as-is for scopes and user-scoped logic. Where a machine token carries **`userId === client_id`** and a **`createdBy`** claim, Python copies that claim onto **`userId`** first, so both halves of the product agree on the identity.

**Operational note:** Prefer tokens whose JWT already carries the identity as **`userId`**; use **`POST /oauth-clients/{appId}/revoke-all-tokens`** and obtain new tokens from **`POST /oauth2/token`** when rotating integrations.


### Available Operations

* [OauthToken](#oauthtoken) - Exchange authorization code for tokens
* [OauthRevoke](#oauthrevoke) - Revoke an access or refresh token
* [OauthIntrospect](#oauthintrospect) - Introspect a token

## OauthToken

OAuth 2.0 Token Endpoint (RFC 6749 Section 4.1.3).

Exchanges an authorization code, client credentials, refresh token, device code or
service token for access tokens.
Successful responses carry `Cache-Control: no-store` and `Pragma: no-cache`.

**Grant Types:**
- `authorization_code`: Exchange auth code for tokens (user-based)
- `client_credentials`: Get tokens for machine-to-machine auth
- `urn:ietf:params:oauth:grant-type:device_code`: Poll for the tokens of a device authorization
  (RFC 8628); answered `400` with `authorization_pending` until the user approves
- `refresh_token`: Get new access token using refresh token. The refresh token must have been issued to the presenting `client_id`; one issued to another client is rejected with `400 invalid_grant` and is left untouched (not rotated, not revoked).
- `urn:ietf:params:oauth:grant-type:token-exchange`: Service-account
  delegation (RFC 8693). A service account exchanges its own token for
  a five-minute token that acts as one named person from its
  allow-list. **Enterprise edition.** The Community build answers this
  grant with `unsupported_grant_type` and leaves it out of discovery,
  so check `grant_types_supported` rather than assuming it is there. **This grant sends no `client_id` or `client_secret`** —
  the `subject_token` is the authentication, since it already carries
  a verified identity and the scopes that bound it. Refusals about the
  person asked for all return `400 invalid_target` with one wording;
  the cause is recorded in the delegation log, which only an
  administrator can read.

For **`client_credentials`**, access tokens represent the **service account the app is pointed at** (`tokenIdentityUserId`), or the **OAuth app creator** for an app registered before that was required and never bound. The JWT may encode **`userId === client_id`**; the **Node API gateway** loads the application record on every request and resolves the identity from it, treating the **`createdBy`** claim as a cross-check and refusing a token whose claim no longer matches it — see **OAuth Provider** tag.

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
| apierrors.BadRequest                          | 400                                           | application/json                              |
| apierrors.OAuthErrorResponse                  | 401                                           | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.ErrorResponse                       | 500                                           | application/json                              |
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
| apierrors.ErrorResponse                       | 400                                           | application/json                              |
| apierrors.OAuthErrorResponse                  | 401                                           | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.ErrorResponse                       | 500                                           | application/json                              |
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
| apierrors.ErrorResponse                       | 400                                           | application/json                              |
| apierrors.OAuthErrorResponse                  | 401                                           | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.ErrorResponse                       | 500                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |