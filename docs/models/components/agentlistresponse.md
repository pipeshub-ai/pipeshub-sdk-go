# AgentListResponse

Paginated response returned by `GET /agents`.

The Node gateway forwards the Python backend response on success. If
the backend returns a non-200 response, the gateway still returns HTTP
200 with `success: true`, an empty `agents` array, and a zeroed
pagination block derived from the requested `page` / `limit`.



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Success`                                                                        | *bool*                                                                           | :heavy_check_mark:                                                               | N/A                                                                              | true                                                                             |
| `Agents`                                                                         | [][components.AgentListItem](../../models/components/agentlistitem.md)           | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `Pagination`                                                                     | [components.AgentListPagination](../../models/components/agentlistpagination.md) | :heavy_check_mark:                                                               | Pagination block returned by `GET /agents`.                                      |                                                                                  |