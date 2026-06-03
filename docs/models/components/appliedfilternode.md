# AppliedFilterNode

A single filter node selected by the user (used for display/persistence of active filters)


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ID`                                                     | **string*                                                | :heavy_minus_sign:                                       | Unique identifier of the filter node                     |
| `Name`                                                   | **string*                                                | :heavy_minus_sign:                                       | Display name of the filter node                          |
| `NodeType`                                               | **string*                                                | :heavy_minus_sign:                                       | Type of the node (e.g. app, recordGroup, folder, record) |
| `Connector`                                              | **string*                                                | :heavy_minus_sign:                                       | Connector identifier associated with this node           |