# SemanticSearch

## Overview

### Available Operations

* [Execute](#execute) - Perform semantic search
* [GetHistory](#gethistory) - Get search history
* [DeleteAllHistory](#deleteallhistory) - Clear all search history
* [GetByID](#getbyid) - Get search by ID
* [Delete](#delete) - Delete search by ID
* [Share](#share) - Share a search
* [Unshare](#unshare) - Unshare a search
* [Archive](#archive) - Archive a search
* [Unarchive](#unarchive) - Unarchive a search

## Execute

Execute a semantic search across your organization's knowledge base.<br><br>
<b>Overview:</b><br>
Semantic search uses AI embeddings to find content based on meaning,
not just keyword matching. This enables finding relevant information
even when the exact words differ.<br><br>
<b>How It Works:</b><br>
<ol>
<li>Your query is converted to a vector embedding</li>
<li>The system finds documents with similar semantic meaning</li>
<li>Results are ranked by relevance score</li>
<li>Matching chunks are returned with metadata</li>
</ol>
<b>Filtering:</b><br>
Use filters to narrow your search:
<ul>
<li><code>filters.apps</code>: Limit to specific connector apps (Google Drive, Confluence, etc.)</li>
<li><code>filters.kb</code>: Limit to specific knowledge bases</li>
</ul>
<b>Results:</b><br>
Each result includes:
<ul>
<li>Matching content chunk</li>
<li>Relevance score (0-1, higher is better)</li>
<li>Source document metadata (name, URL, type)</li>
</ul>
<b>Search History:</b><br>
All searches are saved and can be retrieved via <code>GET /search</code>.


### Example Usage: filtered

<!-- UsageSnippet language="go" operationID="search" method="post" path="/search" example="filtered" -->
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

    res, err := s.SemanticSearch.Execute(ctx, components.SemanticSearchRequest{
        Query: "API documentation examples",
        Filters: &components.Filters{
            Apps: []components.AppType{
                components.AppTypeDrive,
            },
        },
        Limit: pipeshub.Pointer[int64](20),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchResult != nil {
        // handle response
    }
}
```
### Example Usage: simple

<!-- UsageSnippet language="go" operationID="search" method="post" path="/search" example="simple" -->
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

    res, err := s.SemanticSearch.Execute(ctx, components.SemanticSearchRequest{
        Query: "company vacation policy",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.SemanticSearchRequest](../../models/components/semanticsearchrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.SearchResponse](../../models/operations/searchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetHistory

Retrieve your search history with pagination.<br><br>
<b>Overview:</b><br>
Returns a list of all searches performed by the authenticated user.
Each entry includes the original query, results, and metadata.<br><br>
<b>Pagination:</b><br>
Use <code>page</code> and <code>limit</code> to navigate through results.


### Example Usage

<!-- UsageSnippet language="go" operationID="searchHistory" method="get" path="/search" -->
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

    res, err := s.SemanticSearch.GetHistory(ctx, pipeshub.Pointer[int64](10), pipeshub.Pointer[int64](1))
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
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | Number of results per page                               |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SearchHistoryResponse](../../models/operations/searchhistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteAllHistory

Delete all search history for the authenticated user.<br><br>
<b>Warning:</b><br>
This action cannot be undone. All saved searches will be permanently removed.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAllSearchHistory" method="delete" path="/search" -->
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

    res, err := s.SemanticSearch.DeleteAllHistory(ctx)
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

**[*operations.DeleteAllSearchHistoryResponse](../../models/operations/deleteallsearchhistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetByID

Retrieve a specific search result by its ID.


### Example Usage

<!-- UsageSnippet language="go" operationID="getSearchById" method="get" path="/search/{searchId}" -->
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

    res, err := s.SemanticSearch.GetByID(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `searchID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique search identifier                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSearchByIDResponse](../../models/operations/getsearchbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a specific search result by its ID.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSearchById" method="delete" path="/search/{searchId}" -->
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

    res, err := s.SemanticSearch.Delete(ctx, "<value>")
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
| `searchID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique search identifier                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteSearchByIDResponse](../../models/operations/deletesearchbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Share

Share a specific search result, making it accessible to other users.


### Example Usage

<!-- UsageSnippet language="go" operationID="shareSearch" method="patch" path="/search/{searchId}/share" -->
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

    res, err := s.SemanticSearch.Share(ctx, "<value>", operations.ShareSearchRequestBody{})
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
| `searchID`                                                                             | *string*                                                                               | :heavy_check_mark:                                                                     | Unique search identifier                                                               |
| `body`                                                                                 | [operations.ShareSearchRequestBody](../../models/operations/sharesearchrequestbody.md) | :heavy_check_mark:                                                                     | Request payload                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ShareSearchResponse](../../models/operations/sharesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Unshare

Revoke sharing for a specific search result, making it private again.


### Example Usage

<!-- UsageSnippet language="go" operationID="unshareSearch" method="patch" path="/search/{searchId}/unshare" -->
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

    res, err := s.SemanticSearch.Unshare(ctx, "<value>", operations.UnshareSearchRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `searchID`                                                                                 | *string*                                                                                   | :heavy_check_mark:                                                                         | Unique search identifier                                                                   |
| `body`                                                                                     | [operations.UnshareSearchRequestBody](../../models/operations/unsharesearchrequestbody.md) | :heavy_check_mark:                                                                         | Request payload                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UnshareSearchResponse](../../models/operations/unsharesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Archive

Archive a specific search result. Archived searches are hidden from the default search history view.


### Example Usage

<!-- UsageSnippet language="go" operationID="archiveSearch" method="patch" path="/search/{searchId}/archive" -->
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

    res, err := s.SemanticSearch.Archive(ctx, "<value>")
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
| `searchID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique search identifier                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ArchiveSearchResponse](../../models/operations/archivesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Unarchive

Restore a previously archived search result back to the active search history.


### Example Usage

<!-- UsageSnippet language="go" operationID="unarchiveSearch" method="patch" path="/search/{searchId}/unarchive" -->
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

    res, err := s.SemanticSearch.Unarchive(ctx, "<value>")
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
| `searchID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique search identifier                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UnarchiveSearchResponse](../../models/operations/unarchivesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |