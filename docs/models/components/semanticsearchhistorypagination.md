# SemanticSearchHistoryPagination

Pagination block emitted by `buildPaginationMetadata` (utils.ts:417).
`totalPages` is `Math.ceil(totalCount / limit)`, so an empty result
has `totalPages: 0`, not `1`.



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Page`             | `int64`            | :heavy_check_mark: | N/A                |
| `Limit`            | `int64`            | :heavy_check_mark: | N/A                |
| `TotalCount`       | `int64`            | :heavy_check_mark: | N/A                |
| `TotalPages`       | `int64`            | :heavy_check_mark: | N/A                |
| `HasNextPage`      | `bool`             | :heavy_check_mark: | N/A                |
| `HasPrevPage`      | `bool`             | :heavy_check_mark: | N/A                |