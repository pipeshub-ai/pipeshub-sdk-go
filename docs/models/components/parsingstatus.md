# ParsingStatus

Parse-phase status (ahead of indexing/extraction):
- NOT_STARTED: Awaiting parsing
- QUEUED: In parsing queue
- IN_PROGRESS: Currently being parsed
- COMPLETED: Successfully parsed
- FAILED: Parsing failed
- FILE_TYPE_NOT_SUPPORTED: Unsupported file format
- AUTO_INDEX_OFF: Auto-indexing disabled for this record
- EMPTY: File has no extractable content



## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ParsingStatusNotStarted`           | NOT_STARTED                         |
| `ParsingStatusInProgress`           | IN_PROGRESS                         |
| `ParsingStatusFailed`               | FAILED                              |
| `ParsingStatusCompleted`            | COMPLETED                           |
| `ParsingStatusFileTypeNotSupported` | FILE_TYPE_NOT_SUPPORTED             |
| `ParsingStatusAutoIndexOff`         | AUTO_INDEX_OFF                      |
| `ParsingStatusEmpty`                | EMPTY                               |
| `ParsingStatusQueued`               | QUEUED                              |