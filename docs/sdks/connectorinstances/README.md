# ConnectorInstances

## Overview

Create, manage, and delete connector instances for your organization

### Available Operations

* [ListConnectorInstances](#listconnectorinstances) - List connector instances
* [CreateConnectorInstance](#createconnectorinstance) - Create connector instance
* [GetConnectorInstance](#getconnectorinstance) - Get connector instance
* [DeleteConnectorInstance](#deleteconnectorinstance) - Delete connector instance
* [GetConnectorStats](#getconnectorstats) - Get connector statistics
* [ResyncConnector](#resyncconnector) - Resync connector
* [StopConnectorSync](#stopconnectorsync) - Stop the running sync for a connector

## ListConnectorInstances

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
<li><b>desktopOnline:</b> Owner device's desktop app connected (Local FS only)</li>
<li><b>ownerDeviceId / ownerDeviceName:</b> Desktop device that owns the connector, set on first enable (Local FS only)</li>
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

    res, err := s.ConnectorInstances.ListConnectorInstances(ctx, operations.ListConnectorInstancesRequest{})
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
| `request`                                                                                            | [operations.ListConnectorInstancesRequest](../../models/operations/listconnectorinstancesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListConnectorInstancesResponse](../../models/operations/listconnectorinstancesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CreateConnectorInstance

Create a connector instance from a registry type.<br><br>
<b>Body:</b><br>
Put the auth fields the registry schema lists for <code>authType</code> in <code>config.auth</code>,
flat, and the connector's own sync settings in <code>config.sync</code>. A connector that needs no
OAuth sign-in is configured on creation.<br><br>
<b>OAuth:</b><br>
An admin may send the OAuth app's credentials in <code>config.auth</code>: the app is saved and
linked. Anyone may send <code>config.auth.oauthConfigId</code> to use a saved app; non-admins must.
Send <code>baseUrl</code>: the provider redirects to <code>{baseUrl}/{redirectUri}</code>.<br><br>
<b>Scope Permissions:</b><br>
<ul>
<li><code>team</code> scope requires admin privileges</li>
<li><code>personal</code> scope available to all users</li>
</ul>
<b>Next Steps After Creation:</b><br>
<ol>
<li>OAuth only: <code>GET /{id}/oauth/authorize</code>, then <code>GET /oauth/callback</code></li>
<li>Optionally set sync and filters via <code>PUT /{id}/config/filters-sync</code></li>
<li>Turn sync on via <code>POST /{id}/toggle</code></li>
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

    res, err := s.ConnectorInstances.CreateConnectorInstance(ctx, components.CreateConnectorRequest{
        ConnectorType: "confluence",
        InstanceName: "My Confluence",
        Scope: components.ConnectorScopePersonal,
        AuthType: components.ConnectorAuthTypeAPIToken.ToPointer(),
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

    res, err := s.ConnectorInstances.CreateConnectorInstance(ctx, components.CreateConnectorRequest{
        ConnectorType: "google-drive",
        InstanceName: "Company Google Drive",
        Scope: components.ConnectorScopeTeam,
        AuthType: components.ConnectorAuthTypeOauthAdminConsent.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: jiraNewOAuthApp

<!-- UsageSnippet language="go" operationID="createConnectorInstance" method="post" path="/connectors" example="jiraNewOAuthApp" -->
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

    res, err := s.ConnectorInstances.CreateConnectorInstance(ctx, components.CreateConnectorRequest{
        ConnectorType: "Jira",
        InstanceName: "Engineering Jira",
        Scope: components.ConnectorScopeTeam,
        AuthType: components.ConnectorAuthTypeOauth.ToPointer(),
        Config: &components.CreateConnectorRequestConfig{
            Auth: &components.ConnectorAuthConfig{
                OauthInstanceName: pipeshub.Pointer("Atlassian app"),
                AdditionalProperties: map[string]any{
                    "clientId": "your-atlassian-client-id",
                    "clientSecret": "your-atlassian-client-secret",
                },
            },
        },
        BaseURL: pipeshub.Pointer("https://pipeshub.example.com"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: jiraSavedOAuthApp

<!-- UsageSnippet language="go" operationID="createConnectorInstance" method="post" path="/connectors" example="jiraSavedOAuthApp" -->
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

    res, err := s.ConnectorInstances.CreateConnectorInstance(ctx, components.CreateConnectorRequest{
        ConnectorType: "Jira",
        InstanceName: "Engineering Jira",
        Scope: components.ConnectorScopeTeam,
        AuthType: components.ConnectorAuthTypeOauth.ToPointer(),
        Config: &components.CreateConnectorRequestConfig{
            Auth: &components.ConnectorAuthConfig{
                OauthConfigID: pipeshub.Pointer("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
            },
        },
        BaseURL: pipeshub.Pointer("https://pipeshub.example.com"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: postgresBasicAuth

<!-- UsageSnippet language="go" operationID="createConnectorInstance" method="post" path="/connectors" example="postgresBasicAuth" -->
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

    res, err := s.ConnectorInstances.CreateConnectorInstance(ctx, components.CreateConnectorRequest{
        ConnectorType: "PostgreSQL",
        InstanceName: "Analytics database",
        Scope: components.ConnectorScopeTeam,
        AuthType: components.ConnectorAuthTypeBasicAuth.ToPointer(),
        Config: &components.CreateConnectorRequestConfig{
            Auth: &components.ConnectorAuthConfig{
                AdditionalProperties: map[string]any{
                    "host": "db.example.com",
                    "port": 5432,
                    "database": "analytics",
                    "username": "reader",
                    "password": "your-password",
                },
            },
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
### Example Usage: rss

<!-- UsageSnippet language="go" operationID="createConnectorInstance" method="post" path="/connectors" example="rss" -->
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

    res, err := s.ConnectorInstances.CreateConnectorInstance(ctx, components.CreateConnectorRequest{
        ConnectorType: "RSS",
        InstanceName: "Company blog",
        Scope: components.ConnectorScopeTeam,
        Config: &components.CreateConnectorRequestConfig{
            Sync: &components.ConnectorSyncConfig{
                AdditionalProperties: map[string]any{
                    "feed_urls": "https://example.com/feed.xml",
                    "max_articles_per_feed": 50,
                },
            },
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.CreateConnectorRequest](../../models/components/createconnectorrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateConnectorInstanceResponse](../../models/operations/createconnectorinstanceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetConnectorInstance

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

    res, err := s.ConnectorInstances.GetConnectorInstance(ctx, "conn_abc123")
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
| `connectorID`                                            | `string`                                                 | :heavy_check_mark:                                       | Connector instance ID                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetConnectorInstanceResponse](../../models/operations/getconnectorinstanceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteConnectorInstance

Delete a connector instance and everything synced from it. The request
answers `202` as soon as the deletion has started; the deletion itself
runs in the background.<br><br>
<b>Warning:</b><br>
Deletion is permanent. The connector's records, record groups,
permissions and relations are removed from the graph database, its
embeddings are removed from the vector store, its stored configuration
and credentials are deleted, and its stored files are deleted (content
another connector still uses is kept). The records don't go to the
trash and can't be restored. Its sync schedule is removed too.<br><br>
<b>Permissions:</b><br>
The connector's creator or an admin, for either scope. An admin can
delete a member's personal connector.


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

    res, err := s.ConnectorInstances.DeleteConnectorInstance(ctx, "<id>")
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
| `connectorID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteConnectorInstanceResponse](../../models/operations/deleteconnectorinstanceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404, 409      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetConnectorStats

Retrieve statistics for a specific connector instance including record counts,
indexing status breakdown, and sync information.


### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectorStats" method="get" path="/connectors/{connectorId}/stats" -->
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

    res, err := s.ConnectorInstances.GetConnectorStats(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorStats != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `connectorID`                                            | `string`                                                 | :heavy_check_mark:                                       | Connector instance ID                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetConnectorStatsResponse](../../models/operations/getconnectorstatsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ResyncConnector

Trigger a resync of records from a connector instance.

**Overview:**

Fetches content from the external source and updates local records. Use when you suspect data is out of sync.

**Warning:** This can be resource-intensive for large connectors.


### Example Usage

<!-- UsageSnippet language="go" operationID="resyncConnector" method="post" path="/connectors/{connectorId}/resync" -->
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

    res, err := s.ConnectorInstances.ResyncConnector(ctx, "<id>", components.ResyncConnectorRequestBody{
        ConnectorName: "GOOGLE_DRIVE",
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
| `connectorID`                                                                                  | `string`                                                                                       | :heavy_check_mark:                                                                             | Connector instance ID                                                                          |
| `body`                                                                                         | [components.ResyncConnectorRequestBody](../../models/components/resyncconnectorrequestbody.md) | :heavy_check_mark:                                                                             | Request body for resync connector                                                              |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ResyncConnectorResponse](../../models/operations/resyncconnectorresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorResponse           | 400, 401, 403                     | application/json                  |
| apierrors.ResyncConnectorConflict | 409                               | application/json                  |
| apierrors.ErrorResponse           | 500, 503                          | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## StopConnectorSync

Ask the in-flight sync for a connector instance to stop.

**Overview:**

Returns as soon as the request is recorded; the sync halts at its next
interruptible point, which for a connector inside a long provider call
is when that call returns. `stopped` reports whether a sync was actually
running, and `false` is a normal answer rather than a failure.

Deliberately not refused while the connector is busy or locked: this is
the endpoint that has to work precisely then, including against a lock a
crashed process left behind. When nothing is running it also repairs a
stale SYNCING or FULL_SYNCING status, unless the service cannot tell a
live sync on another worker from a stale one, in which case it changes
nothing and says so.


### Example Usage

<!-- UsageSnippet language="go" operationID="stopConnectorSync" method="post" path="/connectors/{connectorId}/sync/stop" -->
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

    res, err := s.ConnectorInstances.StopConnectorSync(ctx, "<id>")
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
| `connectorID`                                            | `string`                                                 | :heavy_check_mark:                                       | Connector instance ID                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.StopConnectorSyncResponse](../../models/operations/stopconnectorsyncresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |