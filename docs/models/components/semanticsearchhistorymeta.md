# SemanticSearchHistoryMeta

`requestId` comes from `req.context?.requestId` and is omitted from
the JSON when upstream middleware did not set it.



## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `RequestID`                               | `*string`                                 | :heavy_minus_sign:                        | N/A                                       |
| `Timestamp`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |
| `Duration`                                | `int64`                                   | :heavy_check_mark:                        | N/A                                       |