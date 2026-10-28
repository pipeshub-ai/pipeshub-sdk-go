# ConversationReasoningEffort

Reasoning effort requested for the turn, when one was set.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.ConversationReasoningEffortNone

// Open enum: custom values can be created with a direct type cast
custom := components.ConversationReasoningEffort("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ConversationReasoningEffortNone`   | none                                |
| `ConversationReasoningEffortLow`    | low                                 |
| `ConversationReasoningEffortMedium` | medium                              |
| `ConversationReasoningEffortHigh`   | high                                |
| `ConversationReasoningEffortMax`    | max                                 |