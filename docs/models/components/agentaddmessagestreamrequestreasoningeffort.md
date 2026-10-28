# AgentAddMessageStreamRequestReasoningEffort

Reasoning effort for this turn, for models that support it. Stored on
the conversation's model info and forwarded to the AI backend.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentAddMessageStreamRequestReasoningEffortNone
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `AgentAddMessageStreamRequestReasoningEffortNone`   | none                                                |
| `AgentAddMessageStreamRequestReasoningEffortLow`    | low                                                 |
| `AgentAddMessageStreamRequestReasoningEffortMedium` | medium                                              |
| `AgentAddMessageStreamRequestReasoningEffortHigh`   | high                                                |
| `AgentAddMessageStreamRequestReasoningEffortMax`    | max                                                 |