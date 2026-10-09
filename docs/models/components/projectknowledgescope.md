# ProjectKnowledgeScope

Retrieval scope (app connector / knowledge-base ids) inherited by
every conversation in the project when the request itself carries no
`filters`. Same id shapes as `Filters`.



## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Apps`                                                           | []`string`                                                       | :heavy_minus_sign:                                               | Connector instance ids scoping this project's default retrieval. |
| `Kb`                                                             | []`string`                                                       | :heavy_minus_sign:                                               | Knowledge-base app ids scoping this project's default retrieval. |