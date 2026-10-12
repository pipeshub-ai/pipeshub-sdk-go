# GetKnowledgeBaseByIDUserRole

User's role in this knowledge base

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.GetKnowledgeBaseByIDUserRoleOwner

// Open enum: custom values can be created with a direct type cast
custom := components.GetKnowledgeBaseByIDUserRole("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `GetKnowledgeBaseByIDUserRoleOwner`  | OWNER                                |
| `GetKnowledgeBaseByIDUserRoleWriter` | WRITER                               |
| `GetKnowledgeBaseByIDUserRoleReader` | READER                               |