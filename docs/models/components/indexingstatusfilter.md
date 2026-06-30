# IndexingStatusFilter

Indexing status used to filter which records are included in a scoped
reindex (record or record-group). Omit `statusFilters` to reindex all
descendants regardless of status.



## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `IndexingStatusFilterNotStarted`           | NOT_STARTED                                |
| `IndexingStatusFilterQueued`               | QUEUED                                     |
| `IndexingStatusFilterInProgress`           | IN_PROGRESS                                |
| `IndexingStatusFilterCompleted`            | COMPLETED                                  |
| `IndexingStatusFilterFailed`               | FAILED                                     |
| `IndexingStatusFilterFileTypeNotSupported` | FILE_TYPE_NOT_SUPPORTED                    |
| `IndexingStatusFilterAutoIndexOff`         | AUTO_INDEX_OFF                             |
| `IndexingStatusFilterEmpty`                | EMPTY                                      |