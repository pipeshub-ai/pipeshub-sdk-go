# Organizations

## Overview

Organization management operations

### Available Operations

* [CheckExists](#checkexists) - Check if organization exists
* [Create](#create) - Create organization
* [GetCurrent](#getcurrent) - Get current organization
* [Update](#update) - Update organization
* [Delete](#delete) - Delete organization
* [UploadLogo](#uploadlogo) - Upload organization logo
* [GetLogo](#getlogo) - Get organization logo
* [DeleteLogo](#deletelogo) - Delete organization logo
* [GetOnboardingStatus](#getonboardingstatus) - Get onboarding status
* [UpdateOnboardingStatus](#updateonboardingstatus) - Update onboarding status

## CheckExists

Check if any organization has been created in the system. This is typically the first API call made during initial setup.<br><br>
<b>Overview:</b><br>
This public endpoint determines whether the system has been initialized with an organization. Used by the frontend to decide whether to show the setup wizard or the login screen.<br><br>
<b>Use Cases:</b><br>
<ul>
<li>First-time setup detection</li>
<li>Onboarding flow decisions</li>
<li>System initialization checks</li>
</ul>
<b>Response:</b><br>
<ul>
<li><code>exists: true</code> - Organization exists, show login</li>
<li><code>exists: false</code> - No organization, show setup wizard</li>
</ul>
<b>Note:</b> This endpoint requires no authentication and is publicly accessible.


### Example Usage

<!-- UsageSnippet language="go" operationID="checkOrgExists" method="get" path="/org/exists" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.Organizations.CheckExists(ctx)
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

**[*operations.CheckOrgExistsResponse](../../models/operations/checkorgexistsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new organization and its first admin user. This is the initial setup endpoint for new PipesHub installations.<br><br>
<b>Overview:</b><br>
This endpoint performs the complete initial setup of a PipesHub instance, including creating the organization entity and its first administrator account. Should only be called once during initial setup.<br><br>
<b>Setup Flow:</b><br>
<ol>
<li>Frontend calls <code>/org/exists</code> to check if setup is needed</li>
<li>If no organization exists, show setup wizard</li>
<li>Collect organization and admin details</li>
<li>Call this endpoint to create organization</li>
<li>User is automatically logged in as admin</li>
</ol>
<b>What Gets Created:</b><br>
<ul>
<li>Organization entity with provided details</li>
<li>Admin user account with provided credentials</li>
<li>Default user groups (admin, everyone)</li>
<li>Default system settings</li>
<li>Initial authentication configuration</li>
</ul>
<b>Account Types:</b><br>
<ul>
<li><code>individual</code>: Single-user account, limited team features</li>
<li><code>business</code>: Multi-user organization with full features</li>
</ul>
<b>Security:</b><br>
<ul>
<li>This endpoint only works if no organization exists</li>
<li>Password must meet security requirements</li>
<li>Email verification may be required based on config</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="createOrganization" method="post" path="/org" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.Organizations.Create(ctx, operations.CreateOrganizationRequest{
        AccountType: operations.AccountTypeBusiness,
        ShortName: pipeshub.Pointer("Acme"),
        RegisteredName: pipeshub.Pointer("Acme Corporation Inc."),
        ContactEmail: "admin@acme.com",
        AdminFullName: "John Smith",
        Password: "SecurePassword123!",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateOrganizationRequest](../../models/operations/createorganizationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateOrganizationResponse](../../models/operations/createorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCurrent

Retrieve details about the authenticated user's organization.<br><br>
<b>Overview:</b><br>
This endpoint returns complete information about the current user's organization, including profile data, settings, and configuration. Use this for organization profile pages and settings.<br><br>
<b>Response Includes:</b><br>
<ul>
<li>Organization profile (name, email, address)</li>
<li>Account type and billing status</li>
<li>Feature flags and limits</li>
<li>Branding settings (logo, colors)</li>
<li>Creation and modification timestamps</li>
</ul>
<b>Use Cases:</b><br>
<ul>
<li>Organization profile pages</li>
<li>Settings and configuration screens</li>
<li>Billing and subscription displays</li>
<li>White-label branding retrieval</li>
</ul>
<b>Note:</b><br>
All authenticated users can access this endpoint to view their organization's details.


### Example Usage

<!-- UsageSnippet language="go" operationID="getCurrentOrganization" method="get" path="/org" -->
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

    res, err := s.Organizations.GetCurrent(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Organization != nil {
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

**[*operations.GetCurrentOrganizationResponse](../../models/operations/getcurrentorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update organization profile and settings information.<br><br>
<b>Overview:</b><br>
This endpoint allows administrators to update the organization's profile information, contact details, and address. Used in the organization settings section of the admin panel.<br><br>
<b>Updatable Fields:</b><br>
<ul>
<li><code>registeredName</code>: Official registered/legal name of the organization</li>
<li><code>shortName</code>: Short display name used in UI</li>
<li><code>phoneNumber</code>: Primary contact phone number</li>
<li><code>permanentAddress</code>: Full address object with street, city, state, country, postal code</li>
</ul>
<b>Restrictions:</b><br>
<ul>
<li>Only organization admins can perform updates</li>
<li>Contact email cannot be changed through this endpoint</li>
<li>Account type cannot be changed after creation</li>
</ul>
<b>Side Effects:</b><br>
<ul>
<li>Organization update event is published</li>
<li>Cached organization data is invalidated</li>
<li>Changes are reflected immediately across all services</li>
</ul>
<b>Validation:</b><br>
<ul>
<li>Phone number must be valid international format</li>
<li>Address fields have maximum length constraints</li>
<li>Names cannot be empty if provided</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="updateOrganization" method="put" path="/org" -->
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

    res, err := s.Organizations.Update(ctx, operations.UpdateOrganizationRequest{
        RegisteredName: pipeshub.Pointer("Acme Corporation Inc."),
        ShortName: pipeshub.Pointer("Acme Corp"),
        PhoneNumber: pipeshub.Pointer("+15551234567"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateOrganizationRequest](../../models/operations/updateorganizationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateOrganizationResponse](../../models/operations/updateorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Permanently delete an organization and all associated data.<br><br>
<b>WARNING:</b> This action is <b>irreversible</b>.<br><br>
<b>Data Deleted:</b><br>
<ul>
<li>All user accounts in the organization</li>
<li>All teams and user groups</li>
<li>All documents and storage data</li>
<li>All configuration and settings</li>
</ul>
<b>Requirements:</b><br>
<ul>
<li>Must be the organization owner (super admin)</li>
<li>Must provide confirmation parameter</li>
<li>All active subscriptions must be cancelled first</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteOrganization" method="delete" path="/org" -->
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

    res, err := s.Organizations.Delete(ctx, operations.ConfirmDelete)
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
| `confirm`                                                | [operations.Confirm](../../models/operations/confirm.md) | :heavy_check_mark:                                       | Must be "DELETE" to confirm deletion                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationResponse](../../models/operations/deleteorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UploadLogo

Upload or update the organization's logo image.<br><br>
<b>Supported Formats:</b><br>
<ul>
<li>PNG (recommended for transparency)</li>
<li>JPG/JPEG</li>
<li>SVG</li>
<li>WebP</li>
</ul>
<b>Requirements:</b><br>
<ul>
<li>Maximum file size: 5MB</li>
<li>Recommended dimensions: 256x256 pixels or higher</li>
<li>Square aspect ratio recommended</li>
</ul>
<b>Side Effects:</b><br>
<ul>
<li>Previous logo is replaced</li>
<li>Multiple sizes may be generated for different use cases</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="uploadOrganizationLogo" method="put" path="/org/logo" -->
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

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Organizations.UploadLogo(ctx, operations.UploadOrganizationLogoRequest{
        Logo: operations.Logo{
            FileName: "example.file",
            Content: example,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UploadOrganizationLogoRequest](../../models/operations/uploadorganizationlogorequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UploadOrganizationLogoResponse](../../models/operations/uploadorganizationlogoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetLogo

Retrieve the organization's logo image or URL.<br><br>
<b>Response Formats:</b><br>
<ul>
<li>Returns a signed URL to access the logo</li>
<li>URL is valid for a limited time (typically 1 hour)</li>
</ul>
<b>Use Cases:</b><br>
<ul>
<li>Displaying logo in navigation/header</li>
<li>Email templates</li>
<li>White-label branding</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getOrganizationLogo" method="get" path="/org/logo" -->
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

    res, err := s.Organizations.GetLogo(ctx)
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

**[*operations.GetOrganizationLogoResponse](../../models/operations/getorganizationlogoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteLogo

Remove the organization's custom logo.<br><br>
<b>Behavior:</b><br>
<ul>
<li>Logo file is permanently deleted from storage</li>
<li>Organization reverts to default/placeholder logo</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteOrganizationLogo" method="delete" path="/org/logo" -->
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

    res, err := s.Organizations.DeleteLogo(ctx)
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

**[*operations.DeleteOrganizationLogoResponse](../../models/operations/deleteorganizationlogoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetOnboardingStatus

Retrieve the organization's onboarding progress and status.<br><br>
<b>Response Details:</b><br>
<ul>
<li>Current onboarding step</li>
<li>Completion status of each step</li>
<li>Overall completion percentage</li>
</ul>
<b>Onboarding Steps:</b><br>
<ul>
<li>Organization profile setup</li>
<li>Admin account configuration</li>
<li>Invite team members</li>
<li>Connect integrations</li>
<li>Configure settings</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getOnboardingStatus" method="get" path="/org/onboarding-status" -->
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

    res, err := s.Organizations.GetOnboardingStatus(ctx)
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

**[*operations.GetOnboardingStatusResponse](../../models/operations/getonboardingstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateOnboardingStatus

Update the organization's onboarding progress.<br><br>
<b>Use Cases:</b><br>
<ul>
<li>Mark a step as completed</li>
<li>Skip optional steps</li>
<li>Complete entire onboarding</li>
</ul>
<b>Behavior:</b><br>
<ul>
<li>Steps must be completed in order (unless skippable)</li>
<li>Completing all required steps marks onboarding as complete</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="updateOnboardingStatus" method="put" path="/org/onboarding-status" -->
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

    res, err := s.Organizations.UpdateOnboardingStatus(ctx, operations.UpdateOnboardingStatusRequest{
        StepID: "invite_team",
        Action: operations.ActionComplete,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateOnboardingStatusRequest](../../models/operations/updateonboardingstatusrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateOnboardingStatusResponse](../../models/operations/updateonboardingstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |