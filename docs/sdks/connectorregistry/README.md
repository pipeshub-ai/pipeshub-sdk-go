# ConnectorRegistry

## Overview

### Available Operations

* [List](#list) - List available connector types
* [GetSchemaForType](#getschemafortype) - Get connector configuration schema

## List

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
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.ConnectorRegistry.List(ctx, components.ConnectorScopeTeam.ToPointer(), pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |                                                                         |
| `scope`                                                                 | [*components.ConnectorScope](../../models/components/connectorscope.md) | :heavy_minus_sign:                                                      | Filter by scope type                                                    | team                                                                    |
| `page`                                                                  | **int64*                                                                | :heavy_minus_sign:                                                      | Page number                                                             |                                                                         |
| `limit`                                                                 | **int64*                                                                | :heavy_minus_sign:                                                      | Items per page                                                          |                                                                         |
| `search`                                                                | **string*                                                               | :heavy_minus_sign:                                                      | Search term for connector names                                         |                                                                         |
| `opts`                                                                  | [][operations.Option](../../models/operations/option.md)                | :heavy_minus_sign:                                                      | The options for this request.                                           |                                                                         |

### Response

**[*operations.GetConnectorRegistryResponse](../../models/operations/getconnectorregistryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSchemaForType

Get the configuration schema for a specific connector type.<br><br>
<b>Overview:</b><br>
Returns JSON Schema definitions for authentication, sync settings, and
filter options. Use this to dynamically build configuration forms.<br><br>
<b>Schema Sections:</b><br>
<ul>
<li><b>authSchema:</b> Fields for authentication (credentials, tokens)</li>
<li><b>syncSchema:</b> Sync settings (schedule, incremental options)</li>
<li><b>filterSchema:</b> Filter field definitions</li>
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

    res, err := s.ConnectorRegistry.GetSchemaForType(ctx, "google-drive")
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
| `connectorType`                                          | *string*                                                 | :heavy_check_mark:                                       | Connector type identifier                                | google-drive                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetConnectorSchemaResponse](../../models/operations/getconnectorschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |