# ModelDefaultReasoningEffort

The model's own reasoning effort, read from its stored configuration. It applies only when the chat request sets no `reasoningEffort` and the agent has no `defaultReasoningEffort`. Omitted when nothing valid is stored; `none` is a legacy value that runs as `low`.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ModelDefaultReasoningEffortNone

// Open enum: custom values can be created with a direct type cast
custom := components.ModelDefaultReasoningEffort("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ModelDefaultReasoningEffortNone`   | none                                |
| `ModelDefaultReasoningEffortLow`    | low                                 |
| `ModelDefaultReasoningEffortMedium` | medium                              |
| `ModelDefaultReasoningEffortHigh`   | high                                |
| `ModelDefaultReasoningEffortMax`    | max                                 |