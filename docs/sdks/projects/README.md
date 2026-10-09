# Projects

## Overview

Workspaces that group related assistant and agent conversations under a
shared name, custom instructions, a knowledge scope, and reference
files. Projects can be shared with teammates; project membership only
grants read access to conversations explicitly marked
`projectVisibility: project` — a member never gets access to another
member's private chats. See `Conversations` for the two fields
(`projectId`, `projectVisibility`) that link a conversation to a
project.


### Available Operations

* [SetConversationProject](#setconversationproject) - Link or unlink a conversation to a project
* [SetConversationProjectVisibility](#setconversationprojectvisibility) - Override a conversation's project visibility
* [CreateProject](#createproject) - Create a project
* [ListProjects](#listprojects) - List projects
* [GetProjectByID](#getprojectbyid) - Get a project
* [UpdateProject](#updateproject) - Update a project
* [DeleteProject](#deleteproject) - Delete a project
* [ArchiveProject](#archiveproject) - Archive a project
* [UnarchiveProject](#unarchiveproject) - Unarchive a project
* [PinProject](#pinproject) - Pin a project
* [UnpinProject](#unpinproject) - Unpin a project
* [GetProjectConversations](#getprojectconversations) - List a project's conversations
* [EnsureProjectKnowledgeBase](#ensureprojectknowledgebase) - Ensure (create-if-absent) the project's hidden file Collection
* [ListProjectMembers](#listprojectmembers) - List project members
* [UpsertProjectMembers](#upsertprojectmembers) - Add or update project members
* [RemoveProjectMember](#removeprojectmember) - Remove a project member
* [SetAgentConversationProject](#setagentconversationproject) - Link or unlink an agent conversation to a project
* [SetAgentConversationProjectVisibility](#setagentconversationprojectvisibility) - Override an agent conversation's project visibility

## SetConversationProject

Set (`projectId: <id>`) or clear (`projectId: null`) the project this
conversation belongs to. Initiator-only.

**Access:**

The caller must be the conversation's initiator. Linking to a
non-null `projectId` also requires at least viewer access to that
project (`404` if not visible to the caller — never `403`, to avoid
leaking project existence across an org boundary).

**Visibility on link:**

When linking, `projectVisibility` defaults from the project's
`chatSharing` setting (`members` → `project`, otherwise `private`)
unless the conversation was already `project`-visible, in which case
that is preserved. Use
`PATCH /conversations/{conversationId}/project-visibility` to
override it explicitly. Unlinking (`projectId: null`) always clears
both `projectId` and `projectVisibility`.


### Example Usage

<!-- UsageSnippet language="go" operationID="setConversationProject" method="put" path="/conversations/{conversationId}/project" -->
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

    res, err := s.Projects.SetConversationProject(ctx, "<value>", operations.SetConversationProjectRequestBody{
        ProjectID: pipeshub.Pointer("<value>"),
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
| `conversationID`                                                                                             | `string`                                                                                                     | :heavy_check_mark:                                                                                           | Unique conversation identifier                                                                               |
| `body`                                                                                                       | [operations.SetConversationProjectRequestBody](../../models/operations/setconversationprojectrequestbody.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.SetConversationProjectResponse](../../models/operations/setconversationprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SetConversationProjectVisibility

Explicitly set whether a project-linked conversation is visible to
other members of that project (`project`) or only to its owner
(`private`). Initiator-only. Requires the conversation to already be
linked to a project via
`PUT /conversations/{conversationId}/project`.


### Example Usage

<!-- UsageSnippet language="go" operationID="setConversationProjectVisibility" method="patch" path="/conversations/{conversationId}/project-visibility" -->
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

    res, err := s.Projects.SetConversationProjectVisibility(ctx, "<value>", operations.SetConversationProjectVisibilityRequestBody{
        Visibility: operations.SetConversationProjectVisibilityVisibilityProject,
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `conversationID`                                                                                                                 | `string`                                                                                                                         | :heavy_check_mark:                                                                                                               | Unique conversation identifier                                                                                                   |
| `body`                                                                                                                           | [operations.SetConversationProjectVisibilityRequestBody](../../models/operations/setconversationprojectvisibilityrequestbody.md) | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.SetConversationProjectVisibilityResponse](../../models/operations/setconversationprojectvisibilityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateProject

Create a new project workspace owned by the caller.


### Example Usage

<!-- UsageSnippet language="go" operationID="createProject" method="post" path="/projects" -->
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

    res, err := s.Projects.CreateProject(ctx, components.CreateProjectRequest{
        Name: "<value>",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.CreateProjectRequest](../../models/components/createprojectrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateProjectResponse](../../models/operations/createprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListProjects

Paginated, non-deleted projects visible to the caller, sorted pinned
first then by `lastActivityAt` descending.


### Example Usage

<!-- UsageSnippet language="go" operationID="listProjects" method="get" path="/projects" -->
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

    res, err := s.Projects.ListProjects(ctx, operations.ListProjectsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListProjectsRequest](../../models/operations/listprojectsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListProjectsResponse](../../models/operations/listprojectsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetProjectByID

Requires at least viewer access. Cross-org ids and ids the caller
cannot see both return `404` — never `403` — to avoid leaking
project existence across an org boundary.


### Example Usage

<!-- UsageSnippet language="go" operationID="getProjectById" method="get" path="/projects/{projectId}" -->
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

    res, err := s.Projects.GetProjectByID(ctx, "<value>")
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetProjectByIDResponse](../../models/operations/getprojectbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateProject

Requires editor access. `visibility` and `chatSharing` are
owner-only fields — including either as a non-owner editor
returns `403`.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateProject" method="patch" path="/projects/{projectId}" -->
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

    res, err := s.Projects.UpdateProject(ctx, "<value>", components.UpdateProjectRequest{})
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
| `projectID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `body`                                                                             | [components.UpdateProjectRequest](../../models/components/updateprojectrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpdateProjectResponse](../../models/operations/updateprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteProject

Owner-only soft delete. Unlinks every `chatSessions` row pointing at
this project (clearing `projectId`/`projectVisibility`) before
marking the project deleted, so no conversation is left pointing at
a deleted project; wrapped in a transaction when the deployment's
replica set supports it. Idempotent — deleting an already-deleted
project returns `200` without re-running the unlink step.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteProject" method="delete" path="/projects/{projectId}" -->
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

    res, err := s.Projects.DeleteProject(ctx, "<value>")
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteProjectResponse](../../models/operations/deleteprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ArchiveProject

Requires editor access. Archived projects are hidden from the
default sidebar but their conversations remain reachable directly.


### Example Usage

<!-- UsageSnippet language="go" operationID="archiveProject" method="post" path="/projects/{projectId}/archive" -->
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

    res, err := s.Projects.ArchiveProject(ctx, "<value>")
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ArchiveProjectResponse](../../models/operations/archiveprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UnarchiveProject

Unarchive a project

### Example Usage

<!-- UsageSnippet language="go" operationID="unarchiveProject" method="post" path="/projects/{projectId}/unarchive" -->
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

    res, err := s.Projects.UnarchiveProject(ctx, "<value>")
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UnarchiveProjectResponse](../../models/operations/unarchiveprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PinProject

Requires viewer access. Pin state is owner-scoped on the document in
V1 — pinning affects the project for every viewer, not just the
caller (per-user pins are deferred).


### Example Usage

<!-- UsageSnippet language="go" operationID="pinProject" method="post" path="/projects/{projectId}/pin" -->
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

    res, err := s.Projects.PinProject(ctx, "<value>")
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PinProjectResponse](../../models/operations/pinprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UnpinProject

Unpin a project

### Example Usage

<!-- UsageSnippet language="go" operationID="unpinProject" method="post" path="/projects/{projectId}/unpin" -->
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

    res, err := s.Projects.UnpinProject(ctx, "<value>")
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UnpinProjectResponse](../../models/operations/unpinprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetProjectConversations

Requires viewer access to the project. Returns both chat and agent
sessions (`chatSessions`, discriminated by `sessionType`/`agentKey`)
that the caller may see: rows they own, plus rows with
`projectVisibility: project`. Access to the project is asserted
first, so a private conversation belonging to a *different* project
member never leaks through this endpoint.


### Example Usage

<!-- UsageSnippet language="go" operationID="getProjectConversations" method="get" path="/projects/{projectId}/conversations" -->
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

    res, err := s.Projects.GetProjectConversations(ctx, "<value>", pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20))
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `page`                                                   | `*int64`                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `limit`                                                  | `*int64`                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetProjectConversationsResponse](../../models/operations/getprojectconversationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## EnsureProjectKnowledgeBase

Requires editor access. Lazily creates the project's own hidden
Knowledge Base (`isHidden: true` in Python `POST /api/v1/kb`) the
first time a caller needs to upload a file, and returns its id
either way. Race-safe: concurrent callers converge on one KB —
the losing request's KB is deleted. The hidden KB is excluded from
the Collections sidebar, Knowledge Hub, and unscoped search, but is
always included in this project's own chat scope
(`filters.kb`) and file uploads go through the normal
`POST /knowledgeBase/{kbId}/upload` SSE pipeline afterward.


### Example Usage

<!-- UsageSnippet language="go" operationID="ensureProjectKnowledgeBase" method="post" path="/projects/{projectId}/knowledge-base" -->
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

    res, err := s.Projects.EnsureProjectKnowledgeBase(ctx, "<value>")
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EnsureProjectKnowledgeBaseResponse](../../models/operations/ensureprojectknowledgebaseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListProjectMembers

Requires viewer access.

### Example Usage

<!-- UsageSnippet language="go" operationID="listProjectMembers" method="get" path="/projects/{projectId}/members" -->
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

    res, err := s.Projects.ListProjectMembers(ctx, "<value>")
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListProjectMembersResponse](../../models/operations/listprojectmembersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpsertProjectMembers

Owner-only. Each `principalId` is validated against this org's IAM
service the same way as
`POST /conversations/{conversationId}/share` — a `principalId` with
no matching user returns `400`. Upserts by `principalId`; the owner
is silently skipped if included.


### Example Usage

<!-- UsageSnippet language="go" operationID="upsertProjectMembers" method="put" path="/projects/{projectId}/members" -->
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

    res, err := s.Projects.UpsertProjectMembers(ctx, "<value>", components.ProjectMembersUpsertRequest{
        Members: []components.Member{},
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
| `projectID`                                                                                      | `string`                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `body`                                                                                           | [components.ProjectMembersUpsertRequest](../../models/components/projectmembersupsertrequest.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpsertProjectMembersResponse](../../models/operations/upsertprojectmembersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RemoveProjectMember

Owner-only.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeProjectMember" method="delete" path="/projects/{projectId}/members/{memberUserId}" -->
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

    res, err := s.Projects.RemoveProjectMember(ctx, "<value>", "<value>", operations.PrincipalTypeUser.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `projectID`                                                                           | `string`                                                                              | :heavy_check_mark:                                                                    | N/A                                                                                   |
| `memberUserID`                                                                        | `string`                                                                              | :heavy_check_mark:                                                                    | N/A                                                                                   |
| `principalType`                                                                       | [*operations.PrincipalType](../../models/operations/principaltype.md)                 | :heavy_minus_sign:                                                                    | Whether `memberUserId` identifies a user or a team.<br/>Defaults to `user` when omitted.<br/> |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

### Response

**[*operations.RemoveProjectMemberResponse](../../models/operations/removeprojectmemberresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SetAgentConversationProject

Agent-conversation equivalent of
`PUT /conversations/{conversationId}/project`. Set
(`projectId: <id>`) or clear (`projectId: null`) the project this
agent conversation belongs to. Initiator-only; linking requires at
least viewer access to the target project.


### Example Usage

<!-- UsageSnippet language="go" operationID="setAgentConversationProject" method="put" path="/agents/{agentKey}/conversations/{conversationId}/project" -->
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

    res, err := s.Projects.SetAgentConversationProject(ctx, "<value>", "<value>", operations.SetAgentConversationProjectRequestBody{
        ProjectID: pipeshub.Pointer("<value>"),
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `agentKey`                                                                                                             | `string`                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `conversationID`                                                                                                       | `string`                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `body`                                                                                                                 | [operations.SetAgentConversationProjectRequestBody](../../models/operations/setagentconversationprojectrequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.SetAgentConversationProjectResponse](../../models/operations/setagentconversationprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SetAgentConversationProjectVisibility

Agent-conversation equivalent of
`PATCH /conversations/{conversationId}/project-visibility`.
Initiator-only; requires the conversation to already be linked to a
project.


### Example Usage

<!-- UsageSnippet language="go" operationID="setAgentConversationProjectVisibility" method="patch" path="/agents/{agentKey}/conversations/{conversationId}/project-visibility" -->
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

    res, err := s.Projects.SetAgentConversationProjectVisibility(ctx, "<value>", "<value>", operations.SetAgentConversationProjectVisibilityRequestBody{
        Visibility: operations.SetAgentConversationProjectVisibilityVisibilityPrivate,
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `agentKey`                                                                                                                                 | `string`                                                                                                                                   | :heavy_check_mark:                                                                                                                         | N/A                                                                                                                                        |
| `conversationID`                                                                                                                           | `string`                                                                                                                                   | :heavy_check_mark:                                                                                                                         | N/A                                                                                                                                        |
| `body`                                                                                                                                     | [operations.SetAgentConversationProjectVisibilityRequestBody](../../models/operations/setagentconversationprojectvisibilityrequestbody.md) | :heavy_check_mark:                                                                                                                         | N/A                                                                                                                                        |
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.SetAgentConversationProjectVisibilityResponse](../../models/operations/setagentconversationprojectvisibilityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |