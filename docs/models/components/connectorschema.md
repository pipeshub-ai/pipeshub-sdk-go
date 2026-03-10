# ConnectorSchema

Schema definition for configuring a connector type


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ConnectorType`                                                     | **string*                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `AuthSchema`                                                        | [*components.AuthSchema](../../models/components/authschema.md)     | :heavy_minus_sign:                                                  | JSON Schema for authentication fields                               |
| `SyncSchema`                                                        | [*components.SyncSchema](../../models/components/syncschema.md)     | :heavy_minus_sign:                                                  | JSON Schema for sync configuration                                  |
| `FilterSchema`                                                      | [*components.FilterSchema](../../models/components/filterschema.md) | :heavy_minus_sign:                                                  | JSON Schema for filter options                                      |
| `RequiredFields`                                                    | []*string*                                                          | :heavy_minus_sign:                                                  | Required field names                                                |