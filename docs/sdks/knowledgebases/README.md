# KnowledgeBases

## Overview

### Available Operations

* [Create](#create) - Create a new knowledge base
* [List](#list) - List all knowledge bases
* [Get](#get) - Get knowledge base by ID
* [Update](#update) - Update knowledge base
* [Delete](#delete) - Delete knowledge base
* [ReindexFailedRecords](#reindexfailedrecords) - Reindex failed records for connector
* [MoveRecord](#moverecord) - Move record to another location
* [GetRootNodes](#getrootnodes) - Get knowledge hub root nodes
* [GetChildNodes](#getchildnodes) - Get knowledge hub child nodes

## Create

Create a new knowledge base for organizing and managing documents within your organization.<br><br>
<b>Overview:</b><br>
A knowledge base is a container for organizing related documents, files, and content. It provides a central location for teams to collaborate on shared information.<br><br>
<b>Features:</b><br>
<ul>
<li>Hierarchical folder structure support</li>
<li>Role-based access control (OWNER, ORGANIZER, WRITER, READER, etc.)</li>
<li>Full-text search across all records</li>
<li>Integration with external connectors (Google Drive, OneDrive, etc.)</li>
<li>Automatic content indexing for AI-powered search</li>
</ul>
<b>Naming Rules:</b><br>
<ul>
<li>Name must be 1-255 characters</li>
<li>Special characters and HTML tags are sanitized</li>
<li>Names don't need to be unique within organization</li>
</ul>
<b>Creator Permissions:</b><br>
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

    res, err := s.KnowledgeBases.Create(ctx, operations.CreateKnowledgeBaseRequest{
        KbName: "Product Documentation",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KnowledgeBase != nil {
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Retrieve a paginated list of all knowledge bases accessible to the authenticated user.<br><br>
<b>Overview:</b><br>
Returns knowledge bases where the user has at least READER permission. Results include the user's role for each KB.<br><br>
<b>Filtering:</b><br>
<ul>
<li><b>search:</b> Full-text search on KB names (max 1000 chars)</li>
<li><b>permissions:</b> Filter by user's role (comma-separated: OWNER,WRITER,READER)</li>
</ul>
<b>Sorting Options:</b><br>
<ul>
<li><code>name</code> - Alphabetical by KB name</li>
<li><code>createdAtTimestamp</code> - By creation date</li>
<li><code>updatedAtTimestamp</code> - By last modification</li>
<li><code>userRole</code> - By permission level</li>
</ul>
<b>Performance:</b><br>
Uses efficient pagination with limit/offset. For large result sets, use smaller page sizes.


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

    res, err := s.KnowledgeBases.List(ctx, operations.ListKnowledgeBasesRequest{
        Permissions: pipeshub.Pointer("OWNER,ORGANIZER,WRITER"),
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
| `request`                                                                                    | [operations.ListKnowledgeBasesRequest](../../models/operations/listknowledgebasesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListKnowledgeBasesResponse](../../models/operations/listknowledgebasesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve detailed information about a specific knowledge base.<br><br>
<b>Overview:</b><br>
Returns complete KB metadata including name, timestamps, and the requesting user's role/permissions.<br><br>
<b>Access Control:</b><br>
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

    res, err := s.KnowledgeBases.Get(ctx, "kb_550e8400-e29b-41d4-a716")
    if err != nil {
        log.Fatal(err)
    }
    if res.KnowledgeBase != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `kbID`                                                   | *string*                                                 | :heavy_check_mark:                                       | Knowledge base ID                                        | kb_550e8400-e29b-41d4-a716                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetKnowledgeBaseResponse](../../models/operations/getknowledgebaseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a knowledge base's name.<br><br>
<b>Required Permission:</b> OWNER or ORGANIZER<br><br>
<b>Validation:</b><br>
<ul>
<li>Name must be 1-255 characters</li>
<li>XSS protection applied to input</li>
</ul>


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

    res, err := s.KnowledgeBases.Update(ctx, "<id>", operations.UpdateKnowledgeBaseRequestBody{
        KbName: pipeshub.Pointer("Updated Documentation Hub"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KnowledgeBase != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `kbID`                                                                                                 | *string*                                                                                               | :heavy_check_mark:                                                                                     | Knowledge base ID                                                                                      |
| `body`                                                                                                 | [operations.UpdateKnowledgeBaseRequestBody](../../models/operations/updateknowledgebaserequestbody.md) | :heavy_check_mark:                                                                                     | Request payload                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateKnowledgeBaseResponse](../../models/operations/updateknowledgebaseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Permanently delete a knowledge base and all its contents.<br><br>
<b>Required Permission:</b> OWNER only<br><br>
<b>What Gets Deleted:</b><br>
<ul>
<li>All folders within the KB</li>
<li>All records and their indexed content</li>
<li>All permission grants</li>
<li>Associated storage files</li>
</ul>
<b>Warning:</b> This action is irreversible. Consider exporting data before deletion.


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

    res, err := s.KnowledgeBases.Delete(ctx, "<id>")
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
| `kbID`                                                   | *string*                                                 | :heavy_check_mark:                                       | Knowledge base ID                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteKnowledgeBaseResponse](../../models/operations/deleteknowledgebaseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReindexFailedRecords

Trigger reindexing of records that previously failed to index for a specific connector.


### Example Usage

<!-- UsageSnippet language="go" operationID="reindexFailedRecords" method="post" path="/knowledgeBase/reindex-failed/connector" -->
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

    res, err := s.KnowledgeBases.ReindexFailedRecords(ctx, operations.ReindexFailedRecordsRequest{})
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
| `request`                                                                                        | [operations.ReindexFailedRecordsRequest](../../models/operations/reindexfailedrecordsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ReindexFailedRecordsResponse](../../models/operations/reindexfailedrecordsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## MoveRecord

Move a record from one location to another within a knowledge base.


### Example Usage

<!-- UsageSnippet language="go" operationID="moveRecord" method="put" path="/knowledgeBase/{kbId}/record/{recordId}/move" -->
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

    res, err := s.KnowledgeBases.MoveRecord(ctx, "702f8ff0-0a01-4354-b592-eea268f40f25", "<id>", operations.MoveRecordRequestBody{})
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
| `kbID`                                                                               | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `recordID`                                                                           | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `body`                                                                               | [operations.MoveRecordRequestBody](../../models/operations/moverecordrequestbody.md) | :heavy_check_mark:                                                                   | Request payload                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.MoveRecordResponse](../../models/operations/moverecordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetRootNodes

Retrieve root-level nodes for unified knowledge hub browsing.<br><br>
<b>Overview:</b><br>
Provides a unified view across all knowledge sources - KBs, connectors, and apps. Use for building file browser UIs.<br><br>
<b>Node Types:</b><br>
<ul>
<li><b>KB:</b> Knowledge bases</li>
<li><b>CONNECTOR:</b> External connector instances</li>
<li><b>APP:</b> Connected applications</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getKnowledgeHubRootNodes" method="get" path="/knowledgeBase/knowledge-hub/nodes" -->
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

    res, err := s.KnowledgeBases.GetRootNodes(ctx, operations.GetKnowledgeHubRootNodesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetChildNodes

Retrieve child nodes under a specific parent in the knowledge hub tree.<br><br>
<b>Navigation:</b><br>
Use this to drill down into KBs, folders, and connector hierarchies.


### Example Usage

<!-- UsageSnippet language="go" operationID="getKnowledgeHubChildNodes" method="get" path="/knowledgeBase/knowledge-hub/nodes/{parentType}/{parentId}" -->
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

    res, err := s.KnowledgeBases.GetChildNodes(ctx, operations.GetKnowledgeHubChildNodesRequest{
        ParentType: "<value>",
        ParentID: "<id>",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetKnowledgeHubChildNodesRequest](../../models/operations/getknowledgehubchildnodesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetKnowledgeHubChildNodesResponse](../../models/operations/getknowledgehubchildnodesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |