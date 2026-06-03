# ChatAttachmentUploadResponse

Success envelope returned by `POST /conversations/attachments/upload`
and `POST /agents/{agentKey}/conversations/attachments/upload`.



## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ConversationID`                                                                                         | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Existing conversation id echoed from the request when the upload is tied to a thread; otherwise `null`.<br/> |
| `Attachments`                                                                                            | [][components.ChatAttachmentUploadRef](../../models/components/chatattachmentuploadref.md)               | :heavy_check_mark:                                                                                       | N/A                                                                                                      |