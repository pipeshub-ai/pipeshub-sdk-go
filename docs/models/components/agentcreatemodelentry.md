# AgentCreateModelEntry


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ModelKey`                                                            | `string`                                                              | :heavy_check_mark:                                                    | Surrounding whitespace is trimmed; a blank key is refused with `400`. |
| `ModelName`                                                           | `*string`                                                             | :heavy_minus_sign:                                                    | N/A                                                                   |
| `Provider`                                                            | `*string`                                                             | :heavy_minus_sign:                                                    | N/A                                                                   |
| `IsReasoning`                                                         | `*bool`                                                               | :heavy_minus_sign:                                                    | N/A                                                                   |