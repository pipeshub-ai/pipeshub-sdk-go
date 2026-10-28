# ConversationConversationError


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `ID`                                              | `*string`                                         | :heavy_minus_sign:                                | Sub-document identifier auto-assigned by MongoDB. |
| `Message`                                         | `string`                                          | :heavy_check_mark:                                | N/A                                               |
| `ErrorType`                                       | `*string`                                         | :heavy_minus_sign:                                | N/A                                               |
| `Timestamp`                                       | [*time.Time](https://pkg.go.dev/time#Time)        | :heavy_minus_sign:                                | N/A                                               |
| `MessageID`                                       | `*string`                                         | :heavy_minus_sign:                                | N/A                                               |
| `Stack`                                           | `*string`                                         | :heavy_minus_sign:                                | N/A                                               |
| `Metadata`                                        | map[string]`any`                                  | :heavy_minus_sign:                                | N/A                                               |