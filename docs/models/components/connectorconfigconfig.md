# ConnectorConfigConfig

Configuration sections


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `Auth`                                                                                  | [*components.Auth](../../models/components/auth.md)                                     | :heavy_minus_sign:                                                                      | Authentication configuration (sensitive data redacted)                                  |
| `Sync`                                                                                  | [*components.ConnectorConfigSync](../../models/components/connectorconfigsync.md)       | :heavy_minus_sign:                                                                      | Sync configuration (schedule, options)                                                  |
| `Filters`                                                                               | [*components.ConnectorConfigFilters](../../models/components/connectorconfigfilters.md) | :heavy_minus_sign:                                                                      | Filter selections for data scope                                                        |