# AgentAddMessageRequestReasoningEffort

Reasoning effort for this turn, for models that support it. Stored on
the conversation's model info and forwarded to the AI backend.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentAddMessageRequestReasoningEffortNone
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `AgentAddMessageRequestReasoningEffortNone`   | none                                          |
| `AgentAddMessageRequestReasoningEffortLow`    | low                                           |
| `AgentAddMessageRequestReasoningEffortMedium` | medium                                        |
| `AgentAddMessageRequestReasoningEffortHigh`   | high                                          |
| `AgentAddMessageRequestReasoningEffortMax`    | max                                           |