# ConnectorStoredConfigSync

Sync strategy plus connector-specific sync settings


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `SelectedStrategy`                                   | `*string`                                            | :heavy_minus_sign:                                   | `MANUAL`, `SCHEDULED` or `WEBHOOK`                   |
| `ScheduledConfig`                                    | map[string]`any`                                     | :heavy_minus_sign:                                   | Schedule used when `selectedStrategy` is `SCHEDULED` |
| `CustomValues`                                       | map[string]`any`                                     | :heavy_minus_sign:                                   | Values of the connector's own sync fields            |
| `AdditionalProperties`                               | map[string]`any`                                     | :heavy_minus_sign:                                   | N/A                                                  |