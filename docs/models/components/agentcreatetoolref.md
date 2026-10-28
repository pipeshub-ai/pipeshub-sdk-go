# AgentCreateToolRef


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Name`                                                                  | `string`                                                                | :heavy_check_mark:                                                      | Surrounding whitespace is trimmed; a blank name is refused with `400`.  |
| `FullName`                                                              | `*string`                                                               | :heavy_minus_sign:                                                      | Surrounding whitespace is trimmed; a blank value is refused with `400`. |
| `Description`                                                           | `*string`                                                               | :heavy_minus_sign:                                                      | N/A                                                                     |