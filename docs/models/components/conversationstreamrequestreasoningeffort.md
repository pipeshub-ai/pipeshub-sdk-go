# ConversationStreamRequestReasoningEffort

How hard a reasoning model should think for this turn. Saved in
the conversation's `modelInfo`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.ConversationStreamRequestReasoningEffortNone
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `ConversationStreamRequestReasoningEffortNone`   | none                                             |
| `ConversationStreamRequestReasoningEffortLow`    | low                                              |
| `ConversationStreamRequestReasoningEffortMedium` | medium                                           |
| `ConversationStreamRequestReasoningEffortHigh`   | high                                             |
| `ConversationStreamRequestReasoningEffortMax`    | max                                              |