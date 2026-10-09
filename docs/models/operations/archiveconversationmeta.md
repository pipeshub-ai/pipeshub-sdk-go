# ArchiveConversationMeta


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `RequestID`                                | `*string`                                  | :heavy_minus_sign:                         | Request correlation identifier             |
| `Timestamp`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | Response generation timestamp              |
| `Duration`                                 | `*int64`                                   | :heavy_minus_sign:                         | Server processing time in milliseconds     |