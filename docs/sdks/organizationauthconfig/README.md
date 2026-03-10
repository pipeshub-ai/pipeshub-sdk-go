# OrganizationAuthConfig

## Overview

### Available Operations

* [GetAuthMethods](#getauthmethods) - Get organization authentication methods
* [UpdateAuthMethod](#updateauthmethod) - Update organization authentication methods
* [SetUp](#setup) - Set up auth configuration

## GetAuthMethods

Retrieve the configured authentication methods for the organization.
<br><br>
<b>Response Structure:</b><br>
Returns an array of authentication steps, each containing:<br>
- <code>order</code>: Step number (1-3)<br>
- <code>allowedMethods</code>: Array of methods allowed for that step
<br><br>
<b>Example Response:</b><br>
<pre>
{
  "authMethods": [
    { "order": 1, "allowedMethods": [{ "type": "password" }, { "type": "google" }] },
    { "order": 2, "allowedMethods": [{ "type": "otp" }] }
  ]
}
</pre>
<br>
<b>Admin Access Required:</b> Only organization admins can view auth configuration.


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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateAuthMethod

Update the authentication methods configuration for an organization.
This allows admins to configure single or multi-factor authentication.
<br><br>
<b>Validation Rules:</b><br>
- Minimum 1 step, maximum 3 steps<br>
- Each step must have a unique order (1, 2, or 3)<br>
- No duplicate methods within the same step<br>
- No method can appear in multiple steps<br>
- Each step must have at least one allowed method
<br><br>
<b>Available Methods:</b><br>
- <code>password</code>: Email/password authentication<br>
- <code>otp</code>: One-time password via email<br>
- <code>google</code>: Google OAuth 2.0<br>
- <code>microsoft</code>: Microsoft OAuth 2.0<br>
- <code>azureAd</code>: Azure Active Directory<br>
- <code>samlSso</code>: SAML 2.0 Single Sign-On<br>
- <code>oauth</code>: Generic OAuth 2.0 provider
<br><br>
<b>Example - Single Factor (Password or Google):</b><br>
<pre>
{
  "authMethods": [
    { "order": 1, "allowedMethods": [{ "type": "password" }, { "type": "google" }] }
  ]
}
</pre>
<br>
<b>Example - Two Factor (Password + OTP):</b><br>
<pre>
{
  "authMethods": [
    { "order": 1, "allowedMethods": [{ "type": "password" }] },
    { "order": 2, "allowedMethods": [{ "type": "otp" }] }
  ]
}
</pre>
<br>
<b>Admin Access Required:</b> Only organization admins can update auth configuration.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateAuthMethod" method="post" path="/orgAuthConfig/updateAuthMethod" -->
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

    res, err := s.OrganizationAuthConfig.UpdateAuthMethod(ctx, components.AuthConfig{
        AuthMethods: []components.AuthStep{
            components.AuthStep{
                Order: 195644,
                AllowedMethods: []components.AuthMethod{
                    components.AuthMethod{
                        Type: components.AuthMethodTypeSamlSso,
                    },
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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.AuthConfig](../../models/components/authconfig.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.UpdateAuthMethodResponse](../../models/operations/updateauthmethodresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.AuthError | 400                 | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |

## SetUp

Set up or initialize the organization's authentication configuration.


### Example Usage

<!-- UsageSnippet language="go" operationID="setUpAuthConfig" method="post" path="/orgAuthConfig/" -->
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

    res, err := s.OrganizationAuthConfig.SetUp(ctx, operations.SetUpAuthConfigRequest{})
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
| `request`                                                                              | [operations.SetUpAuthConfigRequest](../../models/operations/setupauthconfigrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.SetUpAuthConfigResponse](../../models/operations/setupauthconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |