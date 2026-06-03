# AgentConversationArchiveMeta

Request-scoped metadata returned by the archive route. `requestId` is
omitted when upstream middleware did not attach one.



## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `RequestID`                               | **string*                                 | :heavy_minus_sign:                        | N/A                                       |
| `Timestamp`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |
| `Duration`                                | *int64*                                   | :heavy_check_mark:                        | N/A                                       |