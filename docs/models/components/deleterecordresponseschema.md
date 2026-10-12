# DeleteRecordResponseSchema

Response returned by DELETE /knowledgeBase/record/{recordId}.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Success`                                   | `bool`                                      | :heavy_check_mark:                          | N/A                                         |
| `Message`                                   | `string`                                    | :heavy_check_mark:                          | N/A                                         |
| `RecordID`                                  | `string`                                    | :heavy_check_mark:                          | N/A                                         |
| `Connector`                                 | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | N/A                                         |
| `Timestamp`                                 | optionalnullable.OptionalNullable[`int64`]  | :heavy_minus_sign:                          | N/A                                         |