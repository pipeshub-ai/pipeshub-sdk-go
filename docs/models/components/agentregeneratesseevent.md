# AgentRegenerateSSEEvent

SSE event envelope for `POST /agents/{agentKey}/conversations/{conversationId}/message/{messageId}/regenerate`.
AG-UI is the sole wire protocol.

`event` carries the AG-UI type name and `data` is a JSON object that
includes a `"type"` field matching `event`, plus type-specific
fields. Stable gateway-generated top-level outcomes:

- `RUN_FINISHED` returns `{ type, result }` where
  `result` is `{ conversation, recordsUsed, meta }` — the updated
  conversation plus request metadata after the regenerated response
  is persisted.
- `RUN_ERROR` returns `{ type, message, code? }`. Conversation
  lookup failures, unauthorized conversation access, and regenerate
  rule failures such as "not the last message" are reported here.

Other events are forwarded from the agent backend and should be
treated as informational updates. Those forwarded lifecycle and
child-run events may contain `runId`, `threadId`, and `parentRunId`;
the gateway-generated root terminal event does not.



## Fields

| Field                                                                                                                                          | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `Event`                                                                                                                                        | [*components.AgentRegenerateSSEEventEvent](../../models/components/agentregeneratesseeventevent.md)                                            | :heavy_minus_sign:                                                                                                                             | N/A                                                                                                                                            |
| `Data`                                                                                                                                         | **string*                                                                                                                                      | :heavy_minus_sign:                                                                                                                             | JSON-encoded event payload. The decoded JSON includes a `"type"`<br/>field matching `event`, plus type-specific fields. Shape depends<br/>on `event`.<br/> |