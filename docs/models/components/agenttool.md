# AgentTool

A tool that agents can use to perform actions


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Key`                                                             | **string*                                                         | :heavy_minus_sign:                                                | Unique tool identifier                                            | web-search                                                        |
| `Name`                                                            | **string*                                                         | :heavy_minus_sign:                                                | Display name                                                      | Web Search                                                        |
| `Description`                                                     | **string*                                                         | :heavy_minus_sign:                                                | What the tool does                                                |                                                                   |
| `InputSchema`                                                     | [*components.InputSchema](../../models/components/inputschema.md) | :heavy_minus_sign:                                                | JSON Schema for tool inputs                                       |                                                                   |
| `IsEnabled`                                                       | **bool*                                                           | :heavy_minus_sign:                                                | Whether tool is currently available                               |                                                                   |