# ConversationModelInfoReasoningEffort

Reasoning effort requested for the turn, when one was set.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.ConversationModelInfoReasoningEffortNone

// Open enum: custom values can be created with a direct type cast
custom := components.ConversationModelInfoReasoningEffort("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `ConversationModelInfoReasoningEffortNone`   | none                                         |
| `ConversationModelInfoReasoningEffortLow`    | low                                          |
| `ConversationModelInfoReasoningEffortMedium` | medium                                       |
| `ConversationModelInfoReasoningEffortHigh`   | high                                         |
| `ConversationModelInfoReasoningEffortMax`    | max                                          |