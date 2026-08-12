# McpServer

MCP server instance linked to an agent, as projected by the graph
store on `GET /agents/{agentKey}` and `GET /agents` — same shape as
`Toolset`. MCP server nodes carry no secrets, only the attach-time
snapshot of `instanceId`/`typeId`/`name`.



## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Key`                                                                       | **string*                                                                   | :heavy_minus_sign:                                                          | MCP server instance node key in the backing graph store.                    |
| `Name`                                                                      | **string*                                                                   | :heavy_minus_sign:                                                          | MCP server attachment name (attach-time snapshot).                          |
| `DisplayName`                                                               | **string*                                                                   | :heavy_minus_sign:                                                          | Human-readable MCP server product label (for example `Jira MCP`).           |
| `TypeID`                                                                    | **string*                                                                   | :heavy_minus_sign:                                                          | Catalog server type id, when this instance came from a registered template. |
| `InstanceID`                                                                | **string*                                                                   | :heavy_minus_sign:                                                          | Admin-created MCP server instance id.                                       |
| `Tools`                                                                     | [][components.McpServerTool](../../models/components/mcpservertool.md)      | :heavy_minus_sign:                                                          | N/A                                                                         |