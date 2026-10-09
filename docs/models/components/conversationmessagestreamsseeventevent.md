# ConversationMessageStreamSSEEventEvent

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ConversationMessageStreamSSEEventEventRunStarted

// Open enum: custom values can be created with a direct type cast
custom := components.ConversationMessageStreamSSEEventEvent("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ConversationMessageStreamSSEEventEventRunStarted`              | RUN_STARTED                                                     |
| `ConversationMessageStreamSSEEventEventRunFinished`             | RUN_FINISHED                                                    |
| `ConversationMessageStreamSSEEventEventRunError`                | RUN_ERROR                                                       |
| `ConversationMessageStreamSSEEventEventStepStarted`             | STEP_STARTED                                                    |
| `ConversationMessageStreamSSEEventEventStepFinished`            | STEP_FINISHED                                                   |
| `ConversationMessageStreamSSEEventEventTextMessageStart`        | TEXT_MESSAGE_START                                              |
| `ConversationMessageStreamSSEEventEventTextMessageContent`      | TEXT_MESSAGE_CONTENT                                            |
| `ConversationMessageStreamSSEEventEventTextMessageEnd`          | TEXT_MESSAGE_END                                                |
| `ConversationMessageStreamSSEEventEventReasoningStart`          | REASONING_START                                                 |
| `ConversationMessageStreamSSEEventEventReasoningMessageStart`   | REASONING_MESSAGE_START                                         |
| `ConversationMessageStreamSSEEventEventReasoningMessageContent` | REASONING_MESSAGE_CONTENT                                       |
| `ConversationMessageStreamSSEEventEventReasoningMessageEnd`     | REASONING_MESSAGE_END                                           |
| `ConversationMessageStreamSSEEventEventReasoningEnd`            | REASONING_END                                                   |
| `ConversationMessageStreamSSEEventEventToolCallStart`           | TOOL_CALL_START                                                 |
| `ConversationMessageStreamSSEEventEventToolCallArgs`            | TOOL_CALL_ARGS                                                  |
| `ConversationMessageStreamSSEEventEventToolCallEnd`             | TOOL_CALL_END                                                   |
| `ConversationMessageStreamSSEEventEventToolCallResult`          | TOOL_CALL_RESULT                                                |
| `ConversationMessageStreamSSEEventEventStateDelta`              | STATE_DELTA                                                     |
| `ConversationMessageStreamSSEEventEventStateSnapshot`           | STATE_SNAPSHOT                                                  |
| `ConversationMessageStreamSSEEventEventCustom`                  | CUSTOM                                                          |
| `ConversationMessageStreamSSEEventEventHeartbeat`               | HEARTBEAT                                                       |