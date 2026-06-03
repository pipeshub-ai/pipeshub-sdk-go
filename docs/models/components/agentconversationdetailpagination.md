# AgentConversationDetailPagination

Message pagination returned inside the `conversation` object. The
handler paginates backwards from the end of the stored message array,
then sorts the selected page in memory before serialization.



## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Page`                                                             | *int64*                                                            | :heavy_check_mark:                                                 | N/A                                                                |
| `Limit`                                                            | *int64*                                                            | :heavy_check_mark:                                                 | N/A                                                                |
| `TotalCount`                                                       | *int64*                                                            | :heavy_check_mark:                                                 | N/A                                                                |
| `TotalPages`                                                       | *int64*                                                            | :heavy_check_mark:                                                 | N/A                                                                |
| `HasNextPage`                                                      | *bool*                                                             | :heavy_check_mark:                                                 | True when older messages exist outside the returned page           |
| `HasPrevPage`                                                      | *bool*                                                             | :heavy_check_mark:                                                 | True when newer messages exist outside the returned page           |
| `MessageRange`                                                     | [components.MessageRange](../../models/components/messagerange.md) | :heavy_check_mark:                                                 | N/A                                                                |