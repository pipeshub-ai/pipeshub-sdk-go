# SemanticSearch

## Overview

Enterprise semantic search across all indexed knowledge with relevance scoring

### Available Operations

* [Retrieve](#retrieve) - Retrieve chunks with citations
* [Search](#search) - Perform semantic search
* [SearchHistory](#searchhistory) - Get search history
* [DeleteSearchHistory](#deletesearchhistory) - Clear all search history
* [GetSearchByID](#getsearchbyid) - Get search by ID
* [DeleteSearchByID](#deletesearchbyid) - Delete search by ID
* [ArchiveSearch](#archivesearch) - Archive a search
* [UnarchiveSearch](#unarchivesearch) - Unarchive a search

## Retrieve

Ranked passages ("chunks") for a query, each with its record and a
citation link, for building your own RAG pipeline or agent on
PipesHub. Results are filtered by the caller's permissions exactly as
`POST /search` is, and take the same `filters` and date bounds.

Unlike `POST /search`, nothing is stored: no search history and no
citation documents, so an application can call this on every request
it serves. The query text is always searched; `rewriteQuery` adds the
model's rewritten and expanded versions.

Nothing to find (no match, no accessible records, nothing indexed
yet) is a `200` with an empty `chunks` list.


### Example Usage

<!-- UsageSnippet language="go" operationID="retrieve" method="post" path="/retrieve" example="simple" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.SemanticSearch.Retrieve(ctx, components.RetrieveRequest{
        Query: "how do we roll back a canary?",
        Filters: &components.Filters{
            Kb: []string{
                "3f2a9c1e-5b7d-4e8a-9c21-7d4e5f6a8b90",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrieveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.RetrieveRequest](../../models/components/retrieverequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.RetrieveResponse](../../models/operations/retrieveresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 422 | application/json        |
| apierrors.ErrorResponse | 500, 503, 504           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Search

Run a semantic search across your organization's knowledge base.
Matching is meaning-based, so relevant results surface even when
the wording differs from the query.

Use optional `filters` to narrow the scope (see `Filters`):

- `filters.apps` — restrict to whole apps: connector instances
  (for example Google Drive or Confluence) and collections.
  `filters.kb` is the older name for collections and means the
  same.
- `filters.recordGroups` — restrict to record groups (a drive, a
  space, a project), each with the groups nested in it.
- `filters.records` — restrict to folders or records, each with
  everything under it.
- `filters.recordsExact` — restrict to records alone, without
  what is under them.

The keys are a union. With any of the last three present, an
absent or empty `apps` / `kb` means no whole app.

The response returns a `searchId` for the persisted search along
with ranked matches, each carrying a relevance score and the
source document's metadata. Past searches can be retrieved via
`GET /search`. With `saveHistory: false` nothing is saved and
there is no `searchId`.

**Scopes:** an OAuth token needs `semantic:write`. The gateway also lets a token holding
only `semantic:read` through, but the search service then refuses it with the same `403`
("Insufficient scope. Required: semantic:write"). Known defect.


### Example Usage: filtered

<!-- UsageSnippet language="go" operationID="search" method="post" path="/search" example="filtered" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
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
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
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
        RewriteQuery: pipeshub.Pointer(false),
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 422 | application/json        |
| apierrors.ErrorResponse | 500, 503, 504           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SearchHistory

Retrieve the authenticated user's persisted search history.

Returns searches the user owns along with searches shared with them,
scoped to the caller's organization. Archived and deleted entries are
excluded. Citation references on this endpoint are returned as raw
identifier strings; use `GET /search/{searchId}` to fetch a single
search with its citations fully expanded.

Pagination defaults to `page=1, limit=20` (maximum `limit` is 100).
Results are sorted by most recent activity by default.

`pagination.totalCount` (and with it `totalPages` and `hasNextPage`)
counts the searches the listing itself holds: the caller's own and those
shared with them in their organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="searchHistory" method="get" path="/search" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
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
| apierrors.BadRequestError                  | 400                                        | application/json                           |
| apierrors.UnauthorizedError                | 401                                        | application/json                           |
| apierrors.ForbiddenError                   | 403                                        | application/json                           |
| apierrors.SearchHistoryInternalServerError | 500                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

## DeleteSearchHistory

Permanently delete every persisted search row owned by, or shared
with, the authenticated user, along with the citation rows those
searches reference. The action cannot be undone. Searches other
users shared with the caller are deleted for their owners too, even
when the caller only has `read` access.

Scoped to the caller's org and limited to rows where
`isArchived: false`. The optional `shared`, `startDate` and
`endDate` query filters narrow what is deleted; `search` is
validated but has no effect. If nothing matches (including the case
where every row is already archived), the endpoint returns `404`
rather than a successful no-op.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSearchHistory" method="delete" path="/search" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
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

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `search`                                                                                                                                                             | `*string`                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                   | Validated, then dropped: a stored search has no title to match, so it never narrows<br/>the deletion. More than 1000 characters, or HTML markup, is refused with `400`.<br/> |
| `shared`                                                                                                                                                             | [*operations.DeleteSearchHistoryShared](../../models/operations/deletesearchhistoryshared.md)                                                                        | :heavy_minus_sign:                                                                                                                                                   | Restrict the deletion to rows with this `isShared` value. `'true'` / `'1'` and<br/>`'false'` / `'0'`, case-insensitive; anything else is refused with `400`.<br/>    |
| `startDate`                                                                                                                                                          | `*string`                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                   | ISO 8601 date or date-time lower bound for `createdAt`. A value JavaScript's `Date`<br/>cannot parse is refused with `400`.<br/>                                     |
| `endDate`                                                                                                                                                            | `*string`                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                   | ISO 8601 date or date-time upper bound for `createdAt`. A value JavaScript's `Date`<br/>cannot parse is refused with `400`.<br/>                                     |
| `opts`                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.DeleteSearchHistoryResponse](../../models/operations/deletesearchhistoryresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetSearchByID

Retrieve a previously persisted search by its id, scoped to the
caller's org. The caller must own the search or have it shared with
them; archived searches are not returned.

The response body is always an **array** containing zero or one
persisted search document. An unknown id, a search the caller
cannot see, an archived search, or one that fails a query filter
all return an empty array with a `200` status — this route never
answers `404`, so callers should check the array length.

The optional `shared`, `startDate` and `endDate` query filters are
applied to the one search; `search` is validated but has no effect.


### Example Usage

<!-- UsageSnippet language="go" operationID="getSearchById" method="get" path="/search/{searchId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.SemanticSearch.GetSearchByID(ctx, operations.GetSearchByIDRequest{
        SearchID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PersistedSemanticSearchEnvelope != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetSearchByIDRequest](../../models/operations/getsearchbyidrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetSearchByIDResponse](../../models/operations/getsearchbyidresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteSearchByID

Permanently delete a single persisted search row, plus every
citation row referenced by its `citationIds`. The caller must
either own the row or have it shared with them: a user the search
was shared with, even with `read` access only, deletes the owner's
search for everyone.

Scoped to the caller's org and limited to rows where
`isArchived: false`; an archived or already-deleted row is `404`.
The optional `shared`, `startDate` and `endDate` query filters must
also match, otherwise `404`; `search` is validated but has no effect.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSearchById" method="delete" path="/search/{searchId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ArchiveSearch

Archive a specific search result. Archived searches are hidden
from the search history and from `GET /search/{searchId}`, and
cannot be shared, unshared or deleted until they are unarchived.

The caller must own the search or have it shared with them: a user
the search was shared with, even with `read` access only, can
archive the owner's search for everyone.


### Example Usage

<!-- UsageSnippet language="go" operationID="archiveSearch" method="patch" path="/search/{searchId}/archive" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.SemanticSearch.ArchiveSearch(ctx, operations.ArchiveSearchRequest{
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ArchiveSearchRequest](../../models/operations/archivesearchrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ArchiveSearchResponse](../../models/operations/archivesearchresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UnarchiveSearch

Restore a previously archived search result back to the active search history.
The owner or any user the search was shared with can unarchive it.


### Example Usage

<!-- UsageSnippet language="go" operationID="unarchiveSearch" method="patch" path="/search/{searchId}/unarchive" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.SemanticSearch.UnarchiveSearch(ctx, operations.UnarchiveSearchRequest{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UnarchiveSearchRequest](../../models/operations/unarchivesearchrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UnarchiveSearchResponse](../../models/operations/unarchivesearchresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |