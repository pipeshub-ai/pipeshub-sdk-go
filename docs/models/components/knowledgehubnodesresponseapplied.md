# KnowledgeHubNodesResponseApplied

Echo of applied filters; unused slots are JSON `null`.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Q`                                                          | *string*                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `NodeTypes`                                                  | []*string*                                                   | :heavy_check_mark:                                           | N/A                                                          |
| `RecordTypes`                                                | []*string*                                                   | :heavy_check_mark:                                           | N/A                                                          |
| `Origins`                                                    | []*string*                                                   | :heavy_check_mark:                                           | N/A                                                          |
| `ConnectorIds`                                               | []*string*                                                   | :heavy_check_mark:                                           | N/A                                                          |
| `IndexingStatus`                                             | []*string*                                                   | :heavy_check_mark:                                           | N/A                                                          |
| `CreatedAt`                                                  | [components.CreatedAt](../../models/components/createdat.md) | :heavy_check_mark:                                           | N/A                                                          |
| `UpdatedAt`                                                  | [components.UpdatedAt](../../models/components/updatedat.md) | :heavy_check_mark:                                           | N/A                                                          |
| `Size`                                                       | [components.Size](../../models/components/size.md)           | :heavy_check_mark:                                           | N/A                                                          |
| `SortBy`                                                     | *string*                                                     | :heavy_check_mark:                                           | Effective sort field after server normalisation.             |
| `SortOrder`                                                  | *string*                                                     | :heavy_check_mark:                                           | Effective sort order after server normalisation.             |