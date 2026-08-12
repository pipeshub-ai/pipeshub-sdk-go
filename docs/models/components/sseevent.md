# SSEEvent

Server-Sent Event envelope for streaming chat responses. AG-UI is
the sole wire protocol.

`event` carries the AG-UI type name and `data` is a JSON-encoded
object that includes a `"type"` field matching `event`, plus
type-specific fields. Stable gateway-generated top-level outcomes:

- `RUN_FINISHED` — `{ type, result }`; `result` carries the full
  persisted `conversation` and a `meta` block with `requestId`,
  `timestamp` and `duration`.
- `RUN_ERROR` — `{ type, message, code? }`. The conversation row is
  marked FAILED before the stream closes.

Forwarded upstream lifecycle or child-run events may contain `runId`,
`threadId`, and `parentRunId`. Gateway-generated root terminal events
do not. Clients should ignore unknown event names.



## Fields

| Field                                                                                                                                          | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `Event`                                                                                                                                        | [*components.SSEEventEvent](../../models/components/sseeventevent.md)                                                                          | :heavy_minus_sign:                                                                                                                             | N/A                                                                                                                                            |
| `Data`                                                                                                                                         | **string*                                                                                                                                      | :heavy_minus_sign:                                                                                                                             | JSON-encoded event payload. The decoded JSON includes a `"type"`<br/>field matching `event`, plus type-specific fields. Shape depends<br/>on `event`.<br/> |