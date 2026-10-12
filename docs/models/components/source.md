# Source

Origin of the feedback. Always present in responses (server applies the default `user`).

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.SourceUser

// Open enum: custom values can be created with a direct type cast
custom := components.Source("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `SourceUser`   | user           |
| `SourceSystem` | system         |
| `SourceAdmin`  | admin          |
| `SourceAuto`   | auto           |