# ConnectorPagination


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Page`                                      | `int64`                                     | :heavy_check_mark:                          | N/A                                         |
| `Limit`                                     | `int64`                                     | :heavy_check_mark:                          | N/A                                         |
| `Search`                                    | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | The search term the list was filtered by    |
| `TotalCount`                                | `int64`                                     | :heavy_check_mark:                          | N/A                                         |
| `TotalPages`                                | `int64`                                     | :heavy_check_mark:                          | N/A                                         |
| `HasPrev`                                   | `bool`                                      | :heavy_check_mark:                          | N/A                                         |
| `HasNext`                                   | `bool`                                      | :heavy_check_mark:                          | N/A                                         |
| `PrevPage`                                  | optionalnullable.OptionalNullable[`int64`]  | :heavy_minus_sign:                          | N/A                                         |
| `NextPage`                                  | optionalnullable.OptionalNullable[`int64`]  | :heavy_minus_sign:                          | N/A                                         |