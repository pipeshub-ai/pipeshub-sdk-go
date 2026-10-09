# CurrentNode

Node being browsed when `parentId` is in the path; `null` at root.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ID`                                                             | `string`                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `Name`                                                           | `string`                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `NodeType`                                                       | `string`                                                         | :heavy_check_mark:                                               | One of `app`, `recordGroup`, `folder`, `record`.                 |
| `SubType`                                                        | optionalnullable.OptionalNullable[`string`]                      | :heavy_minus_sign:                                               | Connector name or record type when applicable; otherwise `null`. |