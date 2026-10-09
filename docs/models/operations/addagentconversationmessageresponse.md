# AddAgentConversationMessageResponse


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                      | [components.HTTPMetadata](../../models/components/httpmetadata.md)              | :heavy_check_mark:                                                              | N/A                                                                             |
| `AddMessageResponse`                                                            | [*components.AddMessageResponse](../../models/components/addmessageresponse.md) | :heavy_minus_sign:                                                              | The conversation with the new question and answer.                              |
| `Headers`                                                                       | map[string][]`string`                                                           | :heavy_check_mark:                                                              | N/A                                                                             |