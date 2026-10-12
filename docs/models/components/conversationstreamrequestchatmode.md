# ConversationStreamRequestChatMode

Optional execution mode for non-stream consumers of this shared
request schema.
`agent` uses the universal agent loop, while `internal_search`
and `web_search` use their corresponding assistant search paths.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ConversationStreamRequestChatModeAgent
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `ConversationStreamRequestChatModeAgent`          | agent                                             |
| `ConversationStreamRequestChatModeInternalSearch` | internal_search                                   |
| `ConversationStreamRequestChatModeWebSearch`      | web_search                                        |