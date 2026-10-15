# ConnectorAuthField

One field of an auth form. Send its value in `config.auth` under `name`.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Name`                                                        | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `DisplayName`                                                 | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `Placeholder`                                                 | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `Description`                                                 | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `FieldType`                                                   | `*string`                                                     | :heavy_minus_sign:                                            | Input type, for example TEXT, PASSWORD, URL, SELECT, CHECKBOX |
| `Required`                                                    | `*bool`                                                       | :heavy_minus_sign:                                            | N/A                                                           |
| `Usage`                                                       | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `DefaultValue`                                                | `any`                                                         | :heavy_minus_sign:                                            | Default value; its type follows `fieldType`                   |
| `Validation`                                                  | map[string]`any`                                              | :heavy_minus_sign:                                            | N/A                                                           |
| `IsSecret`                                                    | `*bool`                                                       | :heavy_minus_sign:                                            | Stored values come back masked                                |
| `Options`                                                     | []`any`                                                       | :heavy_minus_sign:                                            | N/A                                                           |