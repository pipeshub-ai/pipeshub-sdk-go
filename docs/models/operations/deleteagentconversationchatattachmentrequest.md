# DeleteAgentConversationChatAttachmentRequest


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `AgentKey`                                                                     | *string*                                                                       | :heavy_check_mark:                                                             | Agent key path parameter. Must be non-empty.                                   |
| `RecordID`                                                                     | *string*                                                                       | :heavy_check_mark:                                                             | Attachment record id (from the upload response). Must be non-blank after trim. |