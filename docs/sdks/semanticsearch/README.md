# SemanticSearch

## Overview

### Available Operations

* [Search](#search) - Perform semantic search
* [SearchHistory](#searchhistory) - Get search history
* [DeleteSearchHistory](#deletesearchhistory) - Clear all search history
* [GetSearchByID](#getsearchbyid) - Get search by ID
* [DeleteSearchByID](#deletesearchbyid) - Delete search by ID
* [ArchiveSearch](#archivesearch) - Archive a search
* [UnarchiveSearch](#unarchivesearch) - Unarchive a search

### [Sample App](https://github.com/pipeshub-ai/notebook/tree/main/sample_apps/sdksampleapps/go/enterprise_search)

Runnable Go programs under `semantic_search/` demonstrate `SemanticSearch.Search` with knowledge-base or connector filters. Clone the [notebook](https://github.com/pipeshub-ai/notebook) repo, copy `.env.example` to `.env` in `enterprise_search/`, set your PipesHub URL, credentials, and `KB_ID` / `CONNECTOR_ID`, then from that directory:

```bash
go run ./semantic_search/knowledgebase .env
go run ./semantic_search/connector .env
```

Each program prints the top hits as `Record / ID / Chunk` blocks; edit the `Query` field in the corresponding `main.go` to try different searches.

## Search

Run a semantic search across your organization's knowledge base.
Matching is meaning-based, so relevant results surface even when
the wording differs from the query.

Use optional `filters` to narrow the scope:

- `filters.apps` — restrict to specific connector apps (for
  example Google Drive or Confluence).
- `filters.kb` — restrict to specific knowledge bases.

The response returns a `searchId` for the persisted search along
with ranked matches, each carrying a relevance score and the
source document's metadata. Past searches can be retrieved via
`GET /search`.


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

    res, err := s.SemanticSearch.Search(ctx, components.SemanticSearchRequest{
        Query: "API documentation examples",
        Filters: &components.Filters{
            Apps: []string{
                "drive",
            },
        },
        Limit: pipeshub.Pointer[int64](20),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SemanticSearchExecuteResponse != nil {
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

    res, err := s.SemanticSearch.Search(ctx, components.SemanticSearchRequest{
        Query: "company vacation policy",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SemanticSearchExecuteResponse != nil {
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

## SearchHistory

Retrieve the authenticated user's persisted search history.

Returns searches the user owns along with searches shared with them,
scoped to the caller's organization. Archived and deleted entries are
excluded. Citation references on this endpoint are returned as raw
identifier strings; use `GET /search/{searchId}` to fetch a single
search with its citations fully expanded.

Pagination defaults to `page=1, limit=20` (maximum `limit` is 100).
Results are sorted by most recent activity by default.


### Example Usage

<!-- UsageSnippet language="go" operationID="searchHistory" method="get" path="/search" -->
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

    res, err := s.SemanticSearch.SearchHistory(ctx, operations.SearchHistoryRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SemanticSearchHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.SearchHistoryRequest](../../models/operations/searchhistoryrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.SearchHistoryResponse](../../models/operations/searchhistoryresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.SearchHistoryBadRequestError     | 400                                        | application/json                           |
| apierrors.SearchHistoryUnauthorizedError   | 401                                        | application/json                           |
| apierrors.SearchHistoryForbiddenError      | 403                                        | application/json                           |
| apierrors.SearchHistoryInternalServerError | 500                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

## DeleteSearchHistory

Permanently delete every persisted search row owned by, or shared
with, the authenticated user, along with the citation rows those
searches reference. The action cannot be undone.

Scoped to the caller's org and limited to rows where
`isDeleted: false` and `isArchived: false`. If nothing matches
(including the case where every row is already archived), the
endpoint returns `404` rather than a successful no-op.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSearchHistory" method="delete" path="/search" -->
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

    res, err := s.SemanticSearch.DeleteSearchHistory(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                           | Type                                                                                                                                                                                                                | Required                                                                                                                                                                                                            | Description                                                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                 |
| `search`                                                                                                                                                                                                            | **string*                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                  | Restrict the deletion to rows whose `title` or `messages.content`<br/>matches this case-insensitive substring. Special regex characters<br/>are escaped before the lookup; values over 1000 chars are<br/>rejected with `400`.<br/> |
| `shared`                                                                                                                                                                                                            | [*operations.DeleteSearchHistoryShared](../../models/operations/deletesearchhistoryshared.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                  | Restrict the deletion to rows with this `isShared` value<br/>(`'true'` / `'false'`).<br/>                                                                                                                           |
| `startDate`                                                                                                                                                                                                         | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                  | ISO 8601 lower bound for `createdAt`. Combined with `endDate`<br/>to scope which rows are deleted.<br/>                                                                                                             |
| `endDate`                                                                                                                                                                                                           | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                  | ISO 8601 upper bound for `createdAt`.                                                                                                                                                                               |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |

### Response

**[*operations.DeleteSearchHistoryResponse](../../models/operations/deletesearchhistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSearchByID

Retrieve a previously persisted search by its id, scoped to the
caller's org.

The response body is always an **array** containing zero or one
persisted search document. An unknown id returns an empty array
with a `200` status — callers should check array length rather
than relying on a `404`.


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

    res, err := s.SemanticSearch.GetSearchByID(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.PersistedSemanticSearchEnvelope != nil {
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

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.GetSearchByIDBadRequestError     | 400                                        | application/json                           |
| apierrors.GetSearchByIDUnauthorizedError   | 401                                        | application/json                           |
| apierrors.GetSearchByIDForbiddenError      | 403                                        | application/json                           |
| apierrors.GetSearchByIDNotFoundError       | 404                                        | application/json                           |
| apierrors.GetSearchByIDInternalServerError | 500                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

## DeleteSearchByID

Permanently delete a single persisted search row, plus every
citation row referenced by its `citationIds`. The caller must
either own the row or have it shared with them.

Scoped to the caller's org and limited to rows where
`isDeleted: false` and `isArchived: false`; archived or
already-deleted rows surface as `404`.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSearchById" method="delete" path="/search/{searchId}" -->
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

    res, err := s.SemanticSearch.DeleteSearchByID(ctx, operations.DeleteSearchByIDRequest{
        SearchID: "<value>",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteSearchByIDRequest](../../models/operations/deletesearchbyidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.DeleteSearchByIDResponse](../../models/operations/deletesearchbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ArchiveSearch

Archive a specific search result. Archived searches are hidden
from the default search history view but remain retrievable via
the archive-aware listing endpoints.


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

    res, err := s.SemanticSearch.ArchiveSearch(ctx, "<value>")
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

## UnarchiveSearch

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

    res, err := s.SemanticSearch.UnarchiveSearch(ctx, "<value>")
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