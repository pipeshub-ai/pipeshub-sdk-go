# ParentNode

Parent of `currentNode` when present; `null` when not applicable,
and for a connected app whose limit leaves that parent out. For
such an app, `breadcrumbs` likewise names only the nodes inside
its limit.



## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `ID`                                             | `string`                                         | :heavy_check_mark:                               | N/A                                              |
| `Name`                                           | `string`                                         | :heavy_check_mark:                               | N/A                                              |
| `NodeType`                                       | `string`                                         | :heavy_check_mark:                               | One of `app`, `recordGroup`, `folder`, `record`. |
| `SubType`                                        | optionalnullable.OptionalNullable[`string`]      | :heavy_minus_sign:                               | N/A                                              |