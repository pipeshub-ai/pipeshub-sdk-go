# SemanticSearchGraphRecord

Graph record vertex returned in `records` and as values of `virtual_to_record_map`.
All listed fields are optional in the schema so partial or evolving documents validate; typical Arango documents
usually include `_key`, `_id`, `_rev`, `orgId`, `recordName`, `externalRecordId`, `recordType`, `origin`,
`createdAtTimestamp`, and `connectorId`. Extend this schema when new stable fields appear on Record vertices.



## Fields

| Field                         | Type                          | Required                      | Description                   |
| ----------------------------- | ----------------------------- | ----------------------------- | ----------------------------- |
| `Key`                         | **string*                     | :heavy_minus_sign:            | N/A                           |
| `ID`                          | **string*                     | :heavy_minus_sign:            | N/A                           |
| `Rev`                         | **string*                     | :heavy_minus_sign:            | N/A                           |
| `RecordName`                  | **string*                     | :heavy_minus_sign:            | N/A                           |
| `ExternalRecordID`            | **string*                     | :heavy_minus_sign:            | N/A                           |
| `RecordType`                  | **string*                     | :heavy_minus_sign:            | N/A                           |
| `Origin`                      | **string*                     | :heavy_minus_sign:            | N/A                           |
| `CreatedAtTimestamp`          | **float64*                    | :heavy_minus_sign:            | N/A                           |
| `ConnectorID`                 | **string*                     | :heavy_minus_sign:            | N/A                           |
| `OrgID`                       | **string*                     | :heavy_minus_sign:            | N/A                           |
| `UpdatedAtTimestamp`          | **float64*                    | :heavy_minus_sign:            | N/A                           |
| `ExternalGroupID`             | **string*                     | :heavy_minus_sign:            | N/A                           |
| `ExternalParentID`            | **string*                     | :heavy_minus_sign:            | N/A                           |
| `ExternalRevisionID`          | **string*                     | :heavy_minus_sign:            | N/A                           |
| `ExternalRootGroupID`         | **string*                     | :heavy_minus_sign:            | N/A                           |
| `RecordGroupID`               | **string*                     | :heavy_minus_sign:            | N/A                           |
| `Version`                     | **float64*                    | :heavy_minus_sign:            | N/A                           |
| `ConnectorName`               | **string*                     | :heavy_minus_sign:            | N/A                           |
| `MimeType`                    | **string*                     | :heavy_minus_sign:            | N/A                           |
| `WebURL`                      | **string*                     | :heavy_minus_sign:            | N/A                           |
| `LastSyncTimestamp`           | **float64*                    | :heavy_minus_sign:            | N/A                           |
| `SourceCreatedAtTimestamp`    | **float64*                    | :heavy_minus_sign:            | N/A                           |
| `SourceLastModifiedTimestamp` | **float64*                    | :heavy_minus_sign:            | N/A                           |
| `IsDeleted`                   | **bool*                       | :heavy_minus_sign:            | N/A                           |
| `IsArchived`                  | **bool*                       | :heavy_minus_sign:            | N/A                           |
| `IsVLMOcrProcessed`           | **bool*                       | :heavy_minus_sign:            | N/A                           |
| `DeletedByUserID`             | **string*                     | :heavy_minus_sign:            | N/A                           |
| `ProcessingStartedAt`         | **float64*                    | :heavy_minus_sign:            | N/A                           |
| `ParsingStatus`               | **string*                     | :heavy_minus_sign:            | N/A                           |
| `IndexingStatus`              | **string*                     | :heavy_minus_sign:            | N/A                           |
| `ExtractionStatus`            | **string*                     | :heavy_minus_sign:            | N/A                           |
| `IsLatestVersion`             | **bool*                       | :heavy_minus_sign:            | N/A                           |
| `IsDirty`                     | **bool*                       | :heavy_minus_sign:            | N/A                           |
| `Reason`                      | **string*                     | :heavy_minus_sign:            | N/A                           |
| `LastIndexTimestamp`          | **float64*                    | :heavy_minus_sign:            | N/A                           |
| `LastExtractionTimestamp`     | **float64*                    | :heavy_minus_sign:            | N/A                           |
| `SummaryDocumentID`           | **string*                     | :heavy_minus_sign:            | N/A                           |
| `StorageDocumentID`           | **string*                     | :heavy_minus_sign:            | N/A                           |
| `VirtualRecordID`             | **string*                     | :heavy_minus_sign:            | N/A                           |
| `PreviewRenderable`           | **bool*                       | :heavy_minus_sign:            | N/A                           |
| `IsShared`                    | **bool*                       | :heavy_minus_sign:            | N/A                           |
| `IsDependentNode`             | **bool*                       | :heavy_minus_sign:            | N/A                           |
| `ParentNodeID`                | **string*                     | :heavy_minus_sign:            | N/A                           |
| `HideWeburl`                  | **bool*                       | :heavy_minus_sign:            | N/A                           |
| `IsInternal`                  | **bool*                       | :heavy_minus_sign:            | N/A                           |
| `Md5Checksum`                 | **string*                     | :heavy_minus_sign:            | N/A                           |
| `SizeInBytes`                 | **float64*                    | :heavy_minus_sign:            | N/A                           |
| `Definition`                  | **string*                     | :heavy_minus_sign:            | N/A                           |
| `SourceTables`                | []*string*                    | :heavy_minus_sign:            | N/A                           |
| `RowCount`                    | **float64*                    | :heavy_minus_sign:            | N/A                           |