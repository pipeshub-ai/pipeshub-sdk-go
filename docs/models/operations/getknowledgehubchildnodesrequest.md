# GetKnowledgeHubChildNodesRequest


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `ParentType`                                     | *string*                                         | :heavy_check_mark:                               | Type of parent node (KB, FOLDER, CONNECTOR, APP) |
| `ParentID`                                       | *string*                                         | :heavy_check_mark:                               | ID of parent node                                |
| `Page`                                           | **int64*                                         | :heavy_minus_sign:                               | N/A                                              |
| `Limit`                                          | **int64*                                         | :heavy_minus_sign:                               | N/A                                              |
| `Q`                                              | **string*                                        | :heavy_minus_sign:                               | Search query                                     |