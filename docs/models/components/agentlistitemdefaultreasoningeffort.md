# AgentListItemDefaultReasoningEffort

Agent-level reasoning effort used when a chat request omits its own. Null when unset.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentListItemDefaultReasoningEffortNone

// Open enum: custom values can be created with a direct type cast
custom := components.AgentListItemDefaultReasoningEffort("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AgentListItemDefaultReasoningEffortNone`   | none                                        |
| `AgentListItemDefaultReasoningEffortLow`    | low                                         |
| `AgentListItemDefaultReasoningEffortMedium` | medium                                      |
| `AgentListItemDefaultReasoningEffortHigh`   | high                                        |
| `AgentListItemDefaultReasoningEffortMax`    | max                                         |