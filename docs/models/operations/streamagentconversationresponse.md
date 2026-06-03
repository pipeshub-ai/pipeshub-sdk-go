# StreamAgentConversationResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `AgentStreamSSEEvent`                                              | **stream.EventStream[components.AgentStreamSSEEvent]*              | :heavy_minus_sign:                                                 | SSE stream (text/event-stream)                                     |