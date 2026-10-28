# AppliedFilterNode

A single filter node selected by the user (used for display/persistence of active
filters). The gateway refuses a node that lacks any of the four fields.



## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ID`                                                           | `string`                                                       | :heavy_check_mark:                                             | Unique identifier of the filter node                           |
| `Name`                                                         | `string`                                                       | :heavy_check_mark:                                             | Display name of the filter node                                |
| `NodeType`                                                     | `string`                                                       | :heavy_check_mark:                                             | Type of the node (e.g. app, recordGroup, folder, record)       |
| `Connector`                                                    | `string`                                                       | :heavy_check_mark:                                             | Connector identifier associated with this node                 |
| `Extension`                                                    | `*string`                                                      | :heavy_minus_sign:                                             | File extension of a record node, for its icon                  |
| `MimeType`                                                     | `*string`                                                      | :heavy_minus_sign:                                             | MIME type of a record node, for its icon                       |
| `Mentioned`                                                    | `*bool`                                                        | :heavy_minus_sign:                                             | True when the node was picked with an @ mention in the message |