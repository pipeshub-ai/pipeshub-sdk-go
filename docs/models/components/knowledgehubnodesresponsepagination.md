# KnowledgeHubNodesResponsePagination


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Page`                                                    | `int64`                                                   | :heavy_check_mark:                                        | Current page (1-indexed).                                 |
| `Limit`                                                   | `int64`                                                   | :heavy_check_mark:                                        | Page size.                                                |
| `TotalItems`                                              | `int64`                                                   | :heavy_check_mark:                                        | N/A                                                       |
| `TotalPages`                                              | `int64`                                                   | :heavy_check_mark:                                        | N/A                                                       |
| `HasNext`                                                 | `bool`                                                    | :heavy_check_mark:                                        | N/A                                                       |
| `HasPrev`                                                 | `bool`                                                    | :heavy_check_mark:                                        | N/A                                                       |
| `NextCursor`                                              | optionalnullable.OptionalNullable[`string`]               | :heavy_minus_sign:                                        | Cursor for the next page; null on the last page.          |
| `PrevCursor`                                              | optionalnullable.OptionalNullable[`string`]               | :heavy_minus_sign:                                        | Cursor for the previous page; null on the first page.     |
| `StartIndex`                                              | optionalnullable.OptionalNullable[`int64`]                | :heavy_minus_sign:                                        | 1-based index of the first item on this page, when known. |
| `EndIndex`                                                | optionalnullable.OptionalNullable[`int64`]                | :heavy_minus_sign:                                        | 1-based index of the last item on this page, when known.  |
| `CurrentPageItems`                                        | optionalnullable.OptionalNullable[`int64`]                | :heavy_minus_sign:                                        | Number of items on this page.                             |