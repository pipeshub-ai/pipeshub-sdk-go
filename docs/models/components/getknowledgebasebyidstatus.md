# GetKnowledgeBaseByIDStatus

`DELETING` while the knowledge base is being deleted: reads work, every write is refused with 409
(`error.code` `HTTP_KB_BEING_DELETED`), and deleting it again finishes the job. Otherwise its
processing state (`IDLE`, `QUEUED`, `SYNCING`, `FULL_SYNCING`), or null before it is first processed.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.GetKnowledgeBaseByIDStatusIdle

// Open enum: custom values can be created with a direct type cast
custom := components.GetKnowledgeBaseByIDStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `GetKnowledgeBaseByIDStatusIdle`        | IDLE                                    |
| `GetKnowledgeBaseByIDStatusQueued`      | QUEUED                                  |
| `GetKnowledgeBaseByIDStatusSyncing`     | SYNCING                                 |
| `GetKnowledgeBaseByIDStatusFullSyncing` | FULL_SYNCING                            |
| `GetKnowledgeBaseByIDStatusDeleting`    | DELETING                                |