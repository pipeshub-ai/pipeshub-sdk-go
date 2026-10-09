# AgentRegenerateSSEEventEvent

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentRegenerateSSEEventEventRunStarted

// Open enum: custom values can be created with a direct type cast
custom := components.AgentRegenerateSSEEventEvent("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `AgentRegenerateSSEEventEventRunStarted`              | RUN_STARTED                                           |
| `AgentRegenerateSSEEventEventRunFinished`             | RUN_FINISHED                                          |
| `AgentRegenerateSSEEventEventRunError`                | RUN_ERROR                                             |
| `AgentRegenerateSSEEventEventStepStarted`             | STEP_STARTED                                          |
| `AgentRegenerateSSEEventEventStepFinished`            | STEP_FINISHED                                         |
| `AgentRegenerateSSEEventEventTextMessageStart`        | TEXT_MESSAGE_START                                    |
| `AgentRegenerateSSEEventEventTextMessageContent`      | TEXT_MESSAGE_CONTENT                                  |
| `AgentRegenerateSSEEventEventTextMessageEnd`          | TEXT_MESSAGE_END                                      |
| `AgentRegenerateSSEEventEventReasoningStart`          | REASONING_START                                       |
| `AgentRegenerateSSEEventEventReasoningMessageStart`   | REASONING_MESSAGE_START                               |
| `AgentRegenerateSSEEventEventReasoningMessageContent` | REASONING_MESSAGE_CONTENT                             |
| `AgentRegenerateSSEEventEventReasoningMessageEnd`     | REASONING_MESSAGE_END                                 |
| `AgentRegenerateSSEEventEventReasoningEnd`            | REASONING_END                                         |
| `AgentRegenerateSSEEventEventToolCallStart`           | TOOL_CALL_START                                       |
| `AgentRegenerateSSEEventEventToolCallArgs`            | TOOL_CALL_ARGS                                        |
| `AgentRegenerateSSEEventEventToolCallEnd`             | TOOL_CALL_END                                         |
| `AgentRegenerateSSEEventEventToolCallResult`          | TOOL_CALL_RESULT                                      |
| `AgentRegenerateSSEEventEventStateDelta`              | STATE_DELTA                                           |
| `AgentRegenerateSSEEventEventStateSnapshot`           | STATE_SNAPSHOT                                        |
| `AgentRegenerateSSEEventEventCustom`                  | CUSTOM                                                |
| `AgentRegenerateSSEEventEventHeartbeat`               | HEARTBEAT                                             |