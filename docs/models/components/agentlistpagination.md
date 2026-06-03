# AgentListPagination

Pagination block returned by `GET /agents`.


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       | Example                                           |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `CurrentPage`                                     | *int64*                                           | :heavy_check_mark:                                | Current 1-based page number.                      | 1                                                 |
| `Limit`                                           | *int64*                                           | :heavy_check_mark:                                | Page size actually applied by the backend.        | 20                                                |
| `TotalItems`                                      | *int64*                                           | :heavy_check_mark:                                | Total number of matching agents across all pages. | 2                                                 |
| `TotalPages`                                      | *int64*                                           | :heavy_check_mark:                                | Total number of pages for the current query.      | 1                                                 |
| `HasNext`                                         | *bool*                                            | :heavy_check_mark:                                | Whether a later page exists.                      | false                                             |
| `HasPrev`                                         | *bool*                                            | :heavy_check_mark:                                | Whether an earlier page exists.                   | false                                             |