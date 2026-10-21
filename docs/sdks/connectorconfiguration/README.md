# ConnectorConfiguration

## Overview

Configure authentication, sync settings, and filters for connectors

### Available Operations

* [GetConnectorConfig](#getconnectorconfig) - Get connector configuration
* [UpdateConnectorConfig](#updateconnectorconfig) - Update connector configuration
* [UpdateConnectorAuthConfig](#updateconnectorauthconfig) - Update authentication configuration
* [UpdateConnectorFiltersSyncConfig](#updateconnectorfilterssyncconfig) - Update filters and sync configuration

## GetConnectorConfig

Get the current configuration for a connector instance.<br><br>
<b>Security:</b><br>
OAuth tokens are never returned. Each stored secret in <code>config.auth</code>
is returned as <code>••••••••</code>: a field the connector type's auth schema
marks secret (such as an API token or password), a secret field of its OAuth
app registration (such as <code>clientSecret</code>), or the <code>private_key</code>
of an older Google Workspace service account.


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

    res, err := s.ConnectorConfiguration.GetConnectorConfig(ctx, "<id>")
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

**[*operations.GetConnectorConfigResponse](../../models/operations/getconnectorconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateConnectorConfig

Update authentication, sync, and filter configuration.<br><br>
<b>Prerequisites:</b><br>
Connector must be <b>disabled</b> before updating configuration.
Disable it first using <code>POST /{id}/toggle</code>.<br><br>
<b>Partial Updates:</b><br>
Only provide the sections you want to update. Omitted sections
are not modified.<br><br>
<b>Secrets:</b><br>
A secret field in <code>auth</code> sent as <code>••••••••</code> keeps the
stored value, so the configuration can be sent back as it was read. The
response masks stored secrets the same way.<br><br>
<b>Signing out:</b><br>
A change to <code>auth</code>, a different top-level <code>oauthConfigId</code>,
or a linked OAuth app that now holds a different client signs the connector
out and clears its OAuth tokens. Sending <code>auth</code> back as it was read,
with secrets as the mask or in full, keeps it signed in.<br><br>
<b>OAuth apps (OAUTH connectors):</b><br>
An OAUTH connector signs in with the client ID and secret of its linked OAuth
app, so client credentials an admin sends in <code>auth</code> go into an app,
as on <code>PUT /{connectorId}/config/auth</code>:
<ul>
  <li>On a connector with no app, an app is created from them and linked
  (<code>auth.oauthConfigId</code> in the response) once every required OAuth
  field is present; until then the values are kept on the connector.</li>
  <li>On a linked connector, changed values update the app named in
  <code>auth.oauthConfigId</code>, and only the fields that changed. Without that
  field a new app is created and linked; the previously linked app, which other
  connectors may share, is not changed.</li>
  <li>The new app is named <code>auth.oauthInstanceName</code>, or else after the
  connector.</li>
</ul>
A linked connector stores no copy of its app's secret fields: they are not
stored, and not returned, in <code>config.auth</code>. Its client ID and provider
URLs are the app's. Only an admin's credentials create or change an app.


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

    res, err := s.ConnectorConfiguration.UpdateConnectorConfig(ctx, "<id>", components.UpdateConnectorConfigRequest{
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
| `connectorID`                                                                                      | `string`                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `body`                                                                                             | [components.UpdateConnectorConfigRequest](../../models/components/updateconnectorconfigrequest.md) | :heavy_check_mark:                                                                                 | Request payload                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateConnectorConfigResponse](../../models/operations/updateconnectorconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateConnectorAuthConfig

Update only the authentication configuration.<br><br>
<b>Use Case:</b><br>
Use this when you need to update credentials without changing
sync or filter settings. Useful for credential rotation.<br><br>
<b>Prerequisites:</b><br>
Connector must be disabled. This endpoint clears OAuth state,
requiring re-authentication for OAuth connectors.<br><br>
<b>Secrets:</b><br>
A secret field sent as <code>••••••••</code> keeps the stored value; send a
new value to replace it. The response masks stored secrets the same way.


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

    res, err := s.ConnectorConfiguration.UpdateConnectorAuthConfig(ctx, "<id>", components.UpdateConnectorAuthRequest{
        Auth: components.ConnectorAuthConfig{},
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
| `body`                                                                                         | [components.UpdateConnectorAuthRequest](../../models/components/updateconnectorauthrequest.md) | :heavy_check_mark:                                                                             | Request payload                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateConnectorAuthConfigResponse](../../models/operations/updateconnectorauthconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateConnectorFiltersSyncConfig

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

    res, err := s.ConnectorConfiguration.UpdateConnectorFiltersSyncConfig(ctx, "<id>", components.UpdateConnectorFiltersSyncRequest{})
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
| `connectorID`                                                                                                | `string`                                                                                                     | :heavy_check_mark:                                                                                           | Connector instance ID                                                                                        |
| `body`                                                                                                       | [components.UpdateConnectorFiltersSyncRequest](../../models/components/updateconnectorfilterssyncrequest.md) | :heavy_check_mark:                                                                                           | Request payload                                                                                              |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateConnectorFiltersSyncConfigResponse](../../models/operations/updateconnectorfilterssyncconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |