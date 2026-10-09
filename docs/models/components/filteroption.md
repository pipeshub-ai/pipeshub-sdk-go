# FilterOption

A single filter option for knowledge hub filters.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ID`                                                                | `string`                                                            | :heavy_check_mark:                                                  | Filter ID value to send in requests.                                |
| `Label`                                                             | `string`                                                            | :heavy_check_mark:                                                  | Display label for the filter.                                       |
| `Type`                                                              | optionalnullable.OptionalNullable[`string`]                         | :heavy_minus_sign:                                                  | Additional type information (currently unused, may be null).        |
| `ConnectorType`                                                     | optionalnullable.OptionalNullable[`string`]                         | :heavy_minus_sign:                                                  | Connector type/name. Set only for entries in the `connectors` list. |