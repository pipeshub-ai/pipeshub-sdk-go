# GetConnectorRegistryResponseBody

Connector registry retrieved


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Success`                                                                                | `bool`                                                                                   | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `Connectors`                                                                             | [][components.ConnectorRegistryEntry](../../models/components/connectorregistryentry.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `Pagination`                                                                             | [components.ConnectorPagination](../../models/components/connectorpagination.md)         | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `RegistryCountsByScope`                                                                  | [*operations.RegistryCountsByScope](../../models/operations/registrycountsbyscope.md)    | :heavy_minus_sign:                                                                       | How many connector types each scope offers                                               |