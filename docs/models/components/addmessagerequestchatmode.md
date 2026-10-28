# AddMessageRequestChatMode

Optional execution mode for non-stream consumers of this shared
request schema. `deep`, `quick`, `verification` and `auto` are
legacy values: still accepted, they run as `internal_search`.
The streaming routes accept only `agent`, `internal_search` and
`web_search`. `agent:<mode>` is refused.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.AddMessageRequestChatModeAgent
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `AddMessageRequestChatModeAgent`          | agent                                     |
| `AddMessageRequestChatModeInternalSearch` | internal_search                           |
| `AddMessageRequestChatModeWebSearch`      | web_search                                |
| `AddMessageRequestChatModeDeep`           | deep                                      |
| `AddMessageRequestChatModeQuick`          | quick                                     |
| `AddMessageRequestChatModeVerification`   | verification                              |
| `AddMessageRequestChatModeAuto`           | auto                                      |