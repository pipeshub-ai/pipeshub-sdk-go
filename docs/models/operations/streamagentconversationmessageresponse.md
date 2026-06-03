# StreamAgentConversationMessageResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `AgentMessageStreamSSEEvent`                                       | **stream.EventStream[components.AgentMessageStreamSSEEvent]*       | :heavy_minus_sign:                                                 | SSE stream (text/event-stream)                                     |