# Upload

## Overview

File upload operations

### Available Operations

* [Files](#files) - Upload files to knowledge base
* [ToFolder](#tofolder) - Upload files to folder
* [GetLimits](#getlimits) - Get upload limits

## Files

Upload one or more files directly to a knowledge base.<br><br>
<b>Overview:</b><br>
Batch upload multiple files in a single request. Each file becomes a new record in the KB with automatic content indexing.<br><br>
<b>Upload Limits:</b><br>
<ul>
<li><b>Max files per request:</b> 1000</li>
<li><b>Default max file size:</b> 30MB (configurable via platform settings)</li>
<li>Use <code>GET /knowledgeBase/limits</code> to check current limits</li>
</ul>
<b>Supported File Types:</b><br>
Documents (PDF, DOCX, TXT), Images (PNG, JPG), Videos (MP4), and more.<br><br>
<b>File Metadata:</b><br>
Use <code>files_metadata</code> to provide additional info like file paths and last modified timestamps.<br><br>
<b>Versioning:</b><br>
Set <code>isVersioned: true</code> to enable version tracking for uploaded files.


### Example Usage

<!-- UsageSnippet language="go" operationID="uploadRecordsToKB" method="post" path="/knowledgeBase/{kbId}/upload" -->
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

    res, err := s.Upload.Files(ctx, "<id>", operations.UploadRecordsToKBRequestBody{
        Files: []operations.UploadRecordsToKBFile{
            operations.UploadRecordsToKBFile{
                FileName: "example.file",
                Content: example,
            },
        },
        FilesMetadata: pipeshub.Pointer("[{\"file_path\":\"/docs/report.pdf\",\"last_modified\":\"2024-01-15T10:30:00Z\"}]"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `kbID`                                                                                             | *string*                                                                                           | :heavy_check_mark:                                                                                 | Knowledge base ID                                                                                  |
| `body`                                                                                             | [operations.UploadRecordsToKBRequestBody](../../models/operations/uploadrecordstokbrequestbody.md) | :heavy_check_mark:                                                                                 | Request payload                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UploadRecordsToKBResponse](../../models/operations/uploadrecordstokbresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ToFolder

Upload files directly to a specific folder within a knowledge base.<br><br>
<b>Same as KB upload</b> but files are placed in the specified folder instead of KB root.


### Example Usage

<!-- UsageSnippet language="go" operationID="uploadRecordsToFolder" method="post" path="/knowledgeBase/{kbId}/folder/{folderId}/upload" -->
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

    res, err := s.Upload.ToFolder(ctx, "<id>", "<id>", operations.UploadRecordsToFolderRequestBody{
        Files: []operations.UploadRecordsToFolderFile{
            operations.UploadRecordsToFolderFile{
                FileName: "example.file",
                Content: example,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `kbID`                                                                                                     | *string*                                                                                                   | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `folderID`                                                                                                 | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Target folder ID                                                                                           |
| `body`                                                                                                     | [operations.UploadRecordsToFolderRequestBody](../../models/operations/uploadrecordstofolderrequestbody.md) | :heavy_check_mark:                                                                                         | Request payload                                                                                            |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UploadRecordsToFolderResponse](../../models/operations/uploadrecordstofolderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetLimits

Retrieve current upload constraints for the organization.<br><br>
<b>Use Case:</b><br>
Call this before uploads to validate file sizes on the client side and display appropriate limits to users.


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

    res, err := s.Upload.GetLimits(ctx)
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

**[*operations.GetUploadLimitsResponse](../../models/operations/getuploadlimitsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |