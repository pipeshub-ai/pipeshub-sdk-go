# CreateConnectorInstanceResponseBody

Connector instance created. It is not syncing yet.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Success`                                                                  | `bool`                                                                     | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `Connector`                                                                | [components.ConnectorCreated](../../models/components/connectorcreated.md) | :heavy_check_mark:                                                         | Summary of a connector instance that was just created                      |                                                                            |
| `Message`                                                                  | `*string`                                                                  | :heavy_minus_sign:                                                         | N/A                                                                        | Connector instance created successfully.                                   |