# ConnectorControl

## Overview

Enable/disable connectors for sync or agent functionality

### Available Operations

* [ToggleConnector](#toggleconnector) - Toggle connector sync or agent

## ToggleConnector

Flip a connector's sync or agent use: each call turns it on when it is off and
off when it is on. Read <code>isActive</code> / <code>isAgentActive</code> first. Turning sync on
starts the first sync.<br><br>
<b>Toggle Types:</b><br>
<ul>
<li><code>sync</code> - data synchronization</li>
<li><code>agent</code> - use by AI agents</li>
</ul>
<b>Before Sync Can Be Turned On:</b><br>
<ul>
<li>The connector is configured (<code>isConfigured: true</code>)</li>
<li>OAuth connectors: sign-in has completed</li>
<li>The connection test against the source passes</li>
<li>For agent: the connector supports agents (<code>supportsAgent: true</code>)</li>
<li>Local FS: must be enabled from the owner device via the desktop app, which sends <code>deviceId</code> and <code>deviceName</code>. The first enable claims that device as owner.</li>
</ul>
<b>Permissions:</b><br>
<ul>
<li>Team scope: Requires admin</li>
<li>Personal scope: Only creator can toggle</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="toggleConnector" method="post" path="/connectors/{connectorId}/toggle" -->
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

    res, err := s.ConnectorControl.ToggleConnector(ctx, "<id>", components.ConnectorToggleRequest{
        Type: components.ConnectorToggleRequestTypeSync,
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
| `connectorID`                                                                          | `string`                                                                               | :heavy_check_mark:                                                                     | Connector instance ID                                                                  |
| `body`                                                                                 | [components.ConnectorToggleRequest](../../models/components/connectortogglerequest.md) | :heavy_check_mark:                                                                     | Request payload                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ToggleConnectorResponse](../../models/operations/toggleconnectorresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorResponse           | 400, 401, 403, 404                | application/json                  |
| apierrors.ToggleConnectorConflict | 409                               | application/json                  |
| apierrors.ErrorResponse           | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |