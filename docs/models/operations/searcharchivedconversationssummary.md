# SearchArchivedConversationsSummary


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `TotalMatches`                                            | `*int64`                                                  | :heavy_minus_sign:                                        | Combined match count across both collections              |
| `AssistantMatches`                                        | `*int64`                                                  | :heavy_minus_sign:                                        | Match count in the assistant (`Conversation`) collection  |
| `AgentMatches`                                            | `*int64`                                                  | :heavy_minus_sign:                                        | Match count in the agent (`AgentConversation`) collection |
| `SearchQuery`                                             | `*string`                                                 | :heavy_minus_sign:                                        | Trimmed search term that was applied                      |