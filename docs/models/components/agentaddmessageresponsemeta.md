# AgentAddMessageResponseMeta


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `RequestID`                                                                | `*string`                                                                  | :heavy_minus_sign:                                                         | Request correlation id. Omitted when no request id was set on the context. |
| `Timestamp`                                                                | [time.Time](https://pkg.go.dev/time#Time)                                  | :heavy_check_mark:                                                         | Server timestamp when the response was sent.                               |
| `Duration`                                                                 | `int64`                                                                    | :heavy_check_mark:                                                         | Total handler duration in milliseconds.                                    |
| `RecordsUsed`                                                              | `int64`                                                                    | :heavy_check_mark:                                                         | Same value as the top-level `recordsUsed`.                                 |