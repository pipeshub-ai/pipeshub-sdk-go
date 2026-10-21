# ConnectorStoredConfig

A connector instance's saved settings. Secret auth fields come back masked; saving the mask back keeps the stored value. Provider OAuth tokens are never returned.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `Auth`                                                                                        | [*components.ConnectorStoredConfigAuth](../../models/components/connectorstoredconfigauth.md) | :heavy_minus_sign:                                                                            | Auth fields for the instance's auth type, keyed by field name                                 |
| `Sync`                                                                                        | map[string]`any`                                                                              | :heavy_minus_sign:                                                                            | Sync strategy plus connector-specific sync settings                                           |
| `Filters`                                                                                     | map[string]`any`                                                                              | :heavy_minus_sign:                                                                            | N/A                                                                                           |