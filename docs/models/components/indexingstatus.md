# IndexingStatus

Current indexing/processing status:
- NOT_STARTED: Awaiting indexing
- QUEUED: In indexing queue
- IN_PROGRESS: Currently being indexed
- COMPLETED: Successfully indexed and searchable
- FAILED: Indexing failed (check error details)
- PAUSED: Indexing paused by user
- FILE_TYPE_NOT_SUPPORTED: Unsupported file format
- AUTO_INDEX_OFF: Auto-indexing disabled for this record
- EMPTY: File has no extractable content
- ENABLE_MULTIMODAL_MODELS: Requires multimodal AI models



## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `IndexingStatusNotStarted`             | NOT_STARTED                            |
| `IndexingStatusPaused`                 | PAUSED                                 |
| `IndexingStatusInProgress`             | IN_PROGRESS                            |
| `IndexingStatusCompleted`              | COMPLETED                              |
| `IndexingStatusFailed`                 | FAILED                                 |
| `IndexingStatusFileTypeNotSupported`   | FILE_TYPE_NOT_SUPPORTED                |
| `IndexingStatusAutoIndexOff`           | AUTO_INDEX_OFF                         |
| `IndexingStatusEmpty`                  | EMPTY                                  |
| `IndexingStatusEnableMultimodalModels` | ENABLE_MULTIMODAL_MODELS               |
| `IndexingStatusQueued`                 | QUEUED                                 |