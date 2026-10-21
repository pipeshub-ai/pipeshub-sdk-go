# ConnectorRegistry

## Overview

Browse available connector types and their configuration schemas

### Available Operations

* [GetConnectorRegistry](#getconnectorregistry) - List available connector types
* [GetConnectorSchema](#getconnectorschema) - Get connector configuration schema
* [TestConnectorConnection](#testconnectorconnection) - Test connector credentials before saving
* [GetConnectorEgressIps](#getconnectoregressips) - Get the IP addresses connectors connect from

## GetConnectorRegistry

Get all available connector types from the registry.<br><br>
<b>Overview:</b><br>
The registry contains all connector types that can be configured as instances.
Each type has specific authentication requirements, supported scopes, and capabilities.<br><br>
<b>Connector Types Include:</b><br>
<ul>
<li><b>Google Workspace:</b> Drive, Gmail, Calendar, etc.</li>
<li><b>Microsoft 365:</b> OneDrive, Outlook, SharePoint, etc.</li>
<li><b>Cloud Storage:</b> Dropbox, Box, AWS S3</li>
<li><b>Collaboration:</b> Slack, Confluence, Notion, Jira</li>
<li><b>Databases:</b> PostgreSQL, MySQL, MongoDB</li>
</ul>
<b>Filtering:</b><br>
Use <code>scope</code> to filter by team or personal connectors.
Use <code>search</code> for full-text search across connector names.


### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectorRegistry" method="get" path="/connectors/registry" -->
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

    res, err := s.ConnectorRegistry.GetConnectorRegistry(ctx, operations.ConnectorRegistryScopeTeam.ToPointer(), pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |                                                                                         |
| `scope`                                                                                 | [*operations.ConnectorRegistryScope](../../models/operations/connectorregistryscope.md) | :heavy_minus_sign:                                                                      | Which connector types to list. Defaults to `team`.                                      | team                                                                                    |
| `page`                                                                                  | `*int64`                                                                                | :heavy_minus_sign:                                                                      | Page number                                                                             |                                                                                         |
| `limit`                                                                                 | `*int64`                                                                                | :heavy_minus_sign:                                                                      | Items per page                                                                          |                                                                                         |
| `search`                                                                                | `*string`                                                                               | :heavy_minus_sign:                                                                      | Search term for connector names                                                         |                                                                                         |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |                                                                                         |

### Response

**[*operations.GetConnectorRegistryResponse](../../models/operations/getconnectorregistryresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetConnectorSchema

The configuration form for one connector type.<br><br>
<b>Sections:</b><br>
<ul>
<li><b>auth:</b> <code>supportedAuthTypes</code>, and per auth type in <code>schemas</code> the
fields to send in <code>config.auth</code> (by <code>name</code>) plus, for OAuth, the
<code>redirectUri</code> path</li>
<li><b>sync:</b> <code>supportedStrategies</code>, <code>scheduledConfig</code> defaults, and the
connector's own settings in <code>customFields</code> (sent in <code>config.sync</code> by
<code>name</code>)</li>
<li><b>filters:</b> the <code>sync</code> and <code>indexing</code> filter fields; a field with
<code>optionSourceType: dynamic</code> gets its values from
<code>GET /connectors/{connectorId}/filters/{filterKey}/options</code></li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectorSchema" method="get" path="/connectors/registry/{connectorType}/schema" -->
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

    res, err := s.ConnectorRegistry.GetConnectorSchema(ctx, "Jira")
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
| `connectorType`                                          | `string`                                                 | :heavy_check_mark:                                       | Connector type: `type` from GET /connectors/registry     | Jira                                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetConnectorSchemaResponse](../../models/operations/getconnectorschemaresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## TestConnectorConnection

Try the auth settings from a connector's setup form without creating or
changing an instance.<br><br>
<b>Overview:</b><br>
Only connector types whose schema has <code>supportsConnectionCheck: true</code>
can be tested. A failed check is not an error: the response is
<code>200</code> with <code>success: false</code> and a <code>message</code>
to show the user (for example a wrong password, an unknown database, or a
timeout because a firewall blocks PipesHub's IP address).<br><br>
<b>Editing an instance:</b><br>
Pass <code>connectorId</code> to lay the settings over the instance's saved
ones, as <code>PUT /connectors/{connectorId}/config/auth</code> does. The
caller must be allowed to update that instance.


### Example Usage

<!-- UsageSnippet language="go" operationID="testConnectorConnection" method="post" path="/connectors/registry/{connectorType}/test-connection" -->
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

    res, err := s.ConnectorRegistry.TestConnectorConnection(ctx, "PostgreSQL", operations.TestConnectorConnectionRequestBody{
        Auth: map[string]any{
            "host": "db.example.com",
            "port": "5432",
            "database": "analytics",
            "username": "reader",
            "password": "your-password",
        },
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `connectorType`                                                                                                | `string`                                                                                                       | :heavy_check_mark:                                                                                             | Connector type identifier                                                                                      | PostgreSQL                                                                                                     |
| `body`                                                                                                         | [operations.TestConnectorConnectionRequestBody](../../models/operations/testconnectorconnectionrequestbody.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.TestConnectorConnectionResponse](../../models/operations/testconnectorconnectionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetConnectorEgressIps

Public IP addresses PipesHub's connectors use to reach external systems, for
users to allow in a database's or API's firewall.<br><br>
Set by the deployment through <code>CONNECTOR_EGRESS_IPS</code>; otherwise
looked up by the connector service and cached. The list is empty when the
address cannot be determined, or when the deployment sets
<code>CONNECTOR_EGRESS_IPS=none</code>.


### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectorEgressIps" method="get" path="/connectors/network/egress-ips" -->
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

    res, err := s.ConnectorRegistry.GetConnectorEgressIps(ctx)
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

**[*operations.GetConnectorEgressIpsResponse](../../models/operations/getconnectoregressipsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403                | application/json        |
| apierrors.ErrorResponse | 503                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |