# ConnectorPagination

Pagination information for connector lists


## Fields

| Field                    | Type                     | Required                 | Description              |
| ------------------------ | ------------------------ | ------------------------ | ------------------------ |
| `Page`                   | **int64*                 | :heavy_minus_sign:       | Current page number      |
| `Limit`                  | **int64*                 | :heavy_minus_sign:       | Items per page           |
| `Total`                  | **int64*                 | :heavy_minus_sign:       | Total number of items    |
| `HasMore`                | **bool*                  | :heavy_minus_sign:       | Whether more pages exist |