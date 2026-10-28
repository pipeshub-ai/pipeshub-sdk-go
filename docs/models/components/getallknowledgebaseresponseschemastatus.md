# GetAllKnowledgeBaseResponseSchemaStatus

The knowledge base's processing state. `DELETING` while it is being deleted (writes are
refused with 409). `IDLE`, `QUEUED`, `SYNCING` or `FULL_SYNCING` otherwise, or null on a
knowledge base that has not been processed yet.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.GetAllKnowledgeBaseResponseSchemaStatusIdle

// Open enum: custom values can be created with a direct type cast
custom := components.GetAllKnowledgeBaseResponseSchemaStatus("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `GetAllKnowledgeBaseResponseSchemaStatusIdle`        | IDLE                                                 |
| `GetAllKnowledgeBaseResponseSchemaStatusQueued`      | QUEUED                                               |
| `GetAllKnowledgeBaseResponseSchemaStatusSyncing`     | SYNCING                                              |
| `GetAllKnowledgeBaseResponseSchemaStatusFullSyncing` | FULL_SYNCING                                         |
| `GetAllKnowledgeBaseResponseSchemaStatusDeleting`    | DELETING                                             |