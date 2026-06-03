# AgentArchivedConversationSummary

Archive counts and bounds for the current result page returned by
`GET /agents/{agentKey}/conversations/show/archives`.



## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `TotalArchived`                                                                          | **int64*                                                                                 | :heavy_minus_sign:                                                                       | Total archived conversations matching the filter                                         |
| `OldestArchive`                                                                          | [*time.Time](https://pkg.go.dev/time#Time)                                               | :heavy_minus_sign:                                                                       | Archive timestamp of the first item in the current page. Omitted when the page is empty. |
| `NewestArchive`                                                                          | [*time.Time](https://pkg.go.dev/time#Time)                                               | :heavy_minus_sign:                                                                       | Archive timestamp of the last item in the current page. Omitted when the page is empty.  |