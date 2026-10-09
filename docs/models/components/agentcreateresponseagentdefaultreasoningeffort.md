# AgentCreateResponseAgentDefaultReasoningEffort

Agent-level reasoning effort used when a chat request omits its own. Null when unset.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentCreateResponseAgentDefaultReasoningEffortNone

// Open enum: custom values can be created with a direct type cast
custom := components.AgentCreateResponseAgentDefaultReasoningEffort("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `AgentCreateResponseAgentDefaultReasoningEffortNone`   | none                                                   |
| `AgentCreateResponseAgentDefaultReasoningEffortLow`    | low                                                    |
| `AgentCreateResponseAgentDefaultReasoningEffortMedium` | medium                                                 |
| `AgentCreateResponseAgentDefaultReasoningEffortHigh`   | high                                                   |
| `AgentCreateResponseAgentDefaultReasoningEffortMax`    | max                                                    |