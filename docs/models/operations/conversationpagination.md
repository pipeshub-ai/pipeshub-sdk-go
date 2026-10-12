# ConversationPagination

Pagination over the conversation's messages. Messages are paginated backwards
(newest first), so `messageRange.start`/`messageRange.end` refer to 1-based
positions within the full message list.



## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Page`                                                              | `*int64`                                                            | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Limit`                                                             | `*int64`                                                            | :heavy_minus_sign:                                                  | N/A                                                                 |
| `TotalCount`                                                        | `*int64`                                                            | :heavy_minus_sign:                                                  | Total number of messages in the conversation                        |
| `TotalPages`                                                        | `*int64`                                                            | :heavy_minus_sign:                                                  | N/A                                                                 |
| `HasNextPage`                                                       | `*bool`                                                             | :heavy_minus_sign:                                                  | True if there are older messages available                          |
| `HasPrevPage`                                                       | `*bool`                                                             | :heavy_minus_sign:                                                  | True if there are newer messages available                          |
| `MessageRange`                                                      | [*operations.MessageRange](../../models/operations/messagerange.md) | :heavy_minus_sign:                                                  | N/A                                                                 |