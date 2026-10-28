# ConversationListItemReasoningEffort

Reasoning effort requested when the conversation was created, when one was set.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ConversationListItemReasoningEffortNone

// Open enum: custom values can be created with a direct type cast
custom := components.ConversationListItemReasoningEffort("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `ConversationListItemReasoningEffortNone`   | none                                        |
| `ConversationListItemReasoningEffortLow`    | low                                         |
| `ConversationListItemReasoningEffortMedium` | medium                                      |
| `ConversationListItemReasoningEffortHigh`   | high                                        |
| `ConversationListItemReasoningEffortMax`    | max                                         |