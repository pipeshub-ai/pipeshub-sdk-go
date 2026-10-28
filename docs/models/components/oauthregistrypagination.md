# OAuthRegistryPagination

Pagination block of `GET /oauth/registry` and `GET /oauth/{connectorType}`


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `Page`                                          | `*int64`                                        | :heavy_minus_sign:                              | Current page number                             |
| `Limit`                                         | `*int64`                                        | :heavy_minus_sign:                              | Items per page                                  |
| `Search`                                        | optionalnullable.OptionalNullable[`string`]     | :heavy_minus_sign:                              | The search term that was applied, or null       |
| `TotalCount`                                    | `*int64`                                        | :heavy_minus_sign:                              | Number of items that match                      |
| `TotalPages`                                    | `*int64`                                        | :heavy_minus_sign:                              | N/A                                             |
| `HasPrev`                                       | `*bool`                                         | :heavy_minus_sign:                              | N/A                                             |
| `HasNext`                                       | `*bool`                                         | :heavy_minus_sign:                              | N/A                                             |
| `PrevPage`                                      | optionalnullable.OptionalNullable[`int64`]      | :heavy_minus_sign:                              | Previous page number, or null on the first page |
| `NextPage`                                      | optionalnullable.OptionalNullable[`int64`]      | :heavy_minus_sign:                              | Next page number, or null on the last page      |