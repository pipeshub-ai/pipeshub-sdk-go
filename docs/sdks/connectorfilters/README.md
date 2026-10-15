# ConnectorFilters

## Overview

Dynamic filter options for selecting which data to sync

### Available Operations

* [GetConnectorFilters](#getconnectorfilters) - Get filter options
* [SaveConnectorFilters](#saveconnectorfilters) - Save filter selections
* [GetFilterFieldOptions](#getfilterfieldoptions) - Get dynamic filter options

## GetConnectorFilters

Get available filter options for a connector.<br><br>
<b>Overview:</b><br>
Returns filter fields that can be used to limit what data is synced.
For example, a Google Drive connector might offer filters for
specific folders or file types.<br><br>
<b>Dynamic Filters:</b><br>
Some filter fields have <code>dynamic: true</code> (i.e.
<code>optionSourceType: dynamic</code>), meaning their options are
loaded separately via the filter options endpoint.
For such fields, call
<code>GET /connectors/{connectorId}/filters/{filterKey}/options</code>.
When scoping options to a specific parent (e.g. GitLab projects within
a selected group), pass the parent paths via the
<code>contextGroupPath</code> query parameter on that endpoint.<br><br>
<b>Which connectors answer:</b><br>
Only connectors with auth type <code>API_TOKEN</code>, <code>USERNAME_PASSWORD</code>
or <code>OAUTH_ADMIN_CONSENT</code> (and saved settings), and <code>OAUTH</code>
connectors that have completed sign-in. Every other connector is refused with <code>400</code>:
those with auth type <code>NONE</code> (for example RSS), <code>BASIC_AUTH</code>,
<code>ACCESS_KEY</code>, <code>CONNECTION_STRING</code> or <code>CUSTOM</code>
(<code>Unsupported authentication type: ...</code>), an OAUTH connector that has not
signed in, and a token or password connector with no saved settings.
Their filter fields are in the registry schema; dynamic values come from
<code>GET /connectors/{connectorId}/filters/{filterKey}/options</code>.


### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectorFilters" method="get" path="/connectors/{connectorId}/filters" -->
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

    res, err := s.ConnectorFilters.GetConnectorFilters(ctx, "<id>")
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

**[*operations.GetConnectorFiltersResponse](../../models/operations/getconnectorfiltersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SaveConnectorFilters

Save the user's filter selections for a connector.<br><br>
<b>Overview:</b><br>
After viewing filter options, use this endpoint to save the
selected values. These determine what data will be synced.<br><br>
<b>Request:</b><br>
<code>filters</code> must hold at least one entry: <code>{"filters": {}}</code> is
refused with <code>400</code>. The server stores whatever it is given under the
connector's saved filter values without checking it against the filter schema.
Use <code>PUT /connectors/{connectorId}/config/filters-sync</code> to save typed
sync and indexing filters.


### Example Usage

<!-- UsageSnippet language="go" operationID="saveConnectorFilters" method="post" path="/connectors/{connectorId}/filters" -->
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

    res, err := s.ConnectorFilters.SaveConnectorFilters(ctx, "<id>", components.SaveConnectorFiltersRequest{
        Filters: map[string]any{
            "folders": []any{
                map[string]any{
                    "id": "folder_123",
                    "name": "Documents",
                },
                map[string]any{
                    "id": "folder_456",
                    "name": "Reports",
                },
            },
            "fileTypes": []any{
                "pdf",
                "docx",
            },
            "modifiedAfter": "2024-01-01",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `connectorID`                                                                                    | `string`                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `body`                                                                                           | [components.SaveConnectorFiltersRequest](../../models/components/saveconnectorfiltersrequest.md) | :heavy_check_mark:                                                                               | Request payload                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.SaveConnectorFiltersResponse](../../models/operations/saveconnectorfiltersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetFilterFieldOptions

Get options for a dynamic filter field with pagination.<br><br>
<b>Overview:</b><br>
For filters with <code>optionSourceType: dynamic</code>, options are loaded
from the connected service. This supports pagination and search.<br><br>
<b>Scoped options (GitLab):</b><br>
When fetching options for the <code>project_ids</code> filter key, pass
one or more <code>contextGroupPath</code> query parameters to restrict
the returned repositories to those that belong to the specified GitLab
group namespace paths (e.g. <code>myorg/team</code>). This is used to
populate the project filter after the user has selected groups.<br><br>
<b>Examples:</b><br>
<ul>
<li>Google Drive folders list</li>
<li>Slack channels list</li>
<li>Confluence spaces list</li>
<li>GitLab groups list (<code>filterKey=group_ids</code>)</li>
<li>GitLab projects scoped to groups (<code>filterKey=project_ids&amp;contextGroupPath=myorg/team</code>)</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getFilterFieldOptions" method="get" path="/connectors/{connectorId}/filters/{filterKey}/options" -->
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

    res, err := s.ConnectorFilters.GetFilterFieldOptions(ctx, operations.GetFilterFieldOptionsRequest{
        ConnectorID: "<id>",
        FilterKey: "buckets",
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
| `request`                                                                                          | [operations.GetFilterFieldOptionsRequest](../../models/operations/getfilterfieldoptionsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetFilterFieldOptionsResponse](../../models/operations/getfilterfieldoptionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |