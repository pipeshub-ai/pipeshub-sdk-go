# OAuthConfigPagination


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Page`                                      | `int64`                                     | :heavy_check_mark:                          | N/A                                         |
| `Limit`                                     | `int64`                                     | :heavy_check_mark:                          | N/A                                         |
| `Search`                                    | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | N/A                                         |
| `TotalItems`                                | `int64`                                     | :heavy_check_mark:                          | N/A                                         |
| `TotalPages`                                | `int64`                                     | :heavy_check_mark:                          | N/A                                         |
| `HasPrev`                                   | `bool`                                      | :heavy_check_mark:                          | N/A                                         |
| `HasNext`                                   | `bool`                                      | :heavy_check_mark:                          | N/A                                         |