# ChatAttachmentUploadRef

Concrete attachment metadata returned by `POST /conversations/attachments/upload`
(or the equivalent agent route).



## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `RecordID`                                                    | *string*                                                      | :heavy_check_mark:                                            | Server-assigned attachment record id.                         |
| `RecordName`                                                  | *string*                                                      | :heavy_check_mark:                                            | Original filename stored for the attachment.                  |
| `MimeType`                                                    | *string*                                                      | :heavy_check_mark:                                            | MIME type of the uploaded file.                               |
| `Extension`                                                   | *string*                                                      | :heavy_check_mark:                                            | File extension derived by the backend.                        |
| `VirtualRecordID`                                             | *string*                                                      | :heavy_check_mark:                                            | Synthetic record id used by the graph layer.                  |
| `OcrMode`                                                     | **string*                                                     | :heavy_minus_sign:                                            | Optional backend-reported processing mode for the attachment. |