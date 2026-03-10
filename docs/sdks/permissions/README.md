# Permissions

## Overview

Permission management for knowledge bases

### Available Operations

* [GrantKBAccess](#grantkbaccess) - Grant permissions
* [List](#list) - List permissions
* [Update](#update) - Update permissions
* [DeleteFromKB](#deletefromkb) - Remove permissions

## GrantKBAccess

Grant access permissions to users or teams for a knowledge base.<br><br>
<b>Required Permission:</b> OWNER or ORGANIZER<br><br>
<b>Permission Roles (highest to lowest):</b><br>
<ol>
<li><b>OWNER:</b> Full control, can delete KB, manage all permissions</li>
<li><b>ORGANIZER:</b> Can manage permissions (except OWNER), edit KB settings</li>
<li><b>FILEORGANIZER:</b> Can create/delete folders, organize content</li>
<li><b>WRITER:</b> Can upload, edit, delete records</li>
<li><b>COMMENTER:</b> Can add comments (if supported)</li>
<li><b>READER:</b> View-only access</li>
</ol>
<b>Grant to Multiple:</b><br>
Provide arrays of userIds and/or teamIds to grant the same role to multiple entities.


### Example Usage

<!-- UsageSnippet language="go" operationID="createKBPermission" method="post" path="/knowledgeBase/{kbId}/permissions" -->
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

    res, err := s.Permissions.GrantKBAccess(ctx, "<id>", operations.CreateKBPermissionRequestBody{
        Role: operations.CreateKBPermissionRoleOwner,
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
| `kbID`                                                                                               | *string*                                                                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `body`                                                                                               | [operations.CreateKBPermissionRequestBody](../../models/operations/createkbpermissionrequestbody.md) | :heavy_check_mark:                                                                                   | Request payload                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateKBPermissionResponse](../../models/operations/createkbpermissionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Retrieve all permissions granted on a knowledge base.<br><br>
<b>Required Permission:</b> ORGANIZER or higher to see all permissions, others see only their own.


### Example Usage

<!-- UsageSnippet language="go" operationID="listKBPermissions" method="get" path="/knowledgeBase/{kbId}/permissions" -->
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

    res, err := s.Permissions.List(ctx, "<id>")
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
| `kbID`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListKBPermissionsResponse](../../models/operations/listkbpermissionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update permission roles for users or teams.<br><br>
<b>Required Permission:</b> OWNER or ORGANIZER


### Example Usage

<!-- UsageSnippet language="go" operationID="updateKBPermissions" method="put" path="/knowledgeBase/{kbId}/permissions" -->
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

    res, err := s.Permissions.Update(ctx, "<id>", operations.UpdateKBPermissionsRequestBody{
        Role: operations.UpdateKBPermissionsRoleWriter,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `kbID`                                                                                                 | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `body`                                                                                                 | [operations.UpdateKBPermissionsRequestBody](../../models/operations/updatekbpermissionsrequestbody.md) | :heavy_check_mark:                                                                                     | Request body for Update permissions                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateKBPermissionsResponse](../../models/operations/updatekbpermissionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteFromKB

Remove access permissions from users or teams.<br><br>
<b>Required Permission:</b> OWNER or ORGANIZER<br><br>
<b>Note:</b> Cannot remove the last OWNER from a KB.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteKBPermissions" method="delete" path="/knowledgeBase/{kbId}/permissions" -->
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

    res, err := s.Permissions.DeleteFromKB(ctx, "<id>", operations.DeleteKBPermissionsRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `kbID`                                                                                                 | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `body`                                                                                                 | [operations.DeleteKBPermissionsRequestBody](../../models/operations/deletekbpermissionsrequestbody.md) | :heavy_check_mark:                                                                                     | Request body for Remove permissions                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.DeleteKBPermissionsResponse](../../models/operations/deletekbpermissionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |