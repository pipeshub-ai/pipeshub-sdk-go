# ReIndexRecordGroupResponseSchema

Response returned by POST /knowledgeBase/reindex/record-group/{recordGroupId}.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Success`                                                | `bool`                                                   | :heavy_check_mark:                                       | N/A                                                      |
| `Message`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `RecordGroupID`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `Depth`                                                  | `int64`                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `ConnectorID`                                            | `string`                                                 | :heavy_check_mark:                                       | Id of the connector instance the record group belongs to |
| `Connector`                                              | `*string`                                                | :heavy_check_mark:                                       | N/A                                                      |
| `EventPublished`                                         | `bool`                                                   | :heavy_check_mark:                                       | N/A                                                      |