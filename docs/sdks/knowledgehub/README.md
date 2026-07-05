# ~~KnowledgeHub~~

> [!WARNING]
> This SDK is **DEPRECATED**

## Overview

Unified browse API for root and child nodes (apps, record groups, folders, records) with filtering and search

### Available Operations

* [~~GetKnowledgeHubRootNodes~~](#getknowledgehubrootnodes) - Get knowledge hub root nodes :warning: **Deprecated**
* [~~GetKnowledgeHubChildNodes~~](#getknowledgehubchildnodes) - Get knowledge hub child nodes :warning: **Deprecated**

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


> :warning: **DEPRECATED**: Use the Knowledge Base API instead (sdk.knowledge_base.get_root_nodes / get_child_nodes). This grouping will be removed in a future release.

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

    res, err := s.KnowledgeHub.GetKnowledgeHubRootNodes(ctx, operations.GetKnowledgeHubRootNodesRequest{
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


> :warning: **DEPRECATED**: Use the Knowledge Base API instead (sdk.knowledge_base.get_root_nodes / get_child_nodes). This grouping will be removed in a future release.

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

    res, err := s.KnowledgeHub.GetKnowledgeHubChildNodes(ctx, operations.GetKnowledgeHubChildNodesRequest{
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