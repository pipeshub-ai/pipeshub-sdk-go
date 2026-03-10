# CreateConnectorRequestConfig

Initial configuration (can also be set after creation)


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `Auth`                                                                                  | [*components.ConnectorAuthConfig](../../models/components/connectorauthconfig.md)       | :heavy_minus_sign:                                                                      | Authentication configuration for a connector instance                                   |
| `Sync`                                                                                  | [*components.ConnectorSyncConfig](../../models/components/connectorsyncconfig.md)       | :heavy_minus_sign:                                                                      | Synchronization configuration for a connector instance                                  |
| `Filters`                                                                               | [*components.ConnectorFiltersConfig](../../models/components/connectorfiltersconfig.md) | :heavy_minus_sign:                                                                      | Filter configuration to control what data is synced (sync filters and indexing filters) |