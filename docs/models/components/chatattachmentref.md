# ChatAttachmentRef

Reference to an attachment produced by `POST /conversations/attachments/upload`
(or the equivalent agent route). Include in create/stream/message bodies
so the turn is sent with uploaded files.



## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `RecordID`                                              | *string*                                                | :heavy_check_mark:                                      | Attachment record id returned from the upload endpoint. |
| `RecordName`                                            | **string*                                               | :heavy_minus_sign:                                      | Original display name of the file when known.           |
| `MimeType`                                              | **string*                                               | :heavy_minus_sign:                                      | MIME type of the uploaded file.                         |
| `Extension`                                             | **string*                                               | :heavy_minus_sign:                                      | File extension (e.g. `pdf`).                            |
| `VirtualRecordID`                                       | **string*                                               | :heavy_minus_sign:                                      | Optional synthetic record id used by the graph layer.   |