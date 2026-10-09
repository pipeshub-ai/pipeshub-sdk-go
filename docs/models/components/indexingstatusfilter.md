# IndexingStatusFilter

Indexing status used to filter which records are included in a scoped
reindex (record or record-group). Omit `statusFilters` to reindex all
descendants regardless of status.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.IndexingStatusFilterNotStarted
```


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