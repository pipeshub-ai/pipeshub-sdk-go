# NavigateKnowledgeGraphResponseSchemaPagination

Pagination envelope for a navigate listing.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Page`                                                   | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `Limit`                                                  | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `Total`                                                  | *int64*                                                  | :heavy_check_mark:                                       | Total children of the current node, ignoring pagination. |
| `HasNext`                                                | *bool*                                                   | :heavy_check_mark:                                       | N/A                                                      |
| `HasPrev`                                                | *bool*                                                   | :heavy_check_mark:                                       | N/A                                                      |