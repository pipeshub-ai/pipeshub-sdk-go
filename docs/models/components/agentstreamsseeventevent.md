# AgentStreamSSEEventEvent

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentStreamSSEEventEventRunStarted

// Open enum: custom values can be created with a direct type cast
custom := components.AgentStreamSSEEventEvent("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `AgentStreamSSEEventEventRunStarted`              | RUN_STARTED                                       |
| `AgentStreamSSEEventEventRunFinished`             | RUN_FINISHED                                      |
| `AgentStreamSSEEventEventRunError`                | RUN_ERROR                                         |
| `AgentStreamSSEEventEventStepStarted`             | STEP_STARTED                                      |
| `AgentStreamSSEEventEventStepFinished`            | STEP_FINISHED                                     |
| `AgentStreamSSEEventEventTextMessageStart`        | TEXT_MESSAGE_START                                |
| `AgentStreamSSEEventEventTextMessageContent`      | TEXT_MESSAGE_CONTENT                              |
| `AgentStreamSSEEventEventTextMessageEnd`          | TEXT_MESSAGE_END                                  |
| `AgentStreamSSEEventEventReasoningStart`          | REASONING_START                                   |
| `AgentStreamSSEEventEventReasoningMessageStart`   | REASONING_MESSAGE_START                           |
| `AgentStreamSSEEventEventReasoningMessageContent` | REASONING_MESSAGE_CONTENT                         |
| `AgentStreamSSEEventEventReasoningMessageEnd`     | REASONING_MESSAGE_END                             |
| `AgentStreamSSEEventEventReasoningEnd`            | REASONING_END                                     |
| `AgentStreamSSEEventEventToolCallStart`           | TOOL_CALL_START                                   |
| `AgentStreamSSEEventEventToolCallArgs`            | TOOL_CALL_ARGS                                    |
| `AgentStreamSSEEventEventToolCallEnd`             | TOOL_CALL_END                                     |
| `AgentStreamSSEEventEventToolCallResult`          | TOOL_CALL_RESULT                                  |
| `AgentStreamSSEEventEventStateDelta`              | STATE_DELTA                                       |
| `AgentStreamSSEEventEventStateSnapshot`           | STATE_SNAPSHOT                                    |
| `AgentStreamSSEEventEventCustom`                  | CUSTOM                                            |
| `AgentStreamSSEEventEventHeartbeat`               | HEARTBEAT                                         |