# DefaultReasoningEffort

The model's own reasoning effort. It applies only when the chat request sets no `reasoningEffort` and the agent has no `defaultReasoningEffort`; a level the user picks always wins. Read from `configuration.defaultReasoningEffort`, falling back to a top-level `defaultReasoningEffort`. Present on every model expanded from the entry, and left out when no default is stored. `none` is a legacy value that runs as `low`.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.DefaultReasoningEffortNone

// Open enum: custom values can be created with a direct type cast
custom := operations.DefaultReasoningEffort("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `DefaultReasoningEffortNone`   | none                           |
| `DefaultReasoningEffortLow`    | low                            |
| `DefaultReasoningEffortMedium` | medium                         |
| `DefaultReasoningEffortHigh`   | high                           |
| `DefaultReasoningEffortMax`    | max                            |