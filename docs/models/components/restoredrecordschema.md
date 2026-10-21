# RestoredRecordSchema


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `RecordID`                                                                | `string`                                                                  | :heavy_check_mark:                                                        | N/A                                                                       |
| `Name`                                                                    | optionalnullable.OptionalNullable[`string`]                               | :heavy_minus_sign:                                                        | The name it has now                                                       |
| `RenamedFrom`                                                             | `*string`                                                                 | :heavy_minus_sign:                                                        | Its name before restore, when another item next to it had taken that name |