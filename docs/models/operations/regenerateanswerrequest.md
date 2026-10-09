# RegenerateAnswerRequest


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ConversationID`                                                              | `string`                                                                      | :heavy_check_mark:                                                            | N/A                                                                           |
| `MessageID`                                                                   | `string`                                                                      | :heavy_check_mark:                                                            | ID of the message to regenerate response for                                  |
| `Body`                                                                        | [*components.RegenerateRequest](../../models/components/regeneraterequest.md) | :heavy_minus_sign:                                                            | Request payload                                                               |