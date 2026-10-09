# Relationship

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.RelationshipOwner

// Open enum: custom values can be created with a direct type cast
custom := components.Relationship("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `RelationshipOwner`  | OWNER                |
| `RelationshipWriter` | WRITER               |
| `RelationshipReader` | READER               |