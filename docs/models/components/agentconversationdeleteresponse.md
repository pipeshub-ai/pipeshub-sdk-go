# AgentConversationDeleteResponse

Envelope returned by `DELETE /agents/{agentKey}/conversations/{conversationId}`.
When the conversation does not exist, belongs to a different agent, or
was already deleted, the API still returns HTTP 200 with
`conversation: null`.



## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Message`                                                                                | [components.MessageEnum](../../models/components/messageenum.md)                         | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `Conversation`                                                                           | [components.StoredAgentConversation](../../models/components/storedagentconversation.md) | :heavy_check_mark:                                                                       | Stored agent conversation document returned by non-list endpoints.<br/>                  |