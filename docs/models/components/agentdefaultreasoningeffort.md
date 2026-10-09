# AgentDefaultReasoningEffort

Agent-level reasoning effort used when a chat request omits its own. Null when unset.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentDefaultReasoningEffortNone

// Open enum: custom values can be created with a direct type cast
custom := components.AgentDefaultReasoningEffort("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `AgentDefaultReasoningEffortNone`   | none                                |
| `AgentDefaultReasoningEffortLow`    | low                                 |
| `AgentDefaultReasoningEffortMedium` | medium                              |
| `AgentDefaultReasoningEffortHigh`   | high                                |
| `AgentDefaultReasoningEffortMax`    | max                                 |