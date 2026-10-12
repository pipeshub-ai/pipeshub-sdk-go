# KnowledgeBaseCreateResponseUserRole

User's role in this knowledge base

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.KnowledgeBaseCreateResponseUserRoleOwner

// Open enum: custom values can be created with a direct type cast
custom := components.KnowledgeBaseCreateResponseUserRole("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `KnowledgeBaseCreateResponseUserRoleOwner`  | OWNER                                       |
| `KnowledgeBaseCreateResponseUserRoleWriter` | WRITER                                      |
| `KnowledgeBaseCreateResponseUserRoleReader` | READER                                      |