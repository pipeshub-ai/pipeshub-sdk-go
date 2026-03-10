# Folders

## Overview

Folder organization and management

### Available Operations

* [CreateRoot](#createroot) - Create root folder
* [GetContents](#getcontents) - Get folder contents
* [Update](#update) - Update folder
* [Delete](#delete) - Delete folder
* [GetChildren](#getchildren) - Get folder children (alias for folder contents)
* [Create](#create) - Create subfolder

## CreateRoot

Create a new folder at the root level of a knowledge base.<br><br>
<b>Required Permission:</b> FILEORGANIZER or higher<br><br>
<b>Folder Features:</b><br>
<ul>
<li>Organize records hierarchically</li>
<li>Support nested subfolders</li>
<li>Inherit parent KB permissions</li>
</ul>
<b>Naming Rules:</b><br>
<ul>
<li>1-255 characters</li>
<li>XSS protection applied</li>
<li>Can contain spaces and special characters</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="createRootFolder" method="post" path="/knowledgeBase/{kbId}/folder" -->
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

    res, err := s.Folders.CreateRoot(ctx, "<id>", operations.CreateRootFolderRequestBody{
        FolderName: "Project Documents",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Folder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `kbID`                                                                                           | *string*                                                                                         | :heavy_check_mark:                                                                               | Knowledge base ID                                                                                |
| `body`                                                                                           | [operations.CreateRootFolderRequestBody](../../models/operations/createrootfolderrequestbody.md) | :heavy_check_mark:                                                                               | Request payload                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateRootFolderResponse](../../models/operations/createrootfolderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetContents

Retrieve the contents of a folder including subfolders and records.<br><br>
<b>Overview:</b><br>
Returns paginated list of records within the folder, with same filtering options as KB-level record listing.<br><br>
<b>Navigation:</b><br>
Use this endpoint to browse folder hierarchies. Response includes folder metadata and child items.


### Example Usage

<!-- UsageSnippet language="go" operationID="getFolderContents" method="get" path="/knowledgeBase/{kbId}/folder/{folderId}" -->
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

    res, err := s.Folders.GetContents(ctx, operations.GetFolderContentsRequest{
        KbID: "<id>",
        FolderID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RecordsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetFolderContentsRequest](../../models/operations/getfoldercontentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetFolderContentsResponse](../../models/operations/getfoldercontentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Rename a folder.<br><br>
<b>Required Permission:</b> FILEORGANIZER or higher


### Example Usage

<!-- UsageSnippet language="go" operationID="updateFolder" method="put" path="/knowledgeBase/{kbId}/folder/{folderId}" -->
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

    res, err := s.Folders.Update(ctx, "<id>", "<id>", operations.UpdateFolderRequestBody{
        FolderName: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Folder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `kbID`                                                                                   | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `folderID`                                                                               | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `body`                                                                                   | [operations.UpdateFolderRequestBody](../../models/operations/updatefolderrequestbody.md) | :heavy_check_mark:                                                                       | Request payload                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateFolderResponse](../../models/operations/updatefolderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a folder and all its contents.<br><br>
<b>Required Permission:</b> FILEORGANIZER or higher<br><br>
<b>Cascade Delete:</b><br>
All subfolders and records within will be permanently deleted.<br><br>
<b>Warning:</b> This action is irreversible.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteFolder" method="delete" path="/knowledgeBase/{kbId}/folder/{folderId}" -->
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

    res, err := s.Folders.Delete(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `kbID`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `folderID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteFolderResponse](../../models/operations/deletefolderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetChildren

Retrieve the children (subfolders and records) of a folder.<br><br>
<b>Overview:</b><br>
This is an alias endpoint for <code>/knowledgeBase/{kbId}/folder/{folderId}</code>. Returns paginated list of records within the folder, with same filtering options as KB-level record listing.<br><br>
<b>Navigation:</b><br>
Use this endpoint to browse folder hierarchies. Response includes folder metadata and child items.


### Example Usage

<!-- UsageSnippet language="go" operationID="getFolderChildren" method="get" path="/knowledgeBase/{kbId}/folder/{folderId}/children" -->
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

    res, err := s.Folders.GetChildren(ctx, operations.GetFolderChildrenRequest{
        KbID: "<id>",
        FolderID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RecordsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetFolderChildrenRequest](../../models/operations/getfolderchildrenrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetFolderChildrenResponse](../../models/operations/getfolderchildrenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a nested folder within an existing folder.<br><br>
<b>Required Permission:</b> FILEORGANIZER or higher<br><br>
<b>Nesting:</b><br>
Supports unlimited folder nesting depth for complex organizational structures.


### Example Usage

<!-- UsageSnippet language="go" operationID="createSubfolder" method="post" path="/knowledgeBase/{kbId}/folder/{folderId}/subfolder" -->
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

    res, err := s.Folders.Create(ctx, "<id>", "<id>", operations.CreateSubfolderRequestBody{
        FolderName: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Folder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `kbID`                                                                                         | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `folderID`                                                                                     | *string*                                                                                       | :heavy_check_mark:                                                                             | Parent folder ID                                                                               |
| `body`                                                                                         | [operations.CreateSubfolderRequestBody](../../models/operations/createsubfolderrequestbody.md) | :heavy_check_mark:                                                                             | Request payload                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateSubfolderResponse](../../models/operations/createsubfolderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |