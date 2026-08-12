# AgentStreamSSEEvent

SSE event envelope for `POST /agents/{agentKey}/conversations/stream`.
AG-UI is the sole wire protocol.

`event` carries the AG-UI type name and `data` is a JSON-encoded
object that includes a `"type"` field matching `event`, plus
type-specific fields. The public route requires `chatMode: quick`.
Forwarded lifecycle events may carry `runId`, `threadId`, and
`parentRunId`; gateway-generated root terminal events do not.

Stable gateway-generated top-level outcomes are `RUN_FINISHED` as
`{ type, result }` and `RUN_ERROR` as `{ type, message, code? }`.
Clients should ignore unknown event names rather than treating them
as errors.



## Fields

| Field                                                                                                                                          | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `Event`                                                                                                                                        | [*components.AgentStreamSSEEventEvent](../../models/components/agentstreamsseeventevent.md)                                                    | :heavy_minus_sign:                                                                                                                             | N/A                                                                                                                                            |
| `Data`                                                                                                                                         | **string*                                                                                                                                      | :heavy_minus_sign:                                                                                                                             | JSON-encoded event payload. The decoded JSON includes a `"type"`<br/>field matching `event`, plus type-specific fields. Shape depends<br/>on `event`.<br/> |