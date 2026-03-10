# UpdateConnectorFiltersSyncRequest

Request to update filters and sync config (connector must be authenticated and inactive)


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `Sync`                                                                                  | [*components.ConnectorSyncConfig](../../models/components/connectorsyncconfig.md)       | :heavy_minus_sign:                                                                      | Synchronization configuration for a connector instance                                  |
| `Filters`                                                                               | [*components.ConnectorFiltersConfig](../../models/components/connectorfiltersconfig.md) | :heavy_minus_sign:                                                                      | Filter configuration to control what data is synced (sync filters and indexing filters) |
| `BaseURL`                                                                               | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     |