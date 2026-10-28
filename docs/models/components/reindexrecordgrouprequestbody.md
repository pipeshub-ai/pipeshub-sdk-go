# ReindexRecordGroupRequestBody

Optional body for record-group reindex. Only `depth` and `statusFilters` are read;
any other field is dropped by the validator before the handler runs (there is no
`force` option: the handler always sends `force: false`).



## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `Depth`                                                                                                    | `*int64`                                                                                                   | :heavy_minus_sign:                                                                                         | Depth of records under the record group to include. `-1` (everything) is applied,<br/>and answered, as `100`.<br/> |
| `StatusFilters`                                                                                            | []`string`                                                                                                 | :heavy_minus_sign:                                                                                         | When set, only records matching these indexing statuses are reindexed.<br/>                                |