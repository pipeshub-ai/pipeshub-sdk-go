# SearchArchivedConversationsPagination


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Page`                                            | `*int64`                                          | :heavy_minus_sign:                                | Current page number                               |
| `Limit`                                           | `*int64`                                          | :heavy_minus_sign:                                | Items per page                                    |
| `TotalCount`                                      | `*int64`                                          | :heavy_minus_sign:                                | Total matches across assistant and agent archives |
| `TotalPages`                                      | `*int64`                                          | :heavy_minus_sign:                                | Total pages at the current limit                  |
| `HasNextPage`                                     | `*bool`                                           | :heavy_minus_sign:                                | Whether a next page exists                        |
| `HasPrevPage`                                     | `*bool`                                           | :heavy_minus_sign:                                | Whether a previous page exists                    |