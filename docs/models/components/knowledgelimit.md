# KnowledgeLimit

What a connected app may see: an OAuth app's limit is picked by its
creator and holds for every user who connects it; a personal access
token's by its owner. Search, chat, agent chat and every record read by
id stay inside it, and inside the connecting user's own permissions.
Omit it for everything the user can see. Sent, it must name at least one id and at
most 50 in all, since it is signed into every token the app gets.



## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Apps`                                                                     | []`string`                                                                 | :heavy_minus_sign:                                                         | Whole connectors and collections.                                          |
| `RecordGroups`                                                             | []`string`                                                                 | :heavy_minus_sign:                                                         | Record groups (a drive, space or project), with the groups nested in them. |
| `Records`                                                                  | []`string`                                                                 | :heavy_minus_sign:                                                         | Folders or records, with everything under them.                            |