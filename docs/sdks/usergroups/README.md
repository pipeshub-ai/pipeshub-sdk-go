# UserGroups

## Overview

### Available Operations

* [Create](#create) - Create user group
* [GetAll](#getall) - Get all user groups
* [GetByID](#getbyid) - Get user group by ID
* [Update](#update) - Update user group
* [Delete](#delete) - Delete user group
* [AddUsers](#addusers) - Add users to group
* [RemoveUsers](#removeusers) - Remove users from group
* [GetUserGroups](#getusergroups) - Get groups for a user
* [GetInGroup](#getingroup) - Get users in group
* [GetStatistics](#getstatistics) - Get group statistics

## Create

Create a new user group within the organization.<br><br>
<b>Group Types:</b><br>
<ul>
<li><code>admin</code> - Administrative group with elevated privileges</li>
<li><code>standard</code> - Regular user group</li>
<li><code>everyone</code> - Automatically includes all organization users</li>
<li><code>custom</code> - Custom group with manual membership management</li>
</ul>
<b>Validation Rules:</b><br>
<ul>
<li>Group name must be unique within the organization</li>
<li>Group name is required and cannot be empty</li>
<li>Type must be one of the allowed values</li>
</ul>
<b>Side Effects:</b><br>
<ul>
<li>Creates a unique slug from the group name</li>
<li>Sets createdBy to the authenticated user</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="createUserGroup" method="post" path="/userGroups" -->
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

    res, err := s.UserGroups.Create(ctx, operations.CreateUserGroupRequest{
        Name: "Engineering Team",
        Type: operations.TypeStandard,
        Description: pipeshub.Pointer("All engineering department members"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateUserGroupRequest](../../models/operations/createusergrouprequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateUserGroupResponse](../../models/operations/createusergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAll

Retrieve all user groups in the organization.<br><br>
<b>Response Details:</b><br>
<ul>
<li>Returns all groups including admin, standard, everyone, and custom types</li>
<li>Groups are returned with their member counts</li>
<li>Soft-deleted groups are excluded by default</li>
</ul>
<b>Use Cases:</b><br>
<ul>
<li>Populating group selection dropdowns</li>
<li>Managing group memberships</li>
<li>Access control configuration</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getAllUserGroups" method="get" path="/userGroups" -->
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

    res, err := s.UserGroups.GetAll(ctx)
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

**[*operations.GetAllUserGroupsResponse](../../models/operations/getallusergroupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetByID

Retrieve detailed information about a specific user group.<br><br>
<b>Response Includes:</b><br>
<ul>
<li>Group metadata (name, type, description)</li>
<li>Member count</li>
<li>Creation and modification timestamps</li>
<li>Creator information</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getUserGroupById" method="get" path="/userGroups/{groupId}" -->
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

    res, err := s.UserGroups.GetByID(ctx, "507f1f77bcf86cd799439011")
    if err != nil {
        log.Fatal(err)
    }
    if res.UserGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `groupID`                                                | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the user group                      | 507f1f77bcf86cd799439011                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetUserGroupByIDResponse](../../models/operations/getusergroupbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update an existing user group's information.<br><br>
<b>Updatable Fields:</b><br>
<ul>
<li><code>name</code> - Display name (must remain unique)</li>
<li><code>description</code> - Group description</li>
</ul>
<b>Note:</b> Group type cannot be changed after creation.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateUserGroup" method="put" path="/userGroups/{groupId}" -->
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

    res, err := s.UserGroups.Update(ctx, "507f1f77bcf86cd799439011", operations.UpdateUserGroupRequestBody{
        Name: pipeshub.Pointer("Engineering Team - Updated"),
        Description: pipeshub.Pointer("All engineering and DevOps members"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `groupID`                                                                                      | *string*                                                                                       | :heavy_check_mark:                                                                             | Unique identifier of the user group to update                                                  | 507f1f77bcf86cd799439011                                                                       |
| `body`                                                                                         | [operations.UpdateUserGroupRequestBody](../../models/operations/updateusergrouprequestbody.md) | :heavy_check_mark:                                                                             | Request payload                                                                                |                                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.UpdateUserGroupResponse](../../models/operations/updateusergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Soft delete a user group.<br><br>
<b>Behavior:</b><br>
<ul>
<li>Group is marked as deleted (isDeleted: true)</li>
<li>Group members are not affected</li>
<li>Group can be restored by admin if needed</li>
</ul>
<b>Restrictions:</b><br>
<ul>
<li>Cannot delete system groups (admin, everyone)</li>
<li>Requires admin privileges</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteUserGroup" method="delete" path="/userGroups/{groupId}" -->
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

    res, err := s.UserGroups.Delete(ctx, "507f1f77bcf86cd799439011")
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
| `groupID`                                                | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the user group to delete            | 507f1f77bcf86cd799439011                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteUserGroupResponse](../../models/operations/deleteusergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddUsers

Add one or more users to a user group.<br><br>
<b>Behavior:</b><br>
<ul>
<li>Users already in the group are silently skipped</li>
<li>Invalid user IDs are reported in the response</li>
<li>Operation is atomic - all valid users are added together</li>
</ul>
<b>Validation:</b><br>
<ul>
<li>All user IDs must be valid MongoDB ObjectIds</li>
<li>Users must belong to the same organization</li>
<li>Group must exist and not be deleted</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="addUsersToGroup" method="post" path="/userGroups/add-users" -->
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

    res, err := s.UserGroups.AddUsers(ctx, operations.AddUsersToGroupRequest{
        GroupID: "507f1f77bcf86cd799439011",
        UserIds: []string{
            "507f1f77bcf86cd799439012",
            "507f1f77bcf86cd799439013",
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
| `request`                                                                              | [operations.AddUsersToGroupRequest](../../models/operations/adduserstogrouprequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.AddUsersToGroupResponse](../../models/operations/adduserstogroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RemoveUsers

Remove one or more users from a user group.<br><br>
<b>Behavior:</b><br>
<ul>
<li>Users not in the group are silently skipped</li>
<li>Operation is atomic - all specified users are removed together</li>
</ul>
<b>Restrictions:</b><br>
<ul>
<li>Cannot remove users from "everyone" group type</li>
<li>Cannot remove the last admin from admin group</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="removeUsersFromGroup" method="post" path="/userGroups/remove-users" -->
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

    res, err := s.UserGroups.RemoveUsers(ctx, operations.RemoveUsersFromGroupRequest{
        GroupID: "507f1f77bcf86cd799439011",
        UserIds: []string{
            "507f1f77bcf86cd799439012",
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
| `request`                                                                                        | [operations.RemoveUsersFromGroupRequest](../../models/operations/removeusersfromgrouprequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveUsersFromGroupResponse](../../models/operations/removeusersfromgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetUserGroups

Retrieve all user groups that a specific user belongs to.<br><br>
<b>Response Details:</b><br>
<ul>
<li>Includes all group types (admin, standard, everyone, custom)</li>
<li>Returns group metadata for each membership</li>
</ul>
<b>Use Cases:</b><br>
<ul>
<li>Displaying user's group memberships in profile</li>
<li>Access control checks</li>
<li>Permission inheritance calculations</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getGroupsForUser" method="get" path="/userGroups/users/{userId}" -->
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

    res, err := s.UserGroups.GetUserGroups(ctx, "507f1f77bcf86cd799439012")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the user                            | 507f1f77bcf86cd799439012                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetGroupsForUserResponse](../../models/operations/getgroupsforuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetInGroup

Retrieve all users that belong to a specific user group.


### Example Usage

<!-- UsageSnippet language="go" operationID="getUsersInGroup" method="get" path="/userGroups/{groupId}/users" -->
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

    res, err := s.UserGroups.GetInGroup(ctx, "<id>")
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
| `groupID`                                                | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the user group                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetUsersInGroupResponse](../../models/operations/getusersingroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetStatistics

Retrieve statistics for all user groups including member counts.


### Example Usage

<!-- UsageSnippet language="go" operationID="getGroupStatistics" method="get" path="/userGroups/stats/list" -->
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

    res, err := s.UserGroups.GetStatistics(ctx)
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

**[*operations.GetGroupStatisticsResponse](../../models/operations/getgroupstatisticsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |