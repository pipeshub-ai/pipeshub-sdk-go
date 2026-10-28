# RestoreRecordsRequestBody

Request body for POST /knowledgeBase/records/restore. Fields other than `recordIds` are dropped by the validator, not refused.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `RecordIds`                                                                         | []`string`                                                                          | :heavy_check_mark:                                                                  | Record ids to restore. An id listed more than once is restored, and answered, once. |