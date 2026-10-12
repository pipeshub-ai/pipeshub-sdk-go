# SemanticSearchGraphRecord

Graph record vertex returned in `records` and as values of `virtual_to_record_map`.
All listed fields are optional in the schema so partial or evolving documents validate; typical Arango documents
usually include `_key`, `_id`, `_rev`, `orgId`, `recordName`, `externalRecordId`, `recordType`, `origin`,
`createdAtTimestamp`, and `connectorId`. Extend this schema when new stable fields appear on Record vertices.



## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Key`                                         | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `ID`                                          | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `Rev`                                         | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `RecordName`                                  | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `ExternalRecordID`                            | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `RecordType`                                  | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `Origin`                                      | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `CreatedAtTimestamp`                          | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `ConnectorID`                                 | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `OrgID`                                       | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `UpdatedAtTimestamp`                          | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `ExternalGroupID`                             | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `ExternalParentID`                            | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `ExternalRevisionID`                          | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `ExternalRootGroupID`                         | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `RecordGroupID`                               | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `Version`                                     | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `ConnectorName`                               | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `MimeType`                                    | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `WebURL`                                      | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `LastSyncTimestamp`                           | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `SourceCreatedAtTimestamp`                    | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `SourceLastModifiedTimestamp`                 | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `IsDeleted`                                   | optionalnullable.OptionalNullable[`bool`]     | :heavy_minus_sign:                            | N/A                                           |
| `IsArchived`                                  | optionalnullable.OptionalNullable[`bool`]     | :heavy_minus_sign:                            | N/A                                           |
| `IsVLMOcrProcessed`                           | optionalnullable.OptionalNullable[`bool`]     | :heavy_minus_sign:                            | N/A                                           |
| `DeletedByUserID`                             | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `ProcessingStartedAt`                         | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `QueuedAtTimestamp`                           | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `ParsingStatus`                               | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `IndexingStatus`                              | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `ExtractionStatus`                            | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `IsLatestVersion`                             | optionalnullable.OptionalNullable[`bool`]     | :heavy_minus_sign:                            | N/A                                           |
| `IsDirty`                                     | optionalnullable.OptionalNullable[`bool`]     | :heavy_minus_sign:                            | N/A                                           |
| `Reason`                                      | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `LastIndexTimestamp`                          | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `LastExtractionTimestamp`                     | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `SummaryDocumentID`                           | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `StorageDocumentID`                           | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `VirtualRecordID`                             | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `PreviewRenderable`                           | optionalnullable.OptionalNullable[`bool`]     | :heavy_minus_sign:                            | N/A                                           |
| `IsShared`                                    | optionalnullable.OptionalNullable[`bool`]     | :heavy_minus_sign:                            | N/A                                           |
| `IsDependentNode`                             | optionalnullable.OptionalNullable[`bool`]     | :heavy_minus_sign:                            | N/A                                           |
| `ParentNodeID`                                | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `HideWeburl`                                  | optionalnullable.OptionalNullable[`bool`]     | :heavy_minus_sign:                            | N/A                                           |
| `IsInternal`                                  | optionalnullable.OptionalNullable[`bool`]     | :heavy_minus_sign:                            | N/A                                           |
| `Md5Checksum`                                 | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `SizeInBytes`                                 | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |
| `Definition`                                  | optionalnullable.OptionalNullable[`string`]   | :heavy_minus_sign:                            | N/A                                           |
| `SourceTables`                                | optionalnullable.OptionalNullable[[]`string`] | :heavy_minus_sign:                            | N/A                                           |
| `RowCount`                                    | optionalnullable.OptionalNullable[`float64`]  | :heavy_minus_sign:                            | N/A                                           |