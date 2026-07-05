# KnowledgeBase

## Overview

Knowledge base management operations

### Available Operations

* [CreateKnowledgeBase](#createknowledgebase) - Create a new knowledge base
* [ListKnowledgeBases](#listknowledgebases) - List all knowledge bases
* [GetKnowledgeBase](#getknowledgebase) - Get knowledge base by ID
* [UpdateKnowledgeBase](#updateknowledgebase) - Update knowledge base
* [DeleteKnowledgeBase](#deleteknowledgebase) - Delete knowledge base
* [GetRecordByID](#getrecordbyid) - Get record by ID
* [UpdateRecord](#updaterecord) - Update record
* [DeleteRecord](#deleterecord) - Delete record
* [StreamRecordBuffer](#streamrecordbuffer) - Stream record content
* [CreateFolder](#createfolder) - Create folder
* [UpdateFolder](#updatefolder) - Update folder
* [DeleteFolder](#deletefolder) - Delete folder
* [UploadRecords](#uploadrecords) - Upload files to knowledge base or folder
* [GetUploadLimits](#getuploadlimits) - Get knowledge base upload limits
* [ReindexRecord](#reindexrecord) - Reindex single record
* [ReindexRecordGroup](#reindexrecordgroup) - Reindex record group
* [MoveRecord](#moverecord) - Move record to another location
* [~~GetKnowledgeHubRootNodes~~](#getknowledgehubrootnodes) - Get knowledge hub root nodes :warning: **Deprecated**
* [~~GetKnowledgeHubChildNodes~~](#getknowledgehubchildnodes) - Get knowledge hub child nodes :warning: **Deprecated**

## CreateKnowledgeBase

Create a new knowledge base for organizing and managing documents within your organization.

**Overview:**

A knowledge base is a container for organizing related documents, files, and content. It provides a central location for teams to collaborate on shared information.

**Features:**

- Hierarchical folder structure support
- Role-based access control (OWNER, WRITER, READER)
- Full-text search across all records
- Integration with external connectors (Google Drive, OneDrive, etc.)
- Automatic content indexing for AI-powered search

**Naming Rules:**

- Name must be 1-255 characters
- Special characters and HTML tags are sanitized
- Names don't need to be unique within organization

**Creator Permissions:**

The user creating the KB automatically becomes the OWNER with full administrative rights.


### Example Usage

<!-- UsageSnippet language="go" operationID="createKnowledgeBase" method="post" path="/knowledgeBase" -->
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

    res, err := s.KnowledgeBase.CreateKnowledgeBase(ctx, operations.CreateKnowledgeBaseRequest{
        KbName: "Product Documentation",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KnowledgeBaseCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateKnowledgeBaseRequest](../../models/operations/createknowledgebaserequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateKnowledgeBaseResponse](../../models/operations/createknowledgebaseresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListKnowledgeBases

Retrieve a paginated list of all knowledge bases accessible to the authenticated user.

**Overview:**

Returns knowledge bases where the user has at least READER permission. Results include the user's role for each KB.

**Filtering:**

- **search:** Full-text search on KB names (max 1000 chars)
- **permissions:** Filter by user's role (comma-separated: OWNER, WRITER, READER)

**Sorting Options:**

- `name` — Alphabetical by KB name
- `createdAtTimestamp` — By creation date
- `updatedAtTimestamp` — By last modification
- `userRole` — By permission level

**Performance:**

Uses efficient pagination with limit/offset. For large result sets, use smaller page sizes.

**Query parameters:**

Only `page`, `limit`, `search`, `permissions`, `sortBy`, and `sortOrder` are allowed; unknown query keys are rejected.


### Example Usage

<!-- UsageSnippet language="go" operationID="listKnowledgeBases" method="get" path="/knowledgeBase" -->
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

    res, err := s.KnowledgeBase.ListKnowledgeBases(ctx, operations.ListKnowledgeBasesRequest{
        Permissions: pipeshub.Pointer("OWNER,ORGANIZER,WRITER"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAllKnowledgeBaseResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListKnowledgeBasesRequest](../../models/operations/listknowledgebasesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListKnowledgeBasesResponse](../../models/operations/listknowledgebasesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetKnowledgeBase

Retrieve detailed information about a specific knowledge base.

**Overview:**

Returns complete KB metadata including name, timestamps, root-level folders, and the requesting user's role.

**Access Control:**

User must have at least READER permission to view KB details.


### Example Usage

<!-- UsageSnippet language="go" operationID="getKnowledgeBase" method="get" path="/knowledgeBase/{kbId}" -->
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

    res, err := s.KnowledgeBase.GetKnowledgeBase(ctx, "kb_550e8400-e29b-41d4-a716")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetKnowledgeBaseByID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `kbID`                                                   | *string*                                                 | :heavy_check_mark:                                       | Knowledge base ID (non-empty string)                     | 8a095180-2989-4018-b448-70eb75fba1c7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetKnowledgeBaseResponse](../../models/operations/getknowledgebaseresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404           | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateKnowledgeBase

Update a knowledge base's name.

**Required permission:**

User must have one of `OWNER` or `WRITER` on the knowledge base.

**Validation:**

- `kbId` path parameter must be a valid UUID (`updateKBSchema`)
- When provided, `kbName` must be 1–255 characters
- XSS and format-specifier checks are applied to `kbName` in the gateway controller


### Example Usage

<!-- UsageSnippet language="go" operationID="updateKnowledgeBase" method="put" path="/knowledgeBase/{kbId}" -->
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

    res, err := s.KnowledgeBase.UpdateKnowledgeBase(ctx, "<id>", operations.UpdateKnowledgeBaseRequestBody{
        KbName: pipeshub.Pointer("Updated Documentation Hub"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateKnowledgeBaseByID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `kbID`                                                                                                 | *string*                                                                                               | :heavy_check_mark:                                                                                     | Knowledge base ID (UUID)                                                                               | 8a095180-2989-4018-b448-70eb75fba1c7                                                                   |
| `body`                                                                                                 | [operations.UpdateKnowledgeBaseRequestBody](../../models/operations/updateknowledgebaserequestbody.md) | :heavy_check_mark:                                                                                     | Fields to update. `kbName` is optional; an empty object is valid.                                      |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.UpdateKnowledgeBaseResponse](../../models/operations/updateknowledgebaseresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteKnowledgeBase

Permanently delete a knowledge base and all its contents.

**Required permission:**

User must have `OWNER` role on the knowledge base.

**What gets deleted:**

- All folders within the KB
- All records and their indexed content
- All permission grants
- Associated storage files

**Warning:** This action is irreversible. Consider exporting data before deletion.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteKnowledgeBase" method="delete" path="/knowledgeBase/{kbId}" -->
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

    res, err := s.KnowledgeBase.DeleteKnowledgeBase(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteKnowledgeBaseByID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `kbID`                                                   | *string*                                                 | :heavy_check_mark:                                       | Knowledge base ID (non-empty string)                     | 8a095180-2989-4018-b448-70eb75fba1c7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteKnowledgeBaseResponse](../../models/operations/deleteknowledgebaseresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404           | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetRecordByID

Retrieve detailed information about a specific record.

**Overview:**

Returns complete record metadata including name, type, indexing status, storage information, and version history.

**File conversion:**

Use the optional `convertTo` parameter to request file format conversion (e.g., PDF to text). Supported conversions include PPT to PDF and PPTX to PDF.


### Example Usage

<!-- UsageSnippet language="go" operationID="getRecordById" method="get" path="/knowledgeBase/record/{recordId}" -->
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

    res, err := s.KnowledgeBase.GetRecordByID(ctx, "<id>", pipeshub.Pointer("txt"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetRecordByIDResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                             | Type                                                                                                                  | Required                                                                                                              | Description                                                                                                           | Example                                                                                                               |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                 | :heavy_check_mark:                                                                                                    | The context to use for the request.                                                                                   |                                                                                                                       |
| `recordID`                                                                                                            | *string*                                                                                                              | :heavy_check_mark:                                                                                                    | Record ID                                                                                                             |                                                                                                                       |
| `convertTo`                                                                                                           | **string*                                                                                                             | :heavy_minus_sign:                                                                                                    | Optional format to convert the file to (e.g., PDF to text). Supported conversions include PPT to PDF and PPTX to PDF. | txt                                                                                                                   |
| `opts`                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                              | :heavy_minus_sign:                                                                                                    | The options for this request.                                                                                         |                                                                                                                       |

### Response

**[*operations.GetRecordByIDResponse](../../models/operations/getrecordbyidresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateRecord

Update a record's name and/or file content.

**Overview:**

Allows updating the display name and optionally replacing the file content. Triggers re-indexing when content changes.

**Required permission:**

WRITER or higher

**Updating file content:**

Include a new file in the request to replace the existing content. The file extension must match the original.

**Side effects:**

- Updates `updatedAtTimestamp`
- Increments version if file content changed
- Triggers re-indexing for content changes


### Example Usage

<!-- UsageSnippet language="go" operationID="updateRecord" method="put" path="/knowledgeBase/record/{recordId}" -->
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

    res, err := s.KnowledgeBase.UpdateRecord(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `recordID`                                                                                | *string*                                                                                  | :heavy_check_mark:                                                                        | Record ID                                                                                 |
| `body`                                                                                    | [*operations.UpdateRecordRequestBody](../../models/operations/updaterecordrequestbody.md) | :heavy_minus_sign:                                                                        | Request payload                                                                           |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |

### Response

**[*operations.UpdateRecordResponse](../../models/operations/updaterecordresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteRecord

Permanently delete a record from the knowledge base.

**Required permission:**

WRITER or higher

**What gets deleted:**

- Record metadata
- Associated storage file
- Indexed content and embeddings

**Warning:** This action is irreversible.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteRecord" method="delete" path="/knowledgeBase/record/{recordId}" -->
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

    res, err := s.KnowledgeBase.DeleteRecord(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteRecordResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `recordID`                                               | *string*                                                 | :heavy_check_mark:                                       | Record ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteRecordResponse](../../models/operations/deleterecordresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## StreamRecordBuffer

Stream the binary content of a record's file.

**Overview:**

Returns the raw file content with appropriate `Content-Type` and `Content-Disposition` headers for download or inline viewing.

**Use cases:**

- File downloads
- Inline document preview
- Content extraction pipelines

**Format conversion:**

Use the `convertTo` parameter to convert between formats (e.g. DOCX to PDF).


### Example Usage

<!-- UsageSnippet language="go" operationID="streamRecordBuffer" method="get" path="/knowledgeBase/stream/record/{recordId}" -->
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

    res, err := s.KnowledgeBase.StreamRecordBuffer(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `recordID`                                               | *string*                                                 | :heavy_check_mark:                                       | Record ID                                                |
| `convertTo`                                              | **string*                                                | :heavy_minus_sign:                                       | Target format for conversion                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.StreamRecordBufferResponse](../../models/operations/streamrecordbufferresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| apierrors.ErrorResponse             | 400, 401                            | application/json                    |
| apierrors.Forbidden                 | 403                                 | application/json                    |
| apierrors.StreamRecordErrorResponse | 404, 409                            | application/json                    |
| apierrors.StreamRecordErrorResponse | 500                                 | application/json                    |
| apierrors.APIError                  | 4XX, 5XX                            | \*/\*                               |

## CreateFolder

Create a folder in a knowledge base. Omit `folderId` to create at the KB root;
pass `folderId` as a query parameter to create a nested subfolder inside an
existing parent folder.

**Required permission:** WRITER or higher

**Folder features:**

- Organize records hierarchically
- Support nested subfolders (unlimited depth)
- Inherit parent KB permissions

**Naming rules:**

- 1–255 characters
- XSS protection applied
- Spaces and special characters allowed
- Duplicate names rejected within the same parent (`409`)

**Response:** Returns `id` and `name` for the created folder.


### Example Usage

<!-- UsageSnippet language="go" operationID="createFolder" method="post" path="/knowledgeBase/{kbId}/folder" -->
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

    res, err := s.KnowledgeBase.CreateFolder(ctx, "<id>", operations.CreateFolderRequestBody{
        FolderName: "Project Documents",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FolderCreateResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `kbID`                                                                                   | *string*                                                                                 | :heavy_check_mark:                                                                       | Knowledge base ID                                                                        |
| `body`                                                                                   | [operations.CreateFolderRequestBody](../../models/operations/createfolderrequestbody.md) | :heavy_check_mark:                                                                       | Request payload                                                                          |
| `folderID`                                                                               | **string*                                                                                | :heavy_minus_sign:                                                                       | Parent folder ID. Omit to create at the knowledge base root.                             |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateFolderResponse](../../models/operations/createfolderresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateFolder

Rename a folder.

**Required permission:** WRITER or higher


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

    res, err := s.KnowledgeBase.UpdateFolder(ctx, "<id>", "<id>", operations.UpdateFolderRequestBody{
        FolderName: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FolderUpdateResponseSchema != nil {
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteFolder

Delete a folder and all its contents.

**Required permission:** WRITER or higher

**Cascade delete:**

All subfolders and records within will be permanently deleted.

**Warning:** This action is irreversible.


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

    res, err := s.KnowledgeBase.DeleteFolder(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.FolderDeleteResponseSchema != nil {
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UploadRecords

Upload one or more files to a knowledge base root or to a specific folder.

**Overview**

Batch upload multiple files in a single request. Each file becomes a new record with automatic content indexing.
Omit the `folderId` query parameter to upload to the KB root; include it to upload into that folder.

**Upload Limits**

- **Max files per request:** 1000
- **Default max file size:** 30MB (configurable via platform settings)
- Use `GET /knowledgeBase/limits` to check current limits

**Supported File Types**

Documents (PDF, DOCX, DOC, XLS, XLSX, PPT, PPTX, TXT, CSV, MD), Images (PNG, JPG, JPEG, SVG, WebP), Web (HTML, HTM), and Google Workspace formats.

**File Metadata**

Use `files_metadata` to provide additional info like file paths and last modified timestamps.

**Versioning**

Set `isVersioned: true` to enable version tracking for uploaded files.

**Streaming response**

This endpoint responds with `Content-Type: text/event-stream`.
The upload and its per-file progress are a single request: the body streams
a `file:succeeded` or `file:failed` event per file
(including files rejected up front for size/type), followed by a final
`done` summary, then closes. See the
`UploadStreamSSEEvent` schema for the event/payload contract.


### Example Usage

<!-- UsageSnippet language="go" operationID="uploadRecords" method="post" path="/knowledgeBase/{kbId}/upload" -->
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

    res, err := s.KnowledgeBase.UploadRecords(ctx, "<id>", operations.UploadRecordsRequestBody{
        Files: []operations.UploadRecordsFile{},
        FilesMetadata: pipeshub.Pointer("[{\"file_path\":\"/docs/report.pdf\",\"last_modified\":\"2024-01-15T10:30:00Z\"}]"),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadStreamSSEEvent != nil {
        defer res.UploadStreamSSEEvent.Close()

        for res.UploadStreamSSEEvent.Next() {
            event := res.UploadStreamSSEEvent.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `kbID`                                                                                     | *string*                                                                                   | :heavy_check_mark:                                                                         | Knowledge base ID                                                                          |
| `body`                                                                                     | [operations.UploadRecordsRequestBody](../../models/operations/uploadrecordsrequestbody.md) | :heavy_check_mark:                                                                         | Request payload                                                                            |
| `folderID`                                                                                 | **string*                                                                                  | :heavy_minus_sign:                                                                         | Target folder ID. Omit to upload to the KB root.                                           |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UploadRecordsResponse](../../models/operations/uploadrecordsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorResponse      | 400, 401, 403, 404, 413, 429 | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## GetUploadLimits

Retrieve current upload constraints for the organization.

**Use case:** Call this before uploads to validate file sizes on the client
side and display appropriate limits to users.


### Example Usage

<!-- UsageSnippet language="go" operationID="getUploadLimits" method="get" path="/knowledgeBase/limits" -->
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

    res, err := s.KnowledgeBase.GetUploadLimits(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadLimitsResponseSchema != nil {
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

**[*operations.GetUploadLimitsResponse](../../models/operations/getuploadlimitsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ReindexRecord

Trigger reindexing for a specific record.

**Overview:**

Reprocesses the record's content to update search indexes and AI embeddings. Useful after content changes or to fix indexing failures.

**Depth parameter:**

Controls processing depth for complex documents (`-1` for full depth, `0`–`100` for limited).

**Status filters:**

Optional `statusFilters` array limits reindex to records in matching indexing states
(e.g. `FAILED`, `AUTO_INDEX_OFF`).


### Example Usage

<!-- UsageSnippet language="go" operationID="reindexRecord" method="post" path="/knowledgeBase/reindex/record/{recordId}" -->
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

    res, err := s.KnowledgeBase.ReindexRecord(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ReIndexRecordResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `recordID`                                                                                  | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `body`                                                                                      | [*components.ReindexRecordRequestBody](../../models/components/reindexrecordrequestbody.md) | :heavy_minus_sign:                                                                          | Request payload                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.ReindexRecordResponse](../../models/operations/reindexrecordresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ReindexRecordGroup

Trigger reindexing for all records in a folder or knowledge base.

**Overview:**

Batch reindex operation for entire containers. The `recordGroupId` can be a folder ID or KB ID.

**Status filters:**

Optional `statusFilters` limit which child records are queued (e.g. failed-only or manual-indexing).


### Example Usage

<!-- UsageSnippet language="go" operationID="reindexRecordGroup" method="post" path="/knowledgeBase/reindex/record-group/{recordGroupId}" -->
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

    res, err := s.KnowledgeBase.ReindexRecordGroup(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ReIndexRecordGroupResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `recordGroupID`                                                                                       | *string*                                                                                              | :heavy_check_mark:                                                                                    | Folder ID or KB ID                                                                                    |
| `body`                                                                                                | [*components.ReindexRecordGroupRequestBody](../../models/components/reindexrecordgrouprequestbody.md) | :heavy_minus_sign:                                                                                    | Request payload                                                                                       |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |

### Response

**[*operations.ReindexRecordGroupResponse](../../models/operations/reindexrecordgroupresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MoveRecord

Move a file or folder record to a different location within the same knowledge base.

Set `newParentId` to a folder ID to move the record into that folder, or `null` to move it to the knowledge base root.

**Required Permission:** OWNER or WRITER


### Example Usage

<!-- UsageSnippet language="go" operationID="moveRecord" method="put" path="/knowledgeBase/{kbId}/record/{recordId}/move" -->
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

    res, err := s.KnowledgeBase.MoveRecord(ctx, "702f8ff0-0a01-4354-b592-eea268f40f25", "<id>", components.KnowledgeBaseMoveRecordRequestBody{
        NewParentID: pipeshub.Pointer("<id>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KnowledgeBaseMoveRecordResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `kbID`                                                                                                         | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Knowledge base UUID                                                                                            |
| `recordID`                                                                                                     | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Record identifier (file or folder)                                                                             |
| `body`                                                                                                         | [components.KnowledgeBaseMoveRecordRequestBody](../../models/components/knowledgebasemoverecordrequestbody.md) | :heavy_check_mark:                                                                                             | Target location for the record                                                                                 |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.MoveRecordResponse](../../models/operations/moverecordresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ~~GetKnowledgeHubRootNodes~~

Returns root-level nodes (connector apps and Collection apps) or, when
filters or search are applied, a flat list of matching nodes across the
entire knowledge hub tree.

**Overview**

The Knowledge Hub provides a unified view across all knowledge sources:
- **Collection** — locally uploaded knowledge bases (`origin: COLLECTION`)
- **Connector app** — external connector instances such as Google Drive,
  Slack, Confluence, Jira (`origin: CONNECTOR`)

Use this endpoint to build file-browser UIs and sidebar navigation trees.

**Browsing vs. searching**

When no filters or search query are provided, only top-level app nodes
are returned. Adding `nodeTypes`, `q`, or other filter params triggers a
search across the full tree, returning matching nodes regardless of depth.

For children of a specific node, use
`GET /knowledgeBase/knowledge-hub/nodes/{parentType}/{parentId}`.

**Pagination and sorting**

Results are always paginated. Default sort is `updatedAt` descending.
The `pagination` object in the response contains `hasNext` / `hasPrev`
flags suitable for infinite-scroll or page-based navigation.

**Expanding the response**

Use the `include` parameter to request additional sections:
- `availableFilters` — adds `filters.available` with all filter options
- `counts` — adds a `counts` summary broken down by node type
- `breadcrumbs` — adds the breadcrumb trail (empty at root level)
- `permissions` — adds the caller's permission flags

**Access control**

Requires a valid bearer token. For OAuth tokens the `kb:read` scope
must be present; regular JWT bearer tokens pass through without scope
enforcement.


> :warning: **DEPRECATED**: Use the Knowledge Base API instead. This grouping will be removed in a future release.

### Example Usage

<!-- UsageSnippet language="go" operationID="getKnowledgeHubRootNodes" method="get" path="/knowledgeBase/knowledge-hub/nodes" example="root_apps" -->
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

    res, err := s.KnowledgeBase.GetKnowledgeHubRootNodes(ctx, operations.GetKnowledgeHubRootNodesRequest{
        Q: pipeshub.Pointer("quarterly report"),
        NodeTypes: pipeshub.Pointer("app,recordGroup"),
        RecordTypes: pipeshub.Pointer("FILE,CONFLUENCE_PAGE"),
        Origins: pipeshub.Pointer("CONNECTOR"),
        ConnectorIds: pipeshub.Pointer("f3a4b5b6-5b6c-4e85-9097-3202cfe696fc"),
        IndexingStatus: pipeshub.Pointer("COMPLETED,FAILED"),
        CreatedAt: pipeshub.Pointer("gte:1700000000000,lte:1710000000000"),
        UpdatedAt: pipeshub.Pointer("gte:1700000000000,lte:1710000000000"),
        Size: pipeshub.Pointer("gte:0,lte:10485760"),
        Include: pipeshub.Pointer("availableFilters,counts"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KnowledgeHubNodesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetKnowledgeHubRootNodesRequest](../../models/operations/getknowledgehubrootnodesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetKnowledgeHubRootNodesResponse](../../models/operations/getknowledgehubrootnodesresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.GetKnowledgeHubRootNodesBadRequestError     | 400                                                   | application/json                                      |
| apierrors.GetKnowledgeHubRootNodesUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.GetKnowledgeHubRootNodesForbiddenError      | 403                                                   | application/json                                      |
| apierrors.GetKnowledgeHubRootNodesInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## ~~GetKnowledgeHubChildNodes~~

Returns the children of a specific node in the knowledge hub tree.
Use this endpoint to drill down into Collections, connector app
hierarchies, folders, and record groups.

**Navigation hierarchy**

The typical drill-down path is:
1. Root apps (`GET /knowledgeBase/knowledge-hub/nodes`)
2. Record groups / folders within an app (`parentType=app`)
3. Records within a record group (`parentType=recordGroup`)
4. Sub-records or attachments within a record (`parentType=record`)

**Parent identification**

- `parentType` must be one of: `app`, `recordGroup`, `folder`, `record`
- `parentId` is either a standard UUID or the Collection app sentinel
  `knowledgeBase_<orgId>` (e.g. `knowledgeBase_org123`)

**Filtering and searching**

All query-param filters from the root endpoint are available here and
operate within the scope of the parent node's subtree. When `q` is
provided, the search spans all descendants of the parent node.

**Response extras**

When `include=breadcrumbs` is set, the response contains a
`breadcrumbs` array tracing the path from the root to the current
node. The `currentNode` and `parentNode` objects are always populated
for non-root requests.

**Access control**

Requires a valid bearer token. For OAuth tokens the `kb:read` scope
must be present; regular JWT bearer tokens pass through without scope
enforcement.


> :warning: **DEPRECATED**: Use the Knowledge Base API instead. This grouping will be removed in a future release.

### Example Usage

<!-- UsageSnippet language="go" operationID="getKnowledgeHubChildNodes" method="get" path="/knowledgeBase/knowledge-hub/nodes/{parentType}/{parentId}" example="collection_record_groups" -->
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

    res, err := s.KnowledgeBase.GetKnowledgeHubChildNodes(ctx, operations.GetKnowledgeHubChildNodesRequest{
        ParentType: operations.ParentTypeApp,
        ParentID: "<id>",
        Q: pipeshub.Pointer("quarterly report"),
        NodeTypes: pipeshub.Pointer("recordGroup"),
        RecordTypes: pipeshub.Pointer("FILE,CONFLUENCE_PAGE"),
        Origins: pipeshub.Pointer("CONNECTOR"),
        ConnectorIds: pipeshub.Pointer("f3a4b5b6-5b6c-4e85-9097-3202cfe696fc"),
        IndexingStatus: pipeshub.Pointer("COMPLETED,FAILED"),
        CreatedAt: pipeshub.Pointer("gte:1700000000000,lte:1710000000000"),
        UpdatedAt: pipeshub.Pointer("gte:1700000000000,lte:1710000000000"),
        Size: pipeshub.Pointer("gte:0,lte:10485760"),
        Include: pipeshub.Pointer("breadcrumbs,availableFilters"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KnowledgeHubNodesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetKnowledgeHubChildNodesRequest](../../models/operations/getknowledgehubchildnodesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetKnowledgeHubChildNodesResponse](../../models/operations/getknowledgehubchildnodesresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.GetKnowledgeHubChildNodesBadRequestError    | 400                                                   | application/json                                      |
| apierrors.GetKnowledgeHubRootNodesUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.GetKnowledgeHubRootNodesForbiddenError      | 403                                                   | application/json                                      |
| apierrors.GetKnowledgeHubChildNodesNotFoundError      | 404                                                   | application/json                                      |
| apierrors.GetKnowledgeHubRootNodesInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |