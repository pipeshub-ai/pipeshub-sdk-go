# AgentMessageStreamSSEEvent

Server-Sent Event envelope for `POST /agents/{agentKey}/conversations/{conversationId}/messages/stream`.
AG-UI is the sole wire protocol.

`event` carries the AG-UI type name and `data` is the decoded JSON
object, which includes a `"type"` field matching `event`, plus
type-specific fields. The public route requires `chatMode: quick`.
Forwarded lifecycle events may carry `runId`, `threadId`, and
`parentRunId`. Stable gateway-generated top-level outcomes:

- `CUSTOM` (`name: "conversation_created"`) — fired once after the
  SSE stream opens.
- `RUN_FINISHED` — fired once after the upstream AI's result is
  parsed, citations are saved, and the updated conversation is
  persisted. The gateway emits `{ type, result }`; `result` carries
  `{ conversation, recordsUsed, meta }`.
- `RUN_ERROR` — fired for runtime failures after the stream has
  already started, including conversation lookup failures, upstream
  AI startup failures, save failures, and stream transport errors.

Gateway-generated root terminal events do not contain `runId`,
`threadId`, or `parentRunId`.
Clients should ignore unknown event names rather than treating them
as errors.



## Fields

| Field                                                                                                                      | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `Event`                                                                                                                    | [*components.AgentMessageStreamSSEEventEvent](../../models/components/agentmessagestreamsseeventevent.md)                  | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `Data`                                                                                                                     | map[string]`any`                                                                                                           | :heavy_minus_sign:                                                                                                         | Decoded JSON payload. It includes a `"type"`<br/>field matching `event`, plus type-specific fields. Shape depends<br/>on `event`.<br/> |