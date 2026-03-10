# Users

## Overview

User management operations

### Available Operations

* [GetAll](#getall) - Get all users
* [Create](#create) - Create a new user
* [Get](#get) - Get user by ID
* [Update](#update) - Update user
* [Delete](#delete) - Delete user
* [GetEmail](#getemail) - Get user email by ID
* [UpdateEmail](#updateemail) - Update user email
* [UploadDisplayPicture](#uploaddisplaypicture) - Upload display picture
* [GetDisplayPicture](#getdisplaypicture) - Get display picture
* [RemoveDisplayPicture](#removedisplaypicture) - Remove display picture
* [BulkInvite](#bulkinvite) - Bulk invite users
* [ResendInvite](#resendinvite) - Resend user invite
* [ListGraph](#listgraph) - List users (paginated with graph data)
* [Unblock](#unblock) - Unblock a user in organization
* [GetAllWithGroups](#getallwithgroups) - Get all users with groups
* [GetByIds](#getbyids) - Get users by IDs
* [UpdateFullName](#updatefullname) - Update user full name
* [UpdateFirstName](#updatefirstname) - Update user first name
* [UpdateLastName](#updatelastname) - Update user last name
* [UpdateDesignation](#updatedesignation) - Update user designation
* [CheckAdmin](#checkadmin) - Check if user is admin
* [ListTeams](#listteams) - Get user teams

## GetAll

Retrieve a paginated list of all users in the organization.<br><br>
<b>Overview:</b><br>
This endpoint returns all active users in your organization. It's the primary endpoint for listing and displaying users in admin dashboards, user directories, and selection interfaces.<br><br>
<b>Response Data:</b><br>
<ul>
<li>User profile information (name, email, designation)</li>
<li>Account status (active, pending invitation, disabled)</li>
<li>Login history (hasLoggedIn flag)</li>
<li>Timestamps (createdAt, updatedAt)</li>
</ul>
<b>Privacy Controls:</b><br>
<ul>
<li>Email addresses may be masked based on organization settings</li>
<li>Sensitive fields (password, tokens) are never exposed</li>
<li>Deleted users are excluded from results</li>
</ul>
<b>Performance Notes:</b><br>
<ul>
<li>Results are cached for improved performance</li>
<li>For large organizations, consider using pagination</li>
<li>Use <code>/users/by-ids</code> for fetching specific users</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getAllUsers" method="get" path="/users" -->
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

    res, err := s.Users.GetAll(ctx, pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](50), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Users != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number for pagination (1-based)                     |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | Number of users per page                                 |
| `search`                                                 | **string*                                                | :heavy_minus_sign:                                       | Search users by name or email                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAllUsersResponse](../../models/operations/getallusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new user account in the organization and optionally send an invitation email.<br><br>
<b>Overview:</b><br>
This endpoint creates a new user account. The user will be added to the organization but won't have a password set until they complete the invitation flow or are assigned one by an admin.<br><br>
<b>Invitation Flow:</b><br>
<ol>
<li>Admin creates user with this endpoint</li>
<li>System generates invitation token</li>
<li>User receives invitation email (if sendInvite is true)</li>
<li>User clicks link and sets their password</li>
<li>User can now log in normally</li>
</ol>
<b>Validation Rules:</b><br>
<ul>
<li><code>fullName</code>: Required, 1-100 characters</li>
<li><code>email</code>: Required, valid email format, must be unique in org</li>
<li><code>mobile</code>: Optional, format: +[country][number] (10-15 digits)</li>
<li><code>designation</code>: Optional, job title or role</li>
</ul>
<b>Side Effects:</b><br>
<ul>
<li>User is automatically added to the "everyone" group</li>
<li>Invitation email sent if <code>sendInvite: true</code></li>
<li>User creation event published to event bus</li>
<li>Audit log entry created</li>
</ul>
<b>Authorization:</b><br>
Only organization administrators can create new users.


### Example Usage

<!-- UsageSnippet language="go" operationID="createUser" method="post" path="/users" -->
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

    res, err := s.Users.Create(ctx, operations.CreateUserRequest{
        FullName: "John Smith",
        Email: "john.smith@company.com",
        Mobile: pipeshub.Pointer("+15551234567"),
        Designation: pipeshub.Pointer("Software Engineer"),
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.CreateUserRequest](../../models/operations/createuserrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve detailed information about a specific user by their unique identifier.<br><br>
<b>Overview:</b><br>
This endpoint returns the complete user profile for the specified user ID. Use this to display user details in profiles, settings pages, or when you need full user information.<br><br>
<b>Response Data:</b><br>
<ul>
<li>Basic info: fullName, firstName, lastName, email</li>
<li>Contact: mobile, address</li>
<li>Professional: designation</li>
<li>Status: hasLoggedIn, isDeleted, accountStatus</li>
<li>Metadata: createdAt, updatedAt, createdBy</li>
</ul>
<b>Privacy Notes:</b><br>
<ul>
<li>Email may be masked for non-admin users based on org settings</li>
<li>Password and sensitive tokens are never returned</li>
<li>Display picture URL returned if set</li>
</ul>
<b>Related Endpoints:</b><br>
<ul>
<li><code>GET /users/{id}/email</code> - Get just the email (admin only)</li>
<li><code>GET /users/fetch/with-groups</code> - Get user with group memberships</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getUserById" method="get" path="/users/{id}" -->
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

    res, err := s.Users.Get(ctx, "507f1f77bcf86cd799439011")
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | User ID (24-character MongoDB ObjectId)                  | 507f1f77bcf86cd799439011                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetUserByIDResponse](../../models/operations/getuserbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update user profile information. Users can update their own profile, admins can update any user.<br><br>
<b>Overview:</b><br>
This endpoint allows updating user profile fields. The scope of allowed updates depends on the requester's role and relationship to the user being updated.<br><br>
<b>Authorization Matrix:</b><br>
<ul>
<li><b>Self-update:</b> Users can update their own fullName, mobile, designation, address</li>
<li><b>Admin-update:</b> Admins can update any field for any user</li>
<li><b>Email changes:</b> Require admin privileges and trigger re-verification</li>
</ul>
<b>Updatable Fields:</b><br>
<ul>
<li><code>fullName</code>: Display name (also updates firstName/lastName if parsed)</li>
<li><code>firstName</code>: First name only</li>
<li><code>lastName</code>: Last name only</li>
<li><code>email</code>: Email address (admin only, triggers verification)</li>
<li><code>mobile</code>: Phone number with country code</li>
<li><code>designation</code>: Job title</li>
<li><code>address</code>: Full address object</li>
</ul>
<b>Validation Rules:</b><br>
<ul>
<li>Email must be unique within the organization</li>
<li>Mobile must match pattern: +[country][number]</li>
<li>Name fields: 1-100 characters</li>
</ul>
<b>Side Effects:</b><br>
<ul>
<li>User update event published to event bus</li>
<li>Audit log entry created for admin updates</li>
<li>Email change triggers verification email</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="updateUser" method="put" path="/users/{id}" -->
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

    res, err := s.Users.Update(ctx, "507f1f77bcf86cd799439011", operations.UpdateUserRequestBody{
        FullName: pipeshub.Pointer("John Smith"),
        FirstName: pipeshub.Pointer("John"),
        LastName: pipeshub.Pointer("Smith"),
        Email: pipeshub.Pointer("john.smith@company.com"),
        Mobile: pipeshub.Pointer("+15551234567"),
        Designation: pipeshub.Pointer("Senior Software Engineer"),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | User ID (24-character MongoDB ObjectId)                                              | 507f1f77bcf86cd799439011                                                             |
| `body`                                                                               | [operations.UpdateUserRequestBody](../../models/operations/updateuserrequestbody.md) | :heavy_check_mark:                                                                   | Request payload                                                                      |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.UpdateUserResponse](../../models/operations/updateuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Soft delete a user from the organization. The user account is deactivated but data is retained for audit purposes.<br><br>
<b>Overview:</b><br>
This endpoint performs a soft delete on a user account. The user is marked as deleted and can no longer access the system, but their data is retained for compliance and audit purposes.<br><br>
<b>What Happens on Delete:</b><br>
<ol>
<li>User's <code>isDeleted</code> flag is set to true</li>
<li>User's password is cleared</li>
<li>User is removed from all user groups</li>
<li>User's active sessions are invalidated</li>
<li>User deletion event is published</li>
</ol>
<b>Restrictions:</b><br>
<ul>
<li>Cannot delete organization admins (demote first)</li>
<li>Cannot delete the organization owner</li>
<li>Cannot delete yourself through this endpoint</li>
<li>Cannot delete already-deleted users</li>
</ul>
<b>Data Retention:</b><br>
<ul>
<li>User profile data is retained (soft delete)</li>
<li>User's documents and content remain with updated ownership</li>
<li>Audit logs are preserved</li>
<li>Data can be permanently purged by super admin if required</li>
</ul>
<b>Recovery:</b><br>
Deleted users can be restored by organization admins within a configurable retention period.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteUser" method="delete" path="/users/{id}" -->
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

    res, err := s.Users.Delete(ctx, "507f1f77bcf86cd799439011")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | User ID (24-character MongoDB ObjectId)                  | 507f1f77bcf86cd799439011                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetEmail

Retrieve the email address for a specific user. This is a dedicated endpoint for email lookup with proper access controls.<br><br>
<b>Overview:</b><br>
This endpoint provides direct access to a user's email address. It exists separately from the main user endpoint to allow granular permission control over email visibility.<br><br>
<b>Use Cases:</b><br>
<ul>
<li>Admin communication workflows</li>
<li>Invitation and notification systems</li>
<li>Email-based user lookup</li>
<li>Contact information export</li>
</ul>
<b>Privacy Considerations:</b><br>
<ul>
<li>Only organization admins can access this endpoint</li>
<li>Access is logged for audit purposes</li>
<li>Consider GDPR/privacy regulations when exposing emails</li>
</ul>
<b>Authorization:</b><br>
Requires admin privileges. Regular users should use the main user endpoint which may mask emails based on organization settings.


### Example Usage

<!-- UsageSnippet language="go" operationID="getUserEmailById" method="get" path="/users/{id}/email" -->
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

    res, err := s.Users.GetEmail(ctx, "507f1f77bcf86cd799439011")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | User ID (24-character MongoDB ObjectId)                  | 507f1f77bcf86cd799439011                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetUserEmailByIDResponse](../../models/operations/getuseremailbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateEmail

Update the email address of a user.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateEmail" method="patch" path="/users/{id}/email" -->
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

    res, err := s.Users.UpdateEmail(ctx, "<id>", operations.UpdateEmailRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `body`                                                                                 | [operations.UpdateEmailRequestBody](../../models/operations/updateemailrequestbody.md) | :heavy_check_mark:                                                                     | Request payload                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateEmailResponse](../../models/operations/updateemailresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UploadDisplayPicture

Upload or update the display picture (avatar) for the authenticated user.<br><br>
<b>Overview:</b><br>
This endpoint allows users to upload their profile picture. The image is processed, resized, and stored for use throughout the application.<br><br>
<b>File Requirements:</b><br>
<ul>
<li><b>Allowed types:</b> PNG, JPEG, JPG, WebP, GIF</li>
<li><b>Maximum size:</b> 1MB (1,048,576 bytes)</li>
<li><b>Recommended dimensions:</b> 256x256 pixels or larger</li>
<li><b>Aspect ratio:</b> Square recommended (will be cropped to square)</li>
</ul>
<b>Image Processing:</b><br>
<ul>
<li>Images are automatically resized to standard dimensions</li>
<li>Converted to JPEG for consistency and smaller file size</li>
<li>Multiple sizes may be generated (thumbnail, standard, large)</li>
<li>Original is not preserved</li>
</ul>
<b>Side Effects:</b><br>
<ul>
<li>Previous display picture is replaced</li>
<li>Cached images are invalidated</li>
<li>CDN cache may take time to update</li>
</ul>
<b>Authorization:</b><br>
Users can only upload their own display picture. Admins cannot upload on behalf of other users.


### Example Usage

<!-- UsageSnippet language="go" operationID="uploadUserDisplayPicture" method="put" path="/users/dp" -->
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

    res, err := s.Users.UploadDisplayPicture(ctx, operations.UploadUserDisplayPictureRequest{
        File: operations.UploadUserDisplayPictureFile{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UploadUserDisplayPictureRequest](../../models/operations/uploaduserdisplaypicturerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UploadUserDisplayPictureResponse](../../models/operations/uploaduserdisplaypictureresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetDisplayPicture

Retrieve the current user's display picture image.<br><br>
<b>Overview:</b><br>
This endpoint returns the user's profile picture as binary image data. Use this for displaying the user's avatar in the application.<br><br>
<b>Response Format:</b><br>
<ul>
<li>Returns raw image data (not JSON)</li>
<li>Content-Type header indicates image format (typically image/jpeg)</li>
<li>Suitable for use directly in &lt;img&gt; src or CSS background</li>
</ul>
<b>Caching:</b><br>
<ul>
<li>Response includes cache headers for browser caching</li>
<li>Use ETag for conditional requests</li>
<li>Cache invalidated when picture is updated</li>
</ul>
<b>Alternative:</b><br>
For signed URL access, use the user profile endpoint which returns a <code>displayPictureUrl</code> field.


### Example Usage

<!-- UsageSnippet language="go" operationID="getUserDisplayPicture" method="get" path="/users/dp" -->
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

    res, err := s.Users.GetDisplayPicture(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredImageJpegResponseStream != nil {
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

**[*operations.GetUserDisplayPictureResponse](../../models/operations/getuserdisplaypictureresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RemoveDisplayPicture

Remove the current user's display picture and revert to default avatar.<br><br>
<b>Overview:</b><br>
This endpoint permanently removes the user's uploaded profile picture. After removal, the user will display a default avatar (typically initials or generic icon).<br><br>
<b>What Happens:</b><br>
<ul>
<li>Profile picture file is deleted from storage</li>
<li>User profile updated to remove picture reference</li>
<li>Cached images invalidated</li>
<li>Default avatar will be shown in UI</li>
</ul>
<b>Note:</b><br>
This action is immediate and irreversible. To restore a picture, user must upload a new one.


### Example Usage

<!-- UsageSnippet language="go" operationID="removeUserDisplayPicture" method="delete" path="/users/dp" -->
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

    res, err := s.Users.RemoveDisplayPicture(ctx)
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

**[*operations.RemoveUserDisplayPictureResponse](../../models/operations/removeuserdisplaypictureresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## BulkInvite

Invite multiple users to the organization in a single operation. Ideal for onboarding entire teams at once.<br><br>
<b>Overview:</b><br>
This endpoint creates user accounts for multiple email addresses and sends invitation emails to all of them. It's the most efficient way to add multiple users to your organization.<br><br>
<b>Invitation Flow:</b><br>
<ol>
<li>Validate all email addresses</li>
<li>Check for existing accounts (skip duplicates)</li>
<li>Create user accounts for new emails</li>
<li>Restore any previously deleted accounts</li>
<li>Add users to specified groups (optional)</li>
<li>Send invitation emails to all new users</li>
</ol>
<b>Requirements:</b><br>
<ul>
<li><b>Account Type:</b> Business accounts only (not individual)</li>
<li><b>SMTP:</b> Email configuration must be set up</li>
<li><b>Authorization:</b> Admin privileges required</li>
</ul>
<b>Email Processing:</b><br>
<ul>
<li>Duplicate emails are automatically skipped</li>
<li>Invalid email formats are rejected</li>
<li>Existing users are not re-invited (use resend-invite)</li>
<li>Previously deleted users are restored</li>
</ul>
<b>Response Details:</b><br>
Response includes count of successful invites and any failures with reasons.


### Example Usage

<!-- UsageSnippet language="go" operationID="bulkInviteUsers" method="post" path="/users/bulk/invite" -->
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

    res, err := s.Users.BulkInvite(ctx, operations.BulkInviteUsersRequest{
        Emails: []string{
            "user1@company.com",
            "user2@company.com",
        },
        GroupIds: []string{
            "507f1f77bcf86cd799439011",
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
| `request`                                                                              | [operations.BulkInviteUsersRequest](../../models/operations/bulkinviteusersrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.BulkInviteUsersResponse](../../models/operations/bulkinviteusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ResendInvite

Resend the invitation email to a user who hasn't completed their account setup.<br><br>
<b>Overview:</b><br>
This endpoint resends the invitation email to a user who was previously invited but hasn't logged in yet. Useful when the original invitation email was lost, expired, or ended up in spam.<br><br>
<b>When to Use:</b><br>
<ul>
<li>User didn't receive original invitation</li>
<li>Invitation link expired</li>
<li>User forgot to complete setup</li>
<li>Email went to spam folder</li>
</ul>
<b>Requirements:</b><br>
<ul>
<li>User must exist in the system</li>
<li>User must NOT have logged in yet (hasLoggedIn: false)</li>
<li>SMTP configuration must be active</li>
<li>Admin privileges required</li>
</ul>
<b>What Happens:</b><br>
<ul>
<li>Generates a new invitation token</li>
<li>Invalidates any previous invitation links</li>
<li>Sends new invitation email</li>
<li>Resets invitation expiry timer</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="resendUserInvite" method="post" path="/users/{id}/resend-invite" -->
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

    res, err := s.Users.ResendInvite(ctx, "507f1f77bcf86cd799439011")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | User ID of the user to resend invitation to              | 507f1f77bcf86cd799439011                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ResendUserInviteResponse](../../models/operations/resenduserinviteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListGraph

Retrieve a paginated list of users with enhanced search capabilities using the graph service.<br><br>
<b>Overview:</b><br>
This endpoint provides advanced user listing with full-text search, pagination, and optional relationship data from the knowledge graph. It's optimized for large organizations with thousands of users.<br><br>
<b>Search Capabilities:</b><br>
<ul>
<li>Full-text search across name and email</li>
<li>Fuzzy matching for typo tolerance</li>
<li>Results ranked by relevance</li>
</ul>
<b>Use Cases:</b><br>
<ul>
<li>User directory with search</li>
<li>Autocomplete user selection</li>
<li>Admin user management lists</li>
<li>User analytics dashboards</li>
</ul>
<b>Performance:</b><br>
<ul>
<li>Powered by graph database for fast queries</li>
<li>Supports pagination for large datasets</li>
<li>Results cached for repeated queries</li>
</ul>
<b>vs /users endpoint:</b><br>
Use this endpoint when you need advanced search or are dealing with large user bases. Use <code>/users</code> for simple full-list retrieval.


### Example Usage

<!-- UsageSnippet language="go" operationID="listUsersGraph" method="get" path="/users/graph/list" -->
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

    res, err := s.Users.ListGraph(ctx, operations.ListUsersGraphRequest{
        Search: pipeshub.Pointer("john"),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListUsersGraphRequest](../../models/operations/listusersgraphrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListUsersGraphResponse](../../models/operations/listusersgraphresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Unblock

Unblock a previously blocked user within the authenticated administrator's organization.<br><br>

<b>Overview:</b><br>
This endpoint updates the user's credential record by setting <code>isBlocked</code> to <code>false</code> 
and resetting <code>wrongCredentialCount</code> to <code>0</code>.<br><br>

<b>Authorization:</b><br>
<ul>
<li><b>Admin only:</b> Only organization administrators can unblock users</li>
<li>Requires a valid Bearer token</li>
</ul>

<b>Validation & Conditions:</b><br>
<ul>
<li>User must belong to the same <code>orgId</code> as the authenticated admin</li>
<li>User must currently be blocked (<code>isBlocked: true</code>)</li>
<li>User must not be deleted (<code>isDeleted: false</code>)</li>
</ul>

<b>Note:</b> If the user is not blocked or does not exist in the organization, a 404 response is returned.


### Example Usage

<!-- UsageSnippet language="go" operationID="unblockUser" method="put" path="/users/{id}/unblock" -->
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

    res, err := s.Users.Unblock(ctx, "507f1f77bcf86cd799439011")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | User ID to unblock                                       | 507f1f77bcf86cd799439011                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.UnblockUserResponse](../../models/operations/unblockuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAllWithGroups

Retrieve all users in the organization along with their group memberships.


### Example Usage

<!-- UsageSnippet language="go" operationID="getAllUsersWithGroups" method="get" path="/users/fetch/with-groups" -->
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

    res, err := s.Users.GetAllWithGroups(ctx)
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

**[*operations.GetAllUsersWithGroupsResponse](../../models/operations/getalluserswithgroupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetByIds

Retrieve multiple users by their IDs in a single request.


### Example Usage

<!-- UsageSnippet language="go" operationID="getUsersByIds" method="post" path="/users/by-ids" -->
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

    res, err := s.Users.GetByIds(ctx, operations.GetUsersByIdsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetUsersByIdsRequest](../../models/operations/getusersbyidsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetUsersByIdsResponse](../../models/operations/getusersbyidsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateFullName

Update the full name of a user.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateFullName" method="patch" path="/users/{id}/fullname" -->
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

    res, err := s.Users.UpdateFullName(ctx, "<id>", operations.UpdateFullNameRequestBody{})
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
| `id`                                                                                         | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `body`                                                                                       | [operations.UpdateFullNameRequestBody](../../models/operations/updatefullnamerequestbody.md) | :heavy_check_mark:                                                                           | Request payload                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateFullNameResponse](../../models/operations/updatefullnameresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateFirstName

Update the first name of a user.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateFirstName" method="patch" path="/users/{id}/firstName" -->
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

    res, err := s.Users.UpdateFirstName(ctx, "<id>", operations.UpdateFirstNameRequestBody{})
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
| `id`                                                                                           | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `body`                                                                                         | [operations.UpdateFirstNameRequestBody](../../models/operations/updatefirstnamerequestbody.md) | :heavy_check_mark:                                                                             | Request payload                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateFirstNameResponse](../../models/operations/updatefirstnameresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateLastName

Update the last name of a user.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateLastName" method="patch" path="/users/{id}/lastName" -->
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

    res, err := s.Users.UpdateLastName(ctx, "<id>", operations.UpdateLastNameRequestBody{})
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
| `id`                                                                                         | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `body`                                                                                       | [operations.UpdateLastNameRequestBody](../../models/operations/updatelastnamerequestbody.md) | :heavy_check_mark:                                                                           | Request payload                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateLastNameResponse](../../models/operations/updatelastnameresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateDesignation

Update the designation/title of a user.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateDesignation" method="patch" path="/users/{id}/designation" -->
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

    res, err := s.Users.UpdateDesignation(ctx, "<id>", operations.UpdateDesignationRequestBody{})
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
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `body`                                                                                             | [operations.UpdateDesignationRequestBody](../../models/operations/updatedesignationrequestbody.md) | :heavy_check_mark:                                                                                 | Request payload                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateDesignationResponse](../../models/operations/updatedesignationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CheckAdmin

Check whether the specified user has admin privileges. Returns 200 OK if the user is an admin.


### Example Usage

<!-- UsageSnippet language="go" operationID="adminCheck" method="get" path="/users/{id}/adminCheck" -->
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

    res, err := s.Users.CheckAdmin(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AdminCheckResponse](../../models/operations/admincheckresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListTeams

Retrieve teams associated with the authenticated user.


### Example Usage

<!-- UsageSnippet language="go" operationID="getUserTeamsViaUsers" method="get" path="/users/teams/list" -->
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

    res, err := s.Users.ListTeams(ctx)
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

**[*operations.GetUserTeamsViaUsersResponse](../../models/operations/getuserteamsviausersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |