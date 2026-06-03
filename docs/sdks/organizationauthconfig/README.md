# OrganizationAuthConfig

## Overview

Admin configuration of authentication methods including MFA steps and allowed providers

### Available Operations

* [GetAuthMethods](#getauthmethods) - Get organization authentication methods
* [UpdateAuthMethod](#updateauthmethod) - Update organization authentication methods
* [SetUpAuthConfig](#setupauthconfig) - Set up auth configuration

## GetAuthMethods

Retrieve the configured authentication methods for the organization.

**Response Structure:**

Returns an array of authentication steps, each containing:
- `order`: Step number (1-3)
- `allowedMethods`: Array of methods allowed for that step

**Example Response:**

```json
{
  "authMethods": [
    { "order": 1, "allowedMethods": [{ "type": "password" }, { "type": "google" }] },
    { "order": 2, "allowedMethods": [{ "type": "otp" }] }
  ]
}
```

**Admin Access Required:** Only organization admins can view auth configuration.


### Example Usage

<!-- UsageSnippet language="go" operationID="getAuthMethods" method="get" path="/orgAuthConfig/authMethods" -->
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

    res, err := s.OrganizationAuthConfig.GetAuthMethods(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthConfig != nil {
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

**[*operations.GetAuthMethodsResponse](../../models/operations/getauthmethodsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateAuthMethod

Update the authentication methods configuration for an organization.
This allows admins to configure single or multi-factor authentication.

**Validation Rules:**
- Minimum 1 step, maximum 3 steps
- Each step must have a unique order (1, 2, or 3)
- No duplicate methods within the same step
- No method can appear in multiple steps
- Each step must have at least one allowed method

**Available Methods:**
- `password`: Email/password authentication
- `otp`: One-time password via email
- `google`: Google OAuth 2.0
- `microsoft`: Microsoft OAuth 2.0
- `azureAd`: Azure Active Directory
- `samlSso`: SAML 2.0 Single Sign-On
- `oauth`: Generic OAuth 2.0 provider

**Example - Single Factor (Password or Google):**

```json
{
  "authMethod": [
    { "order": 1, "allowedMethods": [{ "type": "password" }, { "type": "google" }] }
  ]
}
```

**Example - Two Factor (Password + OTP):**

```json
{
  "authMethod": [
    { "order": 1, "allowedMethods": [{ "type": "password" }] },
    { "order": 2, "allowedMethods": [{ "type": "otp" }] }
  ]
}
```

**Admin Access Required:** Only organization admins can update auth configuration.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateAuthMethod" method="post" path="/orgAuthConfig/updateAuthMethod" -->
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

    res, err := s.OrganizationAuthConfig.UpdateAuthMethod(ctx, operations.UpdateAuthMethodRequest{
        AuthMethod: []components.AuthStep{
            components.AuthStep{
                Order: 195644,
                AllowedMethods: []components.AuthMethod{
                    components.AuthMethod{
                        Type: components.TypeSamlSso,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAuthMethodResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateAuthMethodRequest](../../models/operations/updateauthmethodrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateAuthMethodResponse](../../models/operations/updateauthmethodresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SetUpAuthConfig

Set up or initialize the organization's authentication configuration.


### Example Usage

<!-- UsageSnippet language="go" operationID="setUpAuthConfig" method="post" path="/orgAuthConfig" -->
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

    res, err := s.OrganizationAuthConfig.SetUpAuthConfig(ctx, components.OrgAuthConfigCreateRequest{
        ContactEmail: "Buster_Waelchi@yahoo.com",
        RegisteredName: "<value>",
        AdminFullName: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgAuthConfigSetupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.OrgAuthConfigCreateRequest](../../models/components/orgauthconfigcreaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.SetUpAuthConfigResponse](../../models/operations/setupauthconfigresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |