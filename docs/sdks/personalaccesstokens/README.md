# PersonalAccessTokens

## Overview

Self-service, long-lived, scoped, revocable credentials that act as their
creator — unlike an OAuth app's `client_credentials` flow, which acts as
the app.

**Who can create one**
- **Any authenticated org member** — unlike OAuth apps, this is
  deliberately not admin-gated.

**How it's issued**
- Minted through the same OAuth access-token machinery as `/oauth2/token`,
  against one lazily-created, per-org synthetic OAuth app
  (`clientId: pat-system:<orgId>`) that every PAT in that org shares.
  That app is hidden from `/oauth-clients/*` — it never appears in your
  own OAuth app list and can't be managed through those routes.
- The raw token is prefixed `phpat_` ahead of the underlying JWT (see the
  `bearerAuth` security scheme) so it's recognizable to secret scanners.
  It's shown exactly once, at creation.

**Expiry and scopes**
- `expiryDays`: `30` (default), `90`, `365`, or `never`.
- Scopes default to the org's full configured `MCP_SCOPES` set if none
  are requested; `GET /personal-access-tokens/scopes` lists what's
  available.

**Admin visibility**
- Regular members only ever see and revoke their own tokens.
- Org admins can list and revoke *any* member's token via
  `/personal-access-tokens/admin*` — for incident response (a departed
  employee, a compromised laptop) — without needing the OAuth app CRUD
  access described above.


### Available Operations

* [ListPersonalAccessTokens](#listpersonalaccesstokens) - List your own personal access tokens
* [CreatePersonalAccessToken](#createpersonalaccesstoken) - Create a personal access token
* [ListPersonalAccessTokenScopes](#listpersonalaccesstokenscopes) - List scopes available for a new personal access token
* [RevokePersonalAccessToken](#revokepersonalaccesstoken) - Revoke one of your own personal access tokens
* [AdminListPersonalAccessTokens](#adminlistpersonalaccesstokens) - Admin: list every active personal access token in the org
* [AdminRevokePersonalAccessToken](#adminrevokepersonalaccesstoken) - Admin: revoke any user's personal access token by id

## ListPersonalAccessTokens

Lists the caller's own active (non-revoked, unexpired) personal
access tokens, newest first, capped at 100 rows server-side. Never
returns another user's tokens — see `GET /personal-access-tokens/admin`
for the org-admin, cross-user view.

Shares the same per-user rate limiter as `/oauth-clients/*`
(default 1000 req/min, `MAX_OAUTH_CLIENT_REQUESTS_PER_MINUTE`).


### Example Usage

<!-- UsageSnippet language="go" operationID="listPersonalAccessTokens" method="get" path="/personal-access-tokens" -->
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

    res, err := s.PersonalAccessTokens.ListPersonalAccessTokens(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPatResponse != nil {
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

**[*operations.ListPersonalAccessTokensResponse](../../models/operations/listpersonalaccesstokensresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 401                                           | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## CreatePersonalAccessToken

Mints a new personal access token for the caller. Deliberately not
admin-gated — any authenticated org member may create their own,
unlike OAuth app registration.

The token is minted against a lazily-created, per-org synthetic
OAuth app (`clientId: pat-system:<orgId>`) shared by every PAT in
that org — the same signing, hashing, and revocation machinery as
`/oauth2/token`, reused rather than duplicated.

`scopes` is validated against the org's configured `MCP_SCOPES`
env var, not the full role-aware OAuth-app scope catalog — a
non-admin can request any scope in that set.

The response's `accessToken` is shown **once**; only its SHA-256
hash is stored. It's prefixed `phpat_` (see the `bearerAuth`
security scheme).


### Example Usage

<!-- UsageSnippet language="go" operationID="createPersonalAccessToken" method="post" path="/personal-access-tokens" -->
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

    res, err := s.PersonalAccessTokens.CreatePersonalAccessToken(ctx, components.CreatePatRequest{
        Name: "Claude Desktop",
        Scopes: []string{
            "kb:read",
            "semantic:write",
        },
        ExpiryDays: pipeshub.Pointer(components.CreateExpiryDaysExpiryDaysEnum(
            components.ExpiryDaysEnumThirty,
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatePatResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.CreatePatRequest](../../models/components/createpatrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreatePersonalAccessTokenResponse](../../models/operations/createpersonalaccesstokenresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 400, 401                                      | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## ListPersonalAccessTokenScopes

Returns the org's configured `MCP_SCOPES` as a flat array of scope
definitions, for populating the create-token scope picker. Unlike
`GET /oauth-clients/scopes`, this is **not** grouped by category and
**not** role-aware — every org member sees the same set, since PAT
scope selection isn't gated by admin status.


### Example Usage

<!-- UsageSnippet language="go" operationID="listPersonalAccessTokenScopes" method="get" path="/personal-access-tokens/scopes" -->
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

    res, err := s.PersonalAccessTokens.ListPersonalAccessTokenScopes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.PatScopesListResponse != nil {
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

**[*operations.ListPersonalAccessTokenScopesResponse](../../models/operations/listpersonalaccesstokenscopesresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 401                                           | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## RevokePersonalAccessToken

Revokes a token by id, scoped to `{tokenId, clientId, callerUserId}`
— a caller can never revoke another user's token through this route,
even though everyone in the org shares the same underlying
`pat-system:<orgId>` client. Revocation takes effect immediately: the
token's next verification attempt fails, including one already in
flight.


### Example Usage

<!-- UsageSnippet language="go" operationID="revokePersonalAccessToken" method="delete" path="/personal-access-tokens/{tokenId}" -->
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

    res, err := s.PersonalAccessTokens.RevokePersonalAccessToken(ctx, "<id>", &components.RevokePatRequest{
        Reason: pipeshub.Pointer("rotated"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RevokePatResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |
| `tokenID`                                                                   | *string*                                                                    | :heavy_check_mark:                                                          | Personal access token ID                                                    |
| `body`                                                                      | [*components.RevokePatRequest](../../models/components/revokepatrequest.md) | :heavy_minus_sign:                                                          | Optional request body for Revoke personal access token                      |
| `opts`                                                                      | [][operations.Option](../../models/operations/option.md)                    | :heavy_minus_sign:                                                          | The options for this request.                                               |

### Response

**[*operations.RevokePersonalAccessTokenResponse](../../models/operations/revokepersonalaccesstokenresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 401, 404                                      | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## AdminListPersonalAccessTokens

Lists every active personal access token across every member of the
org, paginated, with each token's owner attached (including owners
who've since been deleted from the org — see `ownerDeleted` on
`AdminPatListItem`). For incident response: a departed employee or a
compromised laptop, where only the token's own creator could
otherwise see or revoke it.

Requires org-admin privileges (`userAdminCheck`) — note this returns
**`400`**, not `403`, for a non-admin caller (shared middleware
behavior across the codebase, not specific to this route).


### Example Usage

<!-- UsageSnippet language="go" operationID="adminListPersonalAccessTokens" method="get" path="/personal-access-tokens/admin" -->
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

    res, err := s.PersonalAccessTokens.AdminListPersonalAccessTokens(ctx, pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](100))
    if err != nil {
        log.Fatal(err)
    }
    if res.AdminPatListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `page`                                                            | **int64*                                                          | :heavy_minus_sign:                                                | Page number (defaults to `1` when omitted or empty)               |
| `limit`                                                           | **int64*                                                          | :heavy_minus_sign:                                                | Items per page (defaults to `100` when omitted or empty; max 100) |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |

### Response

**[*operations.AdminListPersonalAccessTokensResponse](../../models/operations/adminlistpersonalaccesstokensresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 400, 401                                      | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## AdminRevokePersonalAccessToken

Revokes a token by id, scoped to the org's PAT client but **not** to
a specific owning user — the admin counterpart to
`DELETE /personal-access-tokens/{tokenId}`. Requires org-admin
privileges (`userAdminCheck`); returns `400` (not `403`) for a
non-admin caller, same as `GET /personal-access-tokens/admin`.


### Example Usage

<!-- UsageSnippet language="go" operationID="adminRevokePersonalAccessToken" method="delete" path="/personal-access-tokens/admin/{tokenId}" -->
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

    res, err := s.PersonalAccessTokens.AdminRevokePersonalAccessToken(ctx, "<id>", &components.RevokePatRequest{
        Reason: pipeshub.Pointer("rotated"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RevokePatResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |
| `tokenID`                                                                   | *string*                                                                    | :heavy_check_mark:                                                          | Personal access token ID                                                    |
| `body`                                                                      | [*components.RevokePatRequest](../../models/components/revokepatrequest.md) | :heavy_minus_sign:                                                          | Optional request body for Admin revoke personal access token                |
| `opts`                                                                      | [][operations.Option](../../models/operations/option.md)                    | :heavy_minus_sign:                                                          | The options for this request.                                               |

### Response

**[*operations.AdminRevokePersonalAccessTokenResponse](../../models/operations/adminrevokepersonalaccesstokenresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.ApplicationJSONErrorResponse        | 400, 401, 404                                 | application/json                              |
| apierrors.OAuthClientManagementRateLimitError | 429                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |