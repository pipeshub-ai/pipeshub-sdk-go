# CreateAgentConversationResponse

Envelope returned by `POST /agents/{agentKey}/conversations`: the
persisted agent conversation (initial user message plus the agent's
answer) and request metadata.



## Fields

| Field                                                                                                                             | Type                                                                                                                              | Required                                                                                                                          | Description                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `Conversation`                                                                                                                    | [components.AgentConversation](../../models/components/agentconversation.md)                                                      | :heavy_check_mark:                                                                                                                | A conversation with a specific AI agent. Similar to regular conversations<br/>but tied to an agent's configuration and capabilities.<br/> |
| `Meta`                                                                                                                            | [components.CreateAgentConversationResponseMeta](../../models/components/createagentconversationresponsemeta.md)                  | :heavy_check_mark:                                                                                                                | N/A                                                                                                                               |