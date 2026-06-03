# OAuthApps

## Overview

Manage OAuth 2.0 client applications registered with PipesHub.

OAuth apps allow third-party applications to access PipesHub APIs on behalf of users
or organizations. Each app receives a client ID and secret for authentication.

**Who can see which apps**
- **Everyone (including org admins)** sees and manages only OAuth apps **they created** (`createdBy`). Other members' apps are hidden (not listed; individual operations return not found).

**Who authorizes vs. client credentials**
- **Authorization code:** Any authenticated user in the workspace may complete consent for a valid `client_id`; issued tokens represent **that user**.
- **Client credentials:** Access tokens represent the **OAuth app creator** (who registered the client), not the caller.

**Scopes**
- `GET /oauth-clients/scopes` returns scopes grouped by category for the **signed-in user's role**.
- **Org admins** may register apps that request additional **admin-only** scopes; non-admins cannot select those scopes when creating or updating an app.

**App Types:**
- **Confidential clients**: Server-side apps that can securely store secrets
- **Public clients**: Browser/mobile apps that cannot securely store secrets (use PKCE)

**App Lifecycle:**
- Create apps with name, redirect URIs, allowed scopes, and optional URLs (homepage, privacy, terms)
- Regenerate secrets if compromised
- Suspend/activate apps to control access
- Revoke all tokens for emergency access removal


### Available Operations

* [ListOAuthApps](#listoauthapps) - List OAuth apps
* [CreateOAuthApp](#createoauthapp) - Create OAuth app
* [ListOAuthScopes](#listoauthscopes) - List available scopes
* [GetOAuthApp](#getoauthapp) - Get OAuth app details
* [UpdateOAuthApp](#updateoauthapp) - Update OAuth app
* [DeleteOAuthApp](#deleteoauthapp) - Delete OAuth app
* [RegenerateOAuthAppSecret](#regenerateoauthappsecret) - Regenerate client secret
* [SuspendOAuthApp](#suspendoauthapp) - Suspend OAuth app
* [ActivateOAuthApp](#activateoauthapp) - Activate suspended OAuth app
* [ListOAuthAppTokens](#listoauthapptokens) - List app tokens
* [RevokeAllOAuthAppTokens](#revokealloauthapptokens) - Revoke all app tokens

## ListOAuthApps

Returns a paginated list of OAuth apps registered by the signed-in user. Access is creator-scoped — even org admins only see apps they created themselves, so this endpoint is safe to use for per-user developer dashboards without leaking org-wide app metadata.

Each entry carries the full app configuration except the client secret, which is only ever returned at creation time and immediately after a regeneration.

Use the `status` query parameter to filter by lifecycle state (`active`, `suspended`, `revoked`) and `search` for a case-insensitive substring match against `name` or `description`.


### Example Usage

<!-- UsageSnippet language="go" operationID="listOAuthApps" method="get" path="/oauth-clients" -->
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

    res, err := s.OAuthApps.ListOAuthApps(ctx, pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAppListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `page`                                                                               | **int64*                                                                             | :heavy_minus_sign:                                                                   | Page number (matches `listAppsQuerySchema`: defaults to `1` when omitted or empty).<br/> |
| `limit`                                                                              | **int64*                                                                             | :heavy_minus_sign:                                                                   | Items per page (defaults to `20` when omitted or empty; max 100).<br/>               |
| `status`                                                                             | [*operations.ListOAuthAppsStatus](../../models/operations/listoauthappsstatus.md)    | :heavy_minus_sign:                                                                   | Filter by status                                                                     |
| `search`                                                                             | **string*                                                                            | :heavy_minus_sign:                                                                   | Search by app name or description (case-insensitive)                                 |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListOAuthAppsResponse](../../models/operations/listoauthappsresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 401, 403                                      | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## CreateOAuthApp

Register a new OAuth app for the organization. Any authenticated org member may create apps; the creator is recorded as the app's owner and is the only user who can subsequently read, update, suspend, activate, regenerate the secret of, or delete it.

The `clientSecret` is returned in this response **only** — it is stored hashed server-side and cannot be retrieved later. Persist it before exiting the create flow; if it is ever lost, rotate via `POST /oauth-clients/{appId}/regenerate-secret`.

`allowedScopes` is validated against the caller's role-aware scope set (see `GET /oauth-clients/scopes`). Org admins may include admin-only scopes; non-admins requesting a restricted scope receive `400`.

All `/oauth-clients/*` routes share a per-user rate limiter (default 1000 req/min, configurable via the `MAX_OAUTH_CLIENT_REQUESTS_PER_MINUTE` env var).


### Example Usage

<!-- UsageSnippet language="go" operationID="createOAuthApp" method="post" path="/oauth-clients" -->
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

    res, err := s.OAuthApps.CreateOAuthApp(ctx, components.CreateOAuthAppRequest{
        Name: "My Integration App",
        Description: pipeshub.Pointer("Integrates PipesHub with our internal tools"),
        RedirectUris: []string{
            "https://myapp.com/callback",
            "http://localhost:3000/callback",
        },
        AllowedGrantTypes: []components.CreateOAuthAppRequestAllowedGrantType{
            components.CreateOAuthAppRequestAllowedGrantTypeAuthorizationCode,
            components.CreateOAuthAppRequestAllowedGrantTypeRefreshToken,
        },
        AllowedScopes: []string{
            "openid",
            "profile",
            "read:records",
        },
        RefreshTokenLifetime: pipeshub.Pointer[int64](604800),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateOAuthAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.CreateOAuthAppRequest](../../models/components/createoauthapprequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateOAuthAppResponse](../../models/operations/createoauthappresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 400, 401, 403                                 | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## ListOAuthScopes

Returns the OAuth scopes the signed-in user is permitted to register on new or updated apps, grouped by category. Use this to populate scope-picker UIs and to validate `allowedScopes` client-side before submitting to `createOAuthApp` / `updateOAuthApp`.

The result is role-aware. Org admins (members of an admin user group) receive every registered scope; everyone else is filtered to exclude admin-only scopes: `org:write`, `org:admin`, `user:invite`, `user:delete`, `usergroup:write`, `team:write`, `config:write`, `crawl:write`, `crawl:delete`.

Each key in the `scopes` map matches the `category` field on the `OAuthScopeInfo` entries it contains. A category may appear with an empty array when every scope it contains is restricted for the caller — treat empty buckets as "no permitted scopes in this group", not as a missing category.

Shares the per-user rate limiter applied to every `/oauth-clients/*` route (default 1000 req/min, `MAX_OAUTH_CLIENT_REQUESTS_PER_MINUTE`).


### Example Usage

<!-- UsageSnippet language="go" operationID="listOAuthScopes" method="get" path="/oauth-clients/scopes" -->
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

    res, err := s.OAuthApps.ListOAuthScopes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthScopesGroupedResponse != nil {
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

**[*operations.ListOAuthScopesResponse](../../models/operations/listoauthscopesresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 401                                           | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## GetOAuthApp

Returns the full configuration of an OAuth app you registered. The `clientSecret` is never echoed back here; if you need a new one, call `POST /oauth-clients/{appId}/regenerate-secret`.

Access is creator-scoped: even org admins receive `404` for apps owned by other users. This avoids leaking app metadata across org members and keeps the read surface symmetric with `listOAuthApps`.


### Example Usage

<!-- UsageSnippet language="go" operationID="getOAuthApp" method="get" path="/oauth-clients/{appId}" -->
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

    res, err := s.OAuthApps.GetOAuthApp(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth app ID (MongoDB ObjectId)                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOAuthAppResponse](../../models/operations/getoauthappresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 401, 403, 404                                 | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## UpdateOAuthApp

Update an OAuth app's configuration. All body fields are optional — supply only what should change. URL fields (`homepageUrl`, `privacyPolicyUrl`, `termsOfServiceUrl`) accept `null` to clear them.

Creator-only: even org admins cannot edit apps owned by other users.

When modifying `allowedScopes`, the new set must remain a subset of the caller's role-aware scope list (same rule as `GET /oauth-clients/scopes`). When adding `authorization_code` to `allowedGrantTypes`, `redirectUris` becomes required and must contain at least one URI; otherwise the request is rejected with `400` by the Zod refine on `updateAppSchema`.

This endpoint never rotates the client secret — use `POST /oauth-clients/{appId}/regenerate-secret` for that.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateOAuthApp" method="put" path="/oauth-clients/{appId}" -->
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

    res, err := s.OAuthApps.UpdateOAuthApp(ctx, "<id>", components.UpdateOAuthAppRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateOAuthAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `appID`                                                                              | *string*                                                                             | :heavy_check_mark:                                                                   | OAuth app ID                                                                         |
| `body`                                                                               | [components.UpdateOAuthAppRequest](../../models/components/updateoauthapprequest.md) | :heavy_check_mark:                                                                   | Request payload                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateOAuthAppResponse](../../models/operations/updateoauthappresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 400, 401, 403, 404                            | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## DeleteOAuthApp

Soft-deletes an OAuth app. The app is flagged `isDeleted=true` on the `OAuthApp` document, removed from list/get responses for every caller, and all of its access and refresh tokens are revoked in the same operation. There is no restore endpoint — deletion is final.

Creator-only: even org admins cannot delete apps owned by other users.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteOAuthApp" method="delete" path="/oauth-clients/{appId}" -->
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

    res, err := s.OAuthApps.DeleteOAuthApp(ctx, "<id>")
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
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth app ID                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOAuthAppResponse](../../models/operations/deleteoauthappresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 401, 403, 404                                 | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## RegenerateOAuthAppSecret

Generates a fresh client secret for an OAuth app. The previous secret is invalidated immediately — any client still presenting it will fail token exchange at `POST /oauth2/token` until updated.

The new secret is returned in this response **only** and cannot be retrieved later. Pair this call with credential propagation to every integration that uses the app. If the rotation was triggered by a suspected leak, also call `POST /oauth-clients/{appId}/revoke-all-tokens` to invalidate already-issued access and refresh tokens instead of waiting for their natural expiry.

Creator-only: even org admins cannot rotate secrets for other users' apps.


### Example Usage

<!-- UsageSnippet language="go" operationID="regenerateOAuthAppSecret" method="post" path="/oauth-clients/{appId}/regenerate-secret" -->
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

    res, err := s.OAuthApps.RegenerateOAuthAppSecret(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegenerateOAuthAppSecretResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth app ID                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RegenerateOAuthAppSecretResponse](../../models/operations/regenerateoauthappsecretresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 401, 403, 404                                 | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## SuspendOAuthApp

Moves an OAuth app to `status: "suspended"`, blocking new token issuance at `POST /oauth2/token` and the authorization-code consent flow. Tokens that have already been issued remain valid until their natural expiry — call `POST /oauth-clients/{appId}/revoke-all-tokens` immediately afterwards if you need an immediate lockout.

Use this for temporary suspensions where you intend to reactivate later. For permanent removal, use `DELETE /oauth-clients/{appId}`. Suspending an app that is already suspended returns `400`.

Creator-only.


### Example Usage

<!-- UsageSnippet language="go" operationID="suspendOAuthApp" method="post" path="/oauth-clients/{appId}/suspend" -->
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

    res, err := s.OAuthApps.SuspendOAuthApp(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuspendOAuthAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth app ID                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SuspendOAuthAppResponse](../../models/operations/suspendoauthappresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 400, 401, 403, 404                            | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## ActivateOAuthApp

Moves a suspended OAuth app back to `status: "active"`, restoring its ability to authenticate and obtain new tokens via `POST /oauth2/token`.

A revoked app cannot be reactivated (returns `400`); the only path back is to register a new app. Activating an app that is already active also returns `400`.

Creator-only.


### Example Usage

<!-- UsageSnippet language="go" operationID="activateOAuthApp" method="post" path="/oauth-clients/{appId}/activate" -->
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

    res, err := s.OAuthApps.ActivateOAuthApp(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ActivateOAuthAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth app ID                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ActivateOAuthAppResponse](../../models/operations/activateoauthappresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 400, 401, 403, 404                            | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## ListOAuthAppTokens

Lists active access and refresh tokens currently issued to an OAuth app, sorted newest first. Useful for auditing app usage and picking specific tokens to investigate before a targeted revocation.

Each entry includes the token type (`access` or `refresh`), the user the token was issued for (omitted for client-credentials access tokens), the granted scopes, the issuance and expiry timestamps, and the revocation flag. Each type is capped at 100 most-recent rows server-side (`listTokensForApp` in `oauth_token.service.ts`); revoked and expired tokens are excluded.

Creator-only.


### Example Usage

<!-- UsageSnippet language="go" operationID="listOAuthAppTokens" method="get" path="/oauth-clients/{appId}/tokens" -->
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

    res, err := s.OAuthApps.ListOAuthAppTokens(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAppTokensListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth app ID                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOAuthAppTokensResponse](../../models/operations/listoauthapptokensresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 401, 403, 404                                 | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## RevokeAllOAuthAppTokens

Revokes every access and refresh token currently issued to an OAuth app, in a single operation. Use this for emergency credential rotation, suspected secret leaks, or as a follow-up to `POST /oauth-clients/{appId}/regenerate-secret` when you want existing sessions invalidated immediately rather than letting them expire naturally.

The response `count` is the total number of tokens revoked across both types. Clients of this app must then obtain new tokens via the standard OAuth flow.

Creator-only.


### Example Usage

<!-- UsageSnippet language="go" operationID="revokeAllOAuthAppTokens" method="post" path="/oauth-clients/{appId}/revoke-all-tokens" -->
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

    res, err := s.OAuthApps.RevokeAllOAuthAppTokens(ctx, "<id>")
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
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth app ID                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RevokeAllOAuthAppTokensResponse](../../models/operations/revokealloauthapptokensresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 401, 403, 404                                 | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |