# CreateConversationRequestChatMode

Optional execution mode for non-stream consumers of this shared
request schema.
`agent` uses the universal agent loop, while `internal_search`
and `web_search` use their corresponding assistant search paths.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.CreateConversationRequestChatModeAgent
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `CreateConversationRequestChatModeAgent`          | agent                                             |
| `CreateConversationRequestChatModeInternalSearch` | internal_search                                   |
| `CreateConversationRequestChatModeWebSearch`      | web_search                                        |