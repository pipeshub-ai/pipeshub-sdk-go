# StageStateSummary

Status of one pipeline stage (e.g. `classify`) for the record's current
content revision.



## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Stage`                                               | `string`                                              | :heavy_check_mark:                                    | Pipeline stage name.                                  | classify                                              |
| `Status`                                              | `string`                                              | :heavy_check_mark:                                    | Stage status (same values as `indexingStatus`).       | COMPLETED                                             |
| `Reason`                                              | `*string`                                             | :heavy_check_mark:                                    | Failure or skip reason, when there is one.            |                                                       |
| `Attempt`                                             | `int64`                                               | :heavy_check_mark:                                    | Number of times this stage has run for this revision. |                                                       |
| `StartedAtMs`                                         | `*int64`                                              | :heavy_check_mark:                                    | N/A                                                   |                                                       |
| `FinishedAtMs`                                        | `*int64`                                              | :heavy_check_mark:                                    | N/A                                                   |                                                       |
| `UpdatedAtMs`                                         | `*int64`                                              | :heavy_check_mark:                                    | N/A                                                   |                                                       |