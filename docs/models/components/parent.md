# Parent

The parent this node is listed under, with what a client needs to
open it and name it; `null` exactly when `parentId` is. It is the
node's placement parent: an ancestor the user cannot open is
replaced, never named.



## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ID`                                                                   | `string`                                                               | :heavy_check_mark:                                                     | Parent node ID.                                                        |
| `NodeType`                                                             | [components.ParentNodeType](../../models/components/parentnodetype.md) | :heavy_check_mark:                                                     | Parent node type; opened as `/nodes/{nodeType}/{id}`.                  |
| `Name`                                                                 | optionalnullable.OptionalNullable[`string`]                            | :heavy_minus_sign:                                                     | Parent display name.                                                   |