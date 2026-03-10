# Records

## Overview

Record management and operations

### Available Operations

* [GetAll](#getall) - Get all records across knowledge bases
* [GetByKB](#getbykb) - Get records for a knowledge base
* [GetChildren](#getchildren) - Get KB children (alias for records)
* [GetByID](#getbyid) - Get record by ID
* [Update](#update) - Update record
* [Delete](#delete) - Delete record
* [StreamContent](#streamcontent) - Stream record content

## GetAll

Retrieve records from all knowledge bases accessible to the user.<br><br>
<b>Overview:</b><br>
Search and filter records across your entire organization. Useful for global search, reporting, and cross-KB content discovery.<br><br>
<b>Filtering Options:</b><br>
<ul>
<li><b>search:</b> Full-text search in record names</li>
<li><b>recordTypes:</b> FILE, WEBPAGE, EMAIL, MESSAGE, TICKET, etc.</li>
<li><b>origins:</b> UPLOAD or CONNECTOR</li>
<li><b>connectors:</b> Filter by connector source</li>
<li><b>indexingStatus:</b> COMPLETED, FAILED, IN_PROGRESS, etc.</li>
<li><b>dateFrom/dateTo:</b> Filter by creation date range</li>
</ul>
<b>Response Includes:</b><br>
<ul>
<li>Paginated record list</li>
<li>Applied and available filter counts</li>
<li>Pagination metadata</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="getAllRecords" method="get" path="/knowledgeBase/records" -->
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

    res, err := s.Records.GetAll(ctx, operations.GetAllRecordsRequest{
        RecordTypes: pipeshub.Pointer("FILE,WEBPAGE,EMAIL"),
        Connectors: pipeshub.Pointer("GOOGLE_DRIVE,ONEDRIVE"),
        IndexingStatus: pipeshub.Pointer("COMPLETED,FAILED"),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetAllRecordsRequest](../../models/operations/getallrecordsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetAllRecordsResponse](../../models/operations/getallrecordsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetByKB

Retrieve a paginated list of records within a specific knowledge base.<br><br>
<b>Overview:</b><br>
Returns all records (documents, files, content) stored in the specified KB, with powerful filtering and sorting capabilities.<br><br>
<b>Filtering:</b><br>
<ul>
<li><b>search:</b> Search by record name (partial match, max 1000 chars)</li>
<li><b>recordTypes:</b> FILE, WEBPAGE, COMMENT, MESSAGE, EMAIL, TICKET</li>
<li><b>origins:</b> UPLOAD (manual uploads) or CONNECTOR (synced)</li>
<li><b>indexingStatus:</b> Filter by processing state</li>
<li><b>dateFrom/dateTo:</b> Creation date range (Unix timestamps)</li>
</ul>
<b>Sorting:</b><br>
Default sorts by <code>createdAtTimestamp</code> descending (newest first).


### Example Usage

<!-- UsageSnippet language="go" operationID="getKBRecords" method="get" path="/knowledgeBase/{kbId}/records" -->
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

    res, err := s.Records.GetByKB(ctx, operations.GetKBRecordsRequest{
        KbID: "<id>",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetKBRecordsRequest](../../models/operations/getkbrecordsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetKBRecordsResponse](../../models/operations/getkbrecordsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetChildren

Retrieve a paginated list of children (folders and records) within a specific knowledge base.<br><br>
<b>Overview:</b><br>
This is an alias endpoint for <code>/knowledgeBase/{kbId}/records</code>. It returns all direct children of the KB root, including both folders and records.<br><br>
<b>Filtering:</b><br>
<ul>
<li><b>search:</b> Search by record name (partial match, max 1000 chars)</li>
<li><b>recordTypes:</b> FILE, WEBPAGE, COMMENT, MESSAGE, EMAIL, TICKET</li>
<li><b>origins:</b> UPLOAD (manual uploads) or CONNECTOR (synced)</li>
<li><b>indexingStatus:</b> Filter by processing state</li>
<li><b>dateFrom/dateTo:</b> Creation date range (Unix timestamps)</li>
</ul>
<b>Sorting:</b><br>
Default sorts by <code>createdAtTimestamp</code> descending (newest first).


### Example Usage

<!-- UsageSnippet language="go" operationID="getKBChildren" method="get" path="/knowledgeBase/{kbId}/children" -->
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

    res, err := s.Records.GetChildren(ctx, operations.GetKBChildrenRequest{
        KbID: "<id>",
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
| `request`                                                                          | [operations.GetKBChildrenRequest](../../models/operations/getkbchildrenrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetKBChildrenResponse](../../models/operations/getkbchildrenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetByID

Retrieve detailed information about a specific record.<br><br>
<b>Overview:</b><br>
Returns complete record metadata including name, type, indexing status, storage information, and version history.<br><br>
<b>File Conversion:</b><br>
Use the optional <code>convertTo</code> parameter to request file format conversion (e.g., PDF to text).


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

    res, err := s.Records.GetByID(ctx, "<id>", pipeshub.Pointer("txt"))
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
| `recordID`                                               | *string*                                                 | :heavy_check_mark:                                       | Record ID                                                |                                                          |
| `convertTo`                                              | **string*                                                | :heavy_minus_sign:                                       | Optional format to convert the file to                   | txt                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetRecordByIDResponse](../../models/operations/getrecordbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a record's name and/or file content.<br><br>
<b>Overview:</b><br>
Allows updating the display name and optionally replacing the file content. Triggers re-indexing when content changes.<br><br>
<b>Required Permission:</b> WRITER or higher<br><br>
<b>Updating File Content:</b><br>
Include a new file in the request to replace the existing content. The file extension must match the original.<br><br>
<b>Side Effects:</b><br>
<ul>
<li>Updates <code>updatedAtTimestamp</code></li>
<li>Increments version if file content changed</li>
<li>Triggers re-indexing for content changes</li>
</ul>


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

    res, err := s.Records.Update(ctx, "<id>", nil)
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Permanently delete a record from the knowledge base.<br><br>
<b>Required Permission:</b> WRITER or higher<br><br>
<b>What Gets Deleted:</b><br>
<ul>
<li>Record metadata</li>
<li>Associated storage file</li>
<li>Indexed content and embeddings</li>
</ul>
<b>Warning:</b> This action is irreversible.


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

    res, err := s.Records.Delete(ctx, "<id>")
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
| `recordID`                                               | *string*                                                 | :heavy_check_mark:                                       | Record ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteRecordResponse](../../models/operations/deleterecordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## StreamContent

Stream the binary content of a record's file.<br><br>
<b>Overview:</b><br>
Returns the raw file content with appropriate Content-Type and Content-Disposition headers for download or inline viewing.<br><br>
<b>Use Cases:</b><br>
<ul>
<li>File downloads</li>
<li>Inline document preview</li>
<li>Content extraction pipelines</li>
</ul>
<b>Format Conversion:</b><br>
Use <code>convertTo</code> parameter to convert between formats (e.g., DOCX to PDF).


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

    res, err := s.Records.StreamContent(ctx, "<id>", nil)
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |