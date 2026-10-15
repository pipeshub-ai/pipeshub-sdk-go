# DeleteRecordResponseSchema

Response returned by DELETE /knowledgeBase/record/{recordId}.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Success`                                                        | `bool`                                                           | :heavy_check_mark:                                               | N/A                                                              |
| `Message`                                                        | `string`                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `RecordID`                                                       | `string`                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `Connector`                                                      | optionalnullable.OptionalNullable[`string`]                      | :heavy_minus_sign:                                               | N/A                                                              |
| `Timestamp`                                                      | optionalnullable.OptionalNullable[`int64`]                       | :heavy_minus_sign:                                               | N/A                                                              |
| `SoftDeleted`                                                    | `*bool`                                                          | :heavy_minus_sign:                                               | True when the record went to the trash instead of being removed  |
| `BatchID`                                                        | optionalnullable.OptionalNullable[`string`]                      | :heavy_minus_sign:                                               | The delete batch; restoring any record in it restores all of it  |
| `VectorCleanupPending`                                           | `*bool`                                                          | :heavy_minus_sign:                                               | The record is deleted, but removing its embeddings is still owed |
| `VectorCleanupFailedRecordIds`                                   | []`string`                                                       | :heavy_minus_sign:                                               | N/A                                                              |
| `VectorCleanupFailedVirtualRecordIds`                            | []`string`                                                       | :heavy_minus_sign:                                               | N/A                                                              |