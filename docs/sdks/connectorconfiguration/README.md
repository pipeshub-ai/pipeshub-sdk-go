# ConnectorConfiguration

## Overview

### Available Operations

* [Get](#get) - Get connector configuration
* [Update](#update) - Update connector configuration
* [UpdateAuth](#updateauth) - Update authentication configuration
* [UpdateFiltersSync](#updatefilterssync) - Update filters and sync configuration

## Get

Get the current configuration for a connector instance.<br><br>
<b>Security:</b><br>
Sensitive data (credentials, OAuth tokens) are redacted from the response.
Only admins can see partial credential information.


### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectorConfig" method="get" path="/connectors/{connectorId}/config" -->
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

    res, err := s.ConnectorConfiguration.Get(ctx, "<id>")
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

**[*operations.GetConnectorConfigResponse](../../models/operations/getconnectorconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update authentication, sync, and filter configuration.<br><br>
<b>Prerequisites:</b><br>
Connector must be <b>disabled</b> before updating configuration.
Disable it first using <code>POST /{id}/toggle</code>.<br><br>
<b>Partial Updates:</b><br>
Only provide the sections you want to update. Omitted sections
are not modified.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateConnectorConfig" method="put" path="/connectors/{connectorId}/config" -->
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

    res, err := s.ConnectorConfiguration.Update(ctx, "<id>", components.UpdateConnectorConfigRequest{
        Auth: &components.ConnectorAuthConfig{
            Values: map[string]any{
                "apiKey": "sk-xxxxx",
                "baseUrl": "https://api.example.com",
            },
            OauthConfigID: pipeshub.Pointer("oauth_config_123"),
        },
        Sync: &components.ConnectorSyncConfig{
            ScheduledConfig: &components.ScheduledConfig{
                CronExpression: pipeshub.Pointer("0 */6 * * *"),
                Timezone: pipeshub.Pointer("America/New_York"),
            },
            WebhookConfig: &components.WebhookConfig{
                Events: []string{
                    "file.created",
                    "file.modified",
                    "file.deleted",
                },
            },
        },
        Filters: &components.ConnectorFiltersConfig{
            Sync: &components.ConnectorFiltersConfigSync{
                Values: map[string]any{
                    "folders": []any{
                        "folder_id_1",
                        "folder_id_2",
                    },
                    "fileTypes": []any{
                        "pdf",
                        "docx",
                        "xlsx",
                    },
                    "includeShared": true,
                },
            },
        },
        BaseURL: pipeshub.Pointer("https://confluence.mycompany.com"),
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `connectorID`                                                                                      | *string*                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `body`                                                                                             | [components.UpdateConnectorConfigRequest](../../models/components/updateconnectorconfigrequest.md) | :heavy_check_mark:                                                                                 | Request payload                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateConnectorConfigResponse](../../models/operations/updateconnectorconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateAuth

Update only the authentication configuration.<br><br>
<b>Use Case:</b><br>
Use this when you need to update credentials without changing
sync or filter settings. Useful for credential rotation.<br><br>
<b>Prerequisites:</b><br>
Connector must be disabled. This endpoint clears OAuth state,
requiring re-authentication for OAuth connectors.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateConnectorAuthConfig" method="put" path="/connectors/{connectorId}/config/auth" -->
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

    res, err := s.ConnectorConfiguration.UpdateAuth(ctx, "<id>", components.UpdateConnectorAuthRequest{
        Auth: components.ConnectorAuthConfig{
            Values: map[string]any{
                "apiKey": "sk-xxxxx",
                "baseUrl": "https://api.example.com",
            },
            OauthConfigID: pipeshub.Pointer("oauth_config_123"),
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `connectorID`                                                                                  | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `body`                                                                                         | [components.UpdateConnectorAuthRequest](../../models/components/updateconnectorauthrequest.md) | :heavy_check_mark:                                                                             | Request payload                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateConnectorAuthConfigResponse](../../models/operations/updateconnectorauthconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateFiltersSync

Update filter selections and sync settings without touching auth.<br><br>
<b>Use Case:</b><br>
Use this to change what data is synced or adjust sync schedule
without re-authenticating.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateConnectorFiltersSyncConfig" method="put" path="/connectors/{connectorId}/config/filters-sync" -->
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

    res, err := s.ConnectorConfiguration.UpdateFiltersSync(ctx, "<id>", components.UpdateConnectorFiltersSyncRequest{
        Sync: &components.ConnectorSyncConfig{
            ScheduledConfig: &components.ScheduledConfig{
                CronExpression: pipeshub.Pointer("0 */6 * * *"),
                Timezone: pipeshub.Pointer("America/New_York"),
            },
            WebhookConfig: &components.WebhookConfig{
                Events: []string{
                    "file.created",
                    "file.modified",
                    "file.deleted",
                },
            },
        },
        Filters: &components.ConnectorFiltersConfig{
            Sync: &components.ConnectorFiltersConfigSync{
                Values: map[string]any{
                    "folders": []any{
                        "folder_id_1",
                        "folder_id_2",
                    },
                    "fileTypes": []any{
                        "pdf",
                        "docx",
                        "xlsx",
                    },
                    "includeShared": true,
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `connectorID`                                                                                                | *string*                                                                                                     | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `body`                                                                                                       | [components.UpdateConnectorFiltersSyncRequest](../../models/components/updateconnectorfilterssyncrequest.md) | :heavy_check_mark:                                                                                           | Request payload                                                                                              |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateConnectorFiltersSyncConfigResponse](../../models/operations/updateconnectorfilterssyncconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |