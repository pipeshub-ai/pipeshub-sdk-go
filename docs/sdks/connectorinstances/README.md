# ConnectorInstances

## Overview

### Available Operations

* [List](#list) - List connector instances
* [Create](#create) - Create connector instance
* [ListActive](#listactive) - List active connector instances
* [ListInactive](#listinactive) - List inactive connector instances
* [ListConfigured](#listconfigured) - List configured connector instances
* [ListActiveAgents](#listactiveagents) - List active agent connectors
* [Get](#get) - Get connector instance
* [Delete](#delete) - Delete connector instance
* [UpdateName](#updatename) - Update connector instance name

## List

Get all configured connector instances for your organization.<br><br>
<b>Overview:</b><br>
Returns instances created by users, filtered by scope and permissions.
Team-scope connectors are visible to all org users. Personal connectors
are only visible to their creators.<br><br>
<b>Instance States:</b><br>
<ul>
<li><b>isConfigured:</b> All required settings are complete</li>
<li><b>isAuthenticated:</b> OAuth flow complete or credentials valid</li>
<li><b>isActive:</b> Connector is enabled for sync/agent</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="listConnectorInstances" method="get" path="/connectors" -->
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

    res, err := s.ConnectorInstances.List(ctx, components.ConnectorScopeTeam.ToPointer(), pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20), nil)
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
| `scope`                                                                 | [*components.ConnectorScope](../../models/components/connectorscope.md) | :heavy_minus_sign:                                                      | Filter by scope (team or personal)                                      | team                                                                    |
| `page`                                                                  | **int64*                                                                | :heavy_minus_sign:                                                      | N/A                                                                     |                                                                         |
| `limit`                                                                 | **int64*                                                                | :heavy_minus_sign:                                                      | N/A                                                                     |                                                                         |
| `search`                                                                | **string*                                                               | :heavy_minus_sign:                                                      | Search by instance name                                                 |                                                                         |
| `opts`                                                                  | [][operations.Option](../../models/operations/option.md)                | :heavy_minus_sign:                                                      | The options for this request.                                           |                                                                         |

### Response

**[*operations.ListConnectorInstancesResponse](../../models/operations/listconnectorinstancesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new connector instance from a registry type.<br><br>
<b>Overview:</b><br>
Creates a new connector instance that can then be configured and enabled.
The instance is created in an unconfigured state and needs authentication
and filter setup before it can be activated.<br><br>
<b>Scope Permissions:</b><br>
<ul>
<li><code>team</code> scope requires admin privileges</li>
<li><code>personal</code> scope available to all users</li>
</ul>
<b>Next Steps After Creation:</b><br>
<ol>
<li>Configure authentication via <code>PUT /{id}/config/auth</code></li>
<li>Complete OAuth flow if needed via <code>GET /{id}/oauth/authorize</code></li>
<li>Set up filters via <code>POST /{id}/filters</code></li>
<li>Enable connector via <code>POST /{id}/toggle</code></li>
</ol>


### Example Usage: confluence

<!-- UsageSnippet language="go" operationID="createConnectorInstance" method="post" path="/connectors" example="confluence" -->
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

    res, err := s.ConnectorInstances.Create(ctx, components.CreateConnectorRequest{
        ConnectorType: "confluence",
        InstanceName: "My Confluence",
        Scope: components.ConnectorScopePersonal,
        AuthType: components.AuthTypeAPIToken.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: googleDrive

<!-- UsageSnippet language="go" operationID="createConnectorInstance" method="post" path="/connectors" example="googleDrive" -->
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

    res, err := s.ConnectorInstances.Create(ctx, components.CreateConnectorRequest{
        ConnectorType: "google-drive",
        InstanceName: "Company Google Drive",
        Scope: components.ConnectorScopeTeam,
        AuthType: components.AuthTypeOauthAdminConsent.ToPointer(),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.CreateConnectorRequest](../../models/components/createconnectorrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateConnectorInstanceResponse](../../models/operations/createconnectorinstanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListActive

Get all active (enabled) connector instances.<br><br>
<b>Overview:</b><br>
Returns only instances where <code>isActive: true</code>.
These are connectors currently syncing data or available to AI agents.


### Example Usage

<!-- UsageSnippet language="go" operationID="listActiveConnectors" method="get" path="/connectors/active" -->
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

    res, err := s.ConnectorInstances.ListActive(ctx)
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

**[*operations.ListActiveConnectorsResponse](../../models/operations/listactiveconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListInactive

Get all inactive (disabled) connector instances.

### Example Usage

<!-- UsageSnippet language="go" operationID="listInactiveConnectors" method="get" path="/connectors/inactive" -->
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

    res, err := s.ConnectorInstances.ListInactive(ctx)
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

**[*operations.ListInactiveConnectorsResponse](../../models/operations/listinactiveconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListConfigured

Get all connector instances that have completed configuration.<br><br>
<b>Overview:</b><br>
Returns instances where <code>isConfigured: true</code>.
These have all required settings but may not be active yet.


### Example Usage

<!-- UsageSnippet language="go" operationID="listConfiguredConnectors" method="get" path="/connectors/configured" -->
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

    res, err := s.ConnectorInstances.ListConfigured(ctx, components.ConnectorScopeTeam.ToPointer(), pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                     | Example                                                                                                                                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                                             |                                                                                                                                                                                                                                                 |
| `scope`                                                                                                                                                                                                                                         | [*components.ConnectorScope](../../models/components/connectorscope.md)                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                              | Scope determines visibility and access control for connectors:<br><br/><ul><br/><li><code>team</code> - Available to all users in the organization (admin-only creation)</li><br/><li><code>personal</code> - Private to the creating user only</li><br/></ul><br/> | team                                                                                                                                                                                                                                            |
| `page`                                                                                                                                                                                                                                          | **int64*                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                              | N/A                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                 |
| `limit`                                                                                                                                                                                                                                         | **int64*                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                              | N/A                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                 |
| `search`                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                              | N/A                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                 |

### Response

**[*operations.ListConfiguredConnectorsResponse](../../models/operations/listconfiguredconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListActiveAgents

Get connector instances enabled for AI agent integration.<br><br>
<b>Overview:</b><br>
Returns connectors where <code>agentEnabled: true</code>.
These are available to AI agents for querying and actions.


### Example Usage

<!-- UsageSnippet language="go" operationID="listActiveAgentConnectors" method="get" path="/connectors/agents/active" -->
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

    res, err := s.ConnectorInstances.ListActiveAgents(ctx, components.ConnectorScopeTeam.ToPointer(), pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                     | Example                                                                                                                                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                                             |                                                                                                                                                                                                                                                 |
| `scope`                                                                                                                                                                                                                                         | [*components.ConnectorScope](../../models/components/connectorscope.md)                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                              | Scope determines visibility and access control for connectors:<br><br/><ul><br/><li><code>team</code> - Available to all users in the organization (admin-only creation)</li><br/><li><code>personal</code> - Private to the creating user only</li><br/></ul><br/> | team                                                                                                                                                                                                                                            |
| `page`                                                                                                                                                                                                                                          | **int64*                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                              | N/A                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                 |
| `limit`                                                                                                                                                                                                                                         | **int64*                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                              | N/A                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                 |
| `search`                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                              | N/A                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                 |

### Response

**[*operations.ListActiveAgentConnectorsResponse](../../models/operations/listactiveagentconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a specific connector instance by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectorInstance" method="get" path="/connectors/{connectorId}" -->
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

    res, err := s.ConnectorInstances.Get(ctx, "conn_abc123")
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
| `connectorID`                                            | *string*                                                 | :heavy_check_mark:                                       | Connector instance ID                                    | conn_abc123                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetConnectorInstanceResponse](../../models/operations/getconnectorinstanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a connector instance and all associated data.<br><br>
<b>Warning:</b><br>
This permanently removes the connector configuration.
Synced records in knowledge bases are NOT deleted.<br><br>
<b>Permissions:</b><br>
<ul>
<li>Team scope: Requires admin</li>
<li>Personal scope: Only creator can delete</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteConnectorInstance" method="delete" path="/connectors/{connectorId}" -->
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

    res, err := s.ConnectorInstances.Delete(ctx, "<id>")
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
| `connectorID`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteConnectorInstanceResponse](../../models/operations/deleteconnectorinstanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateName

Update the display name of a connector instance.<br><br>
<b>Note:</b> This only updates the display name, not the connector configuration.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateConnectorName" method="put" path="/connectors/{connectorId}/name" -->
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

    res, err := s.ConnectorInstances.UpdateName(ctx, "<id>", components.UpdateConnectorNameRequest{
        InstanceName: "Sales Team Drive (Updated)",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `connectorID`                                                                                  | *string*                                                                                       | :heavy_check_mark:                                                                             | Unique connector instance ID                                                                   |
| `body`                                                                                         | [components.UpdateConnectorNameRequest](../../models/components/updateconnectornamerequest.md) | :heavy_check_mark:                                                                             | Request body for Update connector instance name                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateConnectorNameResponse](../../models/operations/updateconnectornameresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |