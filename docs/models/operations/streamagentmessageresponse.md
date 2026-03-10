# StreamAgentMessageResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `SSEEvent`                                                         | **stream.EventStream[components.SSEEvent]*                         | :heavy_minus_sign:                                                 | SSE stream established                                             |