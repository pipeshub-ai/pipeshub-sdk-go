# ConnectorPagination


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Page`                                                | `int64`                                               | :heavy_check_mark:                                    | N/A                                                   |
| `Limit`                                               | `int64`                                               | :heavy_check_mark:                                    | N/A                                                   |
| `Search`                                              | optionalnullable.OptionalNullable[`string`]           | :heavy_minus_sign:                                    | The search term the list was filtered by              |
| `TotalCount`                                          | `int64`                                               | :heavy_check_mark:                                    | Number of items matching the request across all pages |
| `TotalPages`                                          | `int64`                                               | :heavy_check_mark:                                    | Number of pages at the current `limit`                |
| `HasPrev`                                             | `bool`                                                | :heavy_check_mark:                                    | Whether a page exists before this one                 |
| `HasNext`                                             | `bool`                                                | :heavy_check_mark:                                    | Whether a page exists after this one                  |
| `PrevPage`                                            | optionalnullable.OptionalNullable[`int64`]            | :heavy_minus_sign:                                    | Previous page number, `null` on the first page        |
| `NextPage`                                            | optionalnullable.OptionalNullable[`int64`]            | :heavy_minus_sign:                                    | Next page number, `null` on the last page             |