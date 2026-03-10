# DocumentManagement

## Overview

### Available Operations

* [Download](#download) - Download document

## Download

Get a time-limited signed URL to download the document, or receive the file directly for local storage.<br><br>
<b>Overview:</b><br>
This endpoint generates secure download access to documents. For cloud storage (S3/Azure), it returns a presigned URL. For local storage, it streams the file directly.<br><br>
<b>Download Behavior by Storage:</b><br>
<ul>
<li><b>S3/Azure:</b> Returns presigned URL valid for specified duration</li>
<li><b>Local:</b> Streams file directly with appropriate headers</li>
</ul>
<b>Version Download:</b><br>
Specify the <code>version</code> parameter to download a specific historical version. Only available for versioned documents.<br><br>
<b>URL Expiration:</b><br>
<ul>
<li>Default: 3600 seconds (1 hour)</li>
<li>Configurable via <code>expirationTimeInSeconds</code></li>
<li>Maximum depends on storage provider limits</li>
</ul>
<b>Security:</b><br>
Signed URLs are single-use and time-limited. They can be safely shared for temporary access.


### Example Usage

<!-- UsageSnippet language="go" operationID="downloadDocument" method="get" path="/document/{documentId}/download" -->
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

    res, err := s.DocumentManagement.Download(ctx, "507f1f77bcf86cd799439011", pipeshub.Pointer[int64](2), pipeshub.Pointer[int64](7200))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               | Example                                                                   |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |                                                                           |
| `documentID`                                                              | *string*                                                                  | :heavy_check_mark:                                                        | Document ID (24-character MongoDB ObjectId)                               | 507f1f77bcf86cd799439011                                                  |
| `version`                                                                 | **int64*                                                                  | :heavy_minus_sign:                                                        | Version number to download (0-indexed). Must be less than total versions. | 2                                                                         |
| `expirationTimeInSeconds`                                                 | **int64*                                                                  | :heavy_minus_sign:                                                        | Signed URL validity duration in seconds                                   | 7200                                                                      |
| `opts`                                                                    | [][operations.Option](../../models/operations/option.md)                  | :heavy_minus_sign:                                                        | The options for this request.                                             |                                                                           |

### Response

**[*operations.DownloadDocumentResponse](../../models/operations/downloaddocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |