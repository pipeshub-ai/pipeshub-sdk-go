# ConnectorFilterValue

A saved filter. Sync parses only entries that carry `operator` and `type`; one without `type` is saved but skipped when the connector syncs.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Operator`                                                                                 | `string`                                                                                   | :heavy_check_mark:                                                                         | One of the field's `operators`, for example `in` or `is_after`                             |
| `Value`                                                                                    | `any`                                                                                      | :heavy_check_mark:                                                                         | The filter value; its type follows `type`                                                  |
| `Type`                                                                                     | [components.ConnectorFilterValueType](../../models/components/connectorfiltervaluetype.md) | :heavy_check_mark:                                                                         | The field's `filterType` from the registry schema                                          |