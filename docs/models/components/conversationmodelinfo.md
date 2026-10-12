# ConversationModelInfo

AI model configuration recorded against a conversation or message.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ModelKey`                                                     | `*string`                                                      | :heavy_minus_sign:                                             | Stable identifier of the configured model record               |
| `ModelName`                                                    | `*string`                                                      | :heavy_minus_sign:                                             | Provider-facing model name (e.g. `gpt-5.6-luna`)               |
| `ModelProvider`                                                | `*string`                                                      | :heavy_minus_sign:                                             | Provider key (e.g. `openai`, `anthropic`)                      |
| `ModelFriendlyName`                                            | `*string`                                                      | :heavy_minus_sign:                                             | Human-readable display name                                    |
| `ChatMode`                                                     | `*string`                                                      | :heavy_minus_sign:                                             | Chat mode used for this turn (e.g. `quick`, `internal_search`) |