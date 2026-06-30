# ReindexRecordGroupRequestBody

Optional body for record-group (folder/KB container) reindex.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Depth`                                                                              | **int64*                                                                             | :heavy_minus_sign:                                                                   | Depth of records under the record group to include.                                  |
| `Force`                                                                              | **bool*                                                                              | :heavy_minus_sign:                                                                   | Force reindex for all matched records in the group.                                  |
| `StatusFilters`                                                                      | [][components.IndexingStatusFilter](../../models/components/indexingstatusfilter.md) | :heavy_minus_sign:                                                                   | When set, only records matching these indexing statuses are reindexed.<br/>          |