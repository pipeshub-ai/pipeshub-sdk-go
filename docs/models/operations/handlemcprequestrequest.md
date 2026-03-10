# HandleMCPRequestRequest

JSON-RPC 2.0 request object


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Jsonrpc`                                                              | [operations.JsonrpcRequest](../../models/operations/jsonrpcrequest.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `ID`                                                                   | [*operations.IDRequest](../../models/operations/idrequest.md)          | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Method`                                                               | *string*                                                               | :heavy_check_mark:                                                     | MCP method (e.g. initialize, tools/list, tools/call)                   |
| `Params`                                                               | [*operations.Params](../../models/operations/params.md)                | :heavy_minus_sign:                                                     | N/A                                                                    |