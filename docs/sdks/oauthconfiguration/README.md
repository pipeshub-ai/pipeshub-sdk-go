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

Get all connector types that support OAuth authentication.<br><br>
<b>Admin Use:</b><br>
Admins use this to see which connector types need OAuth credentials
to be configured before users can authenticate.


### Example Usage

<!-- UsageSnippet language="go" operationID="getOAuthRegistry" method="get" path="/oauth/registry" -->
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
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetOAuthConnectorType

Get details for a specific OAuth-capable connector type.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOAuthConnectorType" method="get" path="/oauth/registry/{connectorType}" -->
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

    res, err := s.OAuthConfiguration.GetOAuthConnectorType(ctx, "<value>")
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
| `connectorType`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOAuthConnectorTypeResponse](../../models/operations/getoauthconnectortyperesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404           | application/json        |
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `page`                                                   | `*int64`                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `limit`                                                  | `*int64`                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `search`                                                 | `*string`                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOAuthConfigsResponse](../../models/operations/listoauthconfigsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListOAuthConfigsByType

Get all OAuth configurations for a specific connector type.

### Example Usage

<!-- UsageSnippet language="go" operationID="listOAuthConfigsByType" method="get" path="/oauth/{connectorType}" -->
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

    res, err := s.OAuthConfiguration.ListOAuthConfigsByType(ctx, operations.ListOAuthConfigsByTypeRequest{
        ConnectorType: "<value>",
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
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CreateOAuthConfig

Create a new OAuth configuration for a connector type.<br><br>
<b>Admin Only:</b><br>
Only admins can create OAuth configurations. These provide the
OAuth credentials needed for users to authenticate connectors.<br><br>
<b>Use Case:</b><br>
Before users can create Google Drive connectors with OAuth,
an admin must create an OAuth configuration with the Google
OAuth client ID and secret.


### Example Usage

<!-- UsageSnippet language="go" operationID="createOAuthConfig" method="post" path="/oauth/{connectorType}" -->
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

    res, err := s.OAuthConfiguration.CreateOAuthConfig(ctx, "<value>", components.CreateOAuthConfigRequest{
        OauthInstanceName: "Production Google OAuth Credentials",
        Config: components.CreateOAuthConfigRequestConfig{
            ClientID: "123456789-abc.apps.googleusercontent.com",
            ClientSecret: "GOCSPX-xxxxxxxxxxxxx",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `connectorType`                                                                            | `string`                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `body`                                                                                     | [components.CreateOAuthConfigRequest](../../models/components/createoauthconfigrequest.md) | :heavy_check_mark:                                                                         | Request payload                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateOAuthConfigResponse](../../models/operations/createoauthconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 409      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetOAuthConfig

Get a specific OAuth configuration by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOAuthConfig" method="get" path="/oauth/{connectorType}/{configId}" -->
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

    res, err := s.OAuthConfiguration.GetOAuthConfig(ctx, "<value>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                | The context to use for the request.                                                                                                                                                                                                               |
| `connectorType`                                                                                                                                                                                                                                   | `string`                                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                | N/A                                                                                                                                                                                                                                               |
| `configID`                                                                                                                                                                                                                                        | `string`                                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                | N/A                                                                                                                                                                                                                                               |
| `reveal`                                                                                                                                                                                                                                          | [*components.RevealStoredSecrets](../../models/components/revealstoredsecrets.md)                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                | Set to `true` to receive stored secrets instead of placeholders. Honoured only for an<br/>admin's own configuration where the deployment allows it (see `GET<br/>/configurationManager/secretReveal`); ignored for OAuth-app tokens and any other value.<br/> |
| `opts`                                                                                                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                | The options for this request.                                                                                                                                                                                                                     |

### Response

**[*operations.GetOAuthConfigResponse](../../models/operations/getoauthconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
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

    res, err := s.OAuthConfiguration.UpdateOAuthConfig(ctx, "<value>", "<id>", operations.CreateUpdateOAuthConfigRequestBodyRequestBody1(
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

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `connectorType`                                                                                             | `string`                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `configID`                                                                                                  | `string`                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `body`                                                                                                      | [operations.UpdateOAuthConfigRequestBody](../../models/operations/updateoauthconfigrequestbody.md)          | :heavy_check_mark:                                                                                          | Send `oauthInstanceName`, `config`, or both; a body with neither is<br/>refused, even if it carries `baseUrl`.<br/> |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

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

    res, err := s.OAuthConfiguration.DeleteOAuthConfig(ctx, "<value>", "<id>")
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
| `connectorType`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `configID`                                               | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOAuthConfigResponse](../../models/operations/deleteoauthconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |