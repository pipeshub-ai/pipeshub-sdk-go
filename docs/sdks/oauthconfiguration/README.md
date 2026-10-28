# OAuthConfiguration

## Overview

Admin management of OAuth credentials for connector types

### Available Operations

* [GetOAuthRegistry](#getoauthregistry) - List OAuth-capable connector types
* [GetOAuthConnectorType](#getoauthconnectortype) - Get OAuth connector type details
* [ListOAuthConfigs](#listoauthconfigs) - List OAuth configurations
* [ListOAuthConfigsByType](#listoauthconfigsbytype) - List OAuth configs for connector type
* [CreateOAuthConfig](#createoauthconfig) - Create OAuth configuration
* [GetOAuthConfig](#getoauthconfig) - Get OAuth configuration
* [UpdateOAuthConfig](#updateoauthconfig) - Update OAuth configuration
* [DeleteOAuthConfig](#deleteoauthconfig) - Delete OAuth configuration

## GetOAuthRegistry

List the connector and toolset types that support OAuth, from the OAuth registry.<br><br>
Each entry carries the provider URLs, the scopes requested, and the `authFields`
to collect when creating an OAuth configuration for that type.<br><br>
Any authenticated user can call this; it is not limited to admins.


### Example Usage

<!-- UsageSnippet language="go" operationID="getOAuthRegistry" method="get" path="/oauth/registry" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.OAuthConfiguration.GetOAuthRegistry(ctx, pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20), nil)
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
| `page`                                                   | `*int64`                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `limit`                                                  | `*int64`                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `search`                                                 | `*string`                                                | :heavy_minus_sign:                                       | Matches the connector type's name, group or description  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOAuthRegistryResponse](../../models/operations/getoauthregistryresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetOAuthConnectorType

Get the OAuth registry entry of one connector or toolset type.<br><br>
`connectorType` is the registered name and is case-sensitive (e.g. `Confluence`).


### Example Usage

<!-- UsageSnippet language="go" operationID="getOAuthConnectorType" method="get" path="/oauth/registry/{connectorType}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.OAuthConfiguration.GetOAuthConnectorType(ctx, "Confluence", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 | Example                                                                                                     |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |                                                                                                             |
| `connectorType`                                                                                             | `string`                                                                                                    | :heavy_check_mark:                                                                                          | Registered name of the type, case-sensitive                                                                 | Confluence                                                                                                  |
| `page`                                                                                                      | `*int64`                                                                                                    | :heavy_minus_sign:                                                                                          | Validated like the `page` of `GET /oauth/registry`, then ignored; the reply is the same with or without it. |                                                                                                             |
| `limit`                                                                                                     | `*int64`                                                                                                    | :heavy_minus_sign:                                                                                          | Validated like the `limit` of `GET /oauth/registry`, then ignored.                                          |                                                                                                             |
| `search`                                                                                                    | `*string`                                                                                                   | :heavy_minus_sign:                                                                                          | Validated as a single text value, then ignored.                                                             |                                                                                                             |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |                                                                                                             |

### Response

**[*operations.GetOAuthConnectorTypeResponse](../../models/operations/getoauthconnectortyperesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListOAuthConfigs

List the saved OAuth apps of every connector type in the organization.<br><br>
Every caller, admin or not, gets the same essential fields (ID, name, connector type, app
metadata, timestamps). Credentials are never included; an admin reads them per app with
<code>GET /oauth/{connectorType}/{configId}</code>.


### Example Usage

<!-- UsageSnippet language="go" operationID="listOAuthConfigs" method="get" path="/oauth" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.OAuthConfiguration.ListOAuthConfigs(ctx, pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `page`                                                                                                 | `*int64`                                                                                               | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `limit`                                                                                                | `*int64`                                                                                               | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `search`                                                                                               | `*string`                                                                                              | :heavy_minus_sign:                                                                                     | Case-insensitive substring matched against the name, connector type, group, description and categories |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListOAuthConfigsResponse](../../models/operations/listoauthconfigsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListOAuthConfigsByType

List the organization's OAuth configurations for one connector type.<br><br>
<b>What comes back depends on the caller:</b><br>
<ul>
<li>An admin gets each stored record in full, including `config` with the client secret in clear text</li>
<li>Any other caller gets only the credential-free fields</li>
</ul>
A connector type with no configurations, or one that does not exist, yields an empty list.


### Example Usage

<!-- UsageSnippet language="go" operationID="listOAuthConfigsByType" method="get" path="/oauth/{connectorType}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.OAuthConfiguration.ListOAuthConfigsByType(ctx, operations.ListOAuthConfigsByTypeRequest{
        ConnectorType: "Confluence",
        Scope: components.ConnectorScopeTeam.ToPointer(),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListOAuthConfigsByTypeRequest](../../models/operations/listoauthconfigsbytyperequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListOAuthConfigsByTypeResponse](../../models/operations/listoauthconfigsbytyperesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CreateOAuthConfig

Create a new OAuth configuration for a connector type.<br><br>
<b>Admin Only:</b><br>
Only admins can create OAuth configurations. These provide the
OAuth credentials needed for users to authenticate connectors.<br><br>
<b>Use Case:</b><br>
Before users can create Google Drive connectors with OAuth,
an admin must create an OAuth configuration with the Google
OAuth client ID and secret.<br><br>
The response carries the credential-free view of the new configuration; the
credentials are never echoed back. Success is reported as `200`, not `201`.<br><br>
<b>The connector type is not checked.</b> `connectorType` is not compared with the OAuth
registry, so a misspelled or wrongly-cased name (`confluence` for `Confluence`) does not
fail: the configuration is stored under that name with the default icon and an empty
group, description and category list. It is then returned by
`GET /oauth/{connectorType}` for that same name but never by `GET /oauth`, which only
walks the registered types.<br><br>
Body fields other than `oauthInstanceName`, `config` and `baseUrl` are ignored.


### Example Usage

<!-- UsageSnippet language="go" operationID="createOAuthConfig" method="post" path="/oauth/{connectorType}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.OAuthConfiguration.CreateOAuthConfig(ctx, "Confluence", components.CreateOAuthConfigRequest{
        OauthInstanceName: "Production Google OAuth Credentials",
        Config: components.CreateOAuthConfigRequestConfig{
            ClientID: pipeshub.Pointer("123456789-abc.apps.googleusercontent.com"),
            ClientSecret: pipeshub.Pointer("GOCSPX-xxxxxxxxxxxxx"),
            TenantID: pipeshub.Pointer("12345678-1234-1234-1234-123456789abc"),
        },
        BaseURL: pipeshub.Pointer("https://gitlab.mycompany.com"),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `connectorType`                                                                            | `string`                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        | Confluence                                                                                 |
| `body`                                                                                     | [components.CreateOAuthConfigRequest](../../models/components/createoauthconfigrequest.md) | :heavy_check_mark:                                                                         | Request payload                                                                            |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.CreateOAuthConfigResponse](../../models/operations/createoauthconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 409      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetOAuthConfig

Get a specific OAuth configuration by ID.<br><br>
An admin gets `config` with the stored credentials, client secret in clear text.
Any other caller gets the same object without `config`.


### Example Usage

<!-- UsageSnippet language="go" operationID="getOAuthConfig" method="get" path="/oauth/{connectorType}/{configId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.OAuthConfiguration.GetOAuthConfig(ctx, "Confluence", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                       | Example                                                                                                                                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                | The context to use for the request.                                                                                                                                                                                                               |                                                                                                                                                                                                                                                   |
| `connectorType`                                                                                                                                                                                                                                   | `string`                                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                | N/A                                                                                                                                                                                                                                               | Confluence                                                                                                                                                                                                                                        |
| `configID`                                                                                                                                                                                                                                        | `string`                                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                | N/A                                                                                                                                                                                                                                               |                                                                                                                                                                                                                                                   |
| `reveal`                                                                                                                                                                                                                                          | [*components.RevealStoredSecrets](../../models/components/revealstoredsecrets.md)                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                | Set to `true` to receive stored secrets instead of placeholders. Honoured only for an<br/>admin's own configuration where the deployment allows it (see `GET<br/>/configurationManager/secretReveal`); ignored for OAuth-app tokens and any other value.<br/> |                                                                                                                                                                                                                                                   |
| `opts`                                                                                                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                | The options for this request.                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                   |

### Response

**[*operations.GetOAuthConfigResponse](../../models/operations/getoauthconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateOAuthConfig

Update an OAuth configuration.<br><br>
<b>Admin Only:</b><br>
Only an admin can update, including apps another admin created.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateOAuthConfig" method="put" path="/oauth/{connectorType}/{configId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.OAuthConfiguration.UpdateOAuthConfig(ctx, "Confluence", "<id>", operations.CreateUpdateOAuthConfigRequestBodyRequestBody1(
        operations.RequestBody1{
            OauthInstanceName: "<value>",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 | Example                                                                                                     |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |                                                                                                             |
| `connectorType`                                                                                             | `string`                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         | Confluence                                                                                                  |
| `configID`                                                                                                  | `string`                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |                                                                                                             |
| `body`                                                                                                      | [operations.UpdateOAuthConfigRequestBody](../../models/operations/updateoauthconfigrequestbody.md)          | :heavy_check_mark:                                                                                          | Send `oauthInstanceName`, `config`, or both; a body with neither is<br/>refused, even if it carries `baseUrl`.<br/> |                                                                                                             |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |                                                                                                             |

### Response

**[*operations.UpdateOAuthConfigResponse](../../models/operations/updateoauthconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteOAuthConfig

Delete an OAuth configuration.<br><br>
<b>Admin Only:</b><br>
Only admins can delete OAuth configurations.<br><br>
<b>Warning:</b><br>
The server does not check whether connectors still use the configuration:
it is deleted either way, and those connectors lose their OAuth app.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteOAuthConfig" method="delete" path="/oauth/{connectorType}/{configId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.OAuthConfiguration.DeleteOAuthConfig(ctx, "Confluence", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `connectorType`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | Confluence                                               |
| `configID`                                               | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteOAuthConfigResponse](../../models/operations/deleteoauthconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |