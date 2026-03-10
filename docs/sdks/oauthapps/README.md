# OauthApps

## Overview

### Available Operations

* [List](#list) - List OAuth apps
* [Create](#create) - Create OAuth app
* [ListScopes](#listscopes) - List available scopes
* [Get](#get) - Get OAuth app details
* [Update](#update) - Update OAuth app
* [Delete](#delete) - Delete OAuth app
* [RegenerateSecret](#regeneratesecret) - Regenerate client secret
* [Suspend](#suspend) - Suspend OAuth app
* [Activate](#activate) - Activate suspended OAuth app
* [ListTokens](#listtokens) - List app tokens
* [RevokeAllTokens](#revokealltokens) - Revoke all app tokens

## List

List all OAuth apps registered for the organization.
<br><br>
Returns a paginated list of apps with their configuration (excluding secrets).
<br><br>
<b>Filters:</b><br>
- `status` - Filter by app status (active, suspended, revoked)<br>
- `search` - Search by app name


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

    res, err := s.OauthApps.List(ctx, pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAppListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `page`                                                                            | **int64*                                                                          | :heavy_minus_sign:                                                                | Page number                                                                       |
| `limit`                                                                           | **int64*                                                                          | :heavy_minus_sign:                                                                | Items per page                                                                    |
| `status`                                                                          | [*operations.ListOAuthAppsStatus](../../models/operations/listoauthappsstatus.md) | :heavy_minus_sign:                                                                | Filter by status                                                                  |
| `search`                                                                          | **string*                                                                         | :heavy_minus_sign:                                                                | Search by app name                                                                |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*operations.ListOAuthAppsResponse](../../models/operations/listoauthappsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new OAuth app for the organization.
<br><br>
<b>Important:</b> The client secret is only returned once during creation.
Store it securely - it cannot be retrieved later. If lost, you'll need to
regenerate it.
<br><br>
<b>Admin Only:</b> Requires admin privileges.
<br><br>
<b>Rate Limited:</b> 10 requests per minute.


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

    res, err := s.OauthApps.Create(ctx, components.CreateOAuthAppRequest{
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAppWithSecret != nil {
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListScopes

List all available OAuth scopes that can be requested by apps.
<br><br>
Returns scope names, descriptions, and categories for display
in app configuration UI.


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

    res, err := s.OauthApps.ListScopes(ctx)
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

**[*operations.ListOAuthScopesResponse](../../models/operations/listoauthscopesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get details of a specific OAuth app.
<br><br>
Returns app configuration without the client secret.


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

    res, err := s.OauthApps.Get(ctx, "<id>")
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update an OAuth app's configuration.
<br><br>
<b>Admin Only:</b> Requires admin privileges.
<br><br>
<b>Rate Limited:</b> 10 requests per minute.
<br><br>
<b>Note:</b> To regenerate the client secret, use the
`/oauth-clients/{appId}/regenerate-secret` endpoint.


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

    res, err := s.OauthApps.Update(ctx, "<id>", components.UpdateOAuthAppRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAppResponse != nil {
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete (soft delete) an OAuth app.
<br><br>
This marks the app as deleted and revokes all its tokens.
The app cannot be restored after deletion.
<br><br>
<b>Admin Only:</b> Requires admin privileges.
<br><br>
<b>Rate Limited:</b> 10 requests per minute.


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

    res, err := s.OauthApps.Delete(ctx, "<id>")
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RegenerateSecret

Regenerate the client secret for an OAuth app.
<br><br>
The old secret is immediately invalidated. Any clients using the old
secret will fail to authenticate until updated with the new secret.
<br><br>
<b>Important:</b> The new secret is only returned once. Store it securely.
<br><br>
<b>Admin Only:</b> Requires admin privileges.
<br><br>
<b>Rate Limited:</b> 10 requests per minute.


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

    res, err := s.OauthApps.RegenerateSecret(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthAppWithSecret != nil {
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Suspend

Suspend an OAuth app, preventing it from authenticating or issuing tokens.
<br><br>
Existing tokens remain valid until they expire, but no new tokens can
be obtained. Use this for temporary access suspension.
<br><br>
<b>Admin Only:</b> Requires admin privileges.
<br><br>
<b>Rate Limited:</b> 10 requests per minute.


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

    res, err := s.OauthApps.Suspend(ctx, "<id>")
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
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth app ID                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SuspendOAuthAppResponse](../../models/operations/suspendoauthappresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Activate

Reactivate a suspended OAuth app, allowing it to authenticate and issue tokens again.
<br><br>
<b>Admin Only:</b> Requires admin privileges.
<br><br>
<b>Rate Limited:</b> 10 requests per minute.


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

    res, err := s.OauthApps.Activate(ctx, "<id>")
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
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth app ID                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ActivateOAuthAppResponse](../../models/operations/activateoauthappresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListTokens

List all active tokens issued to an OAuth app.
<br><br>
Useful for monitoring app usage and identifying tokens to revoke.
<br><br>
<b>Admin Only:</b> Requires admin privileges.


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

    res, err := s.OauthApps.ListTokens(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthTokenListItems != nil {
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RevokeAllTokens

Revoke all tokens (access and refresh) issued to an OAuth app.
<br><br>
Use this for emergency access removal or when rotating credentials.
<br><br>
<b>Admin Only:</b> Requires admin privileges.
<br><br>
<b>Rate Limited:</b> 10 requests per minute.


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

    res, err := s.OauthApps.RevokeAllTokens(ctx, "<id>")
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |