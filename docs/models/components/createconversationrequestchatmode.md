# CreateConversationRequestChatMode

Optional execution mode for non-stream consumers of this shared
request schema. `deep`, `quick`, `verification` and `auto` are
legacy values: still accepted, they run as `internal_search`.
The streaming routes accept only `agent`, `internal_search` and
`web_search`. `agent:<mode>` is refused.
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
| `CreateConversationRequestChatModeDeep`           | deep                                              |
| `CreateConversationRequestChatModeQuick`          | quick                                             |
| `CreateConversationRequestChatModeVerification`   | verification                                      |
| `CreateConversationRequestChatModeAuto`           | auto                                              |