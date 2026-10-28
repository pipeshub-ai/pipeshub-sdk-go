# GetFilterFieldOptionsResponseBody

Filter options


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Success`                                                      | `bool`                                                         | :heavy_check_mark:                                             | N/A                                                            |
| `Options`                                                      | [][operations.OptionObj](../../models/operations/optionobj.md) | :heavy_check_mark:                                             | N/A                                                            |
| `Page`                                                         | `*int64`                                                       | :heavy_minus_sign:                                             | N/A                                                            |
| `Limit`                                                        | `*int64`                                                       | :heavy_minus_sign:                                             | N/A                                                            |
| `HasMore`                                                      | `*bool`                                                        | :heavy_minus_sign:                                             | N/A                                                            |
| `Cursor`                                                       | optionalnullable.OptionalNullable[`string`]                    | :heavy_minus_sign:                                             | Pass back to get the next page                                 |
| `Message`                                                      | optionalnullable.OptionalNullable[`string`]                    | :heavy_minus_sign:                                             | Informational message, only when the connector sets one        |