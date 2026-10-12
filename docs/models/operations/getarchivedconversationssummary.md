# GetArchivedConversationsSummary


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `TotalArchived`                                         | `*int64`                                                | :heavy_minus_sign:                                      | Total archived conversations matching the filter        |
| `OldestArchive`                                         | [*time.Time](https://pkg.go.dev/time#Time)              | :heavy_minus_sign:                                      | Archive timestamp of the first item in the current page |
| `NewestArchive`                                         | [*time.Time](https://pkg.go.dev/time#Time)              | :heavy_minus_sign:                                      | Archive timestamp of the last item in the current page  |