# CreateUserGroupRequest

Request payload


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Name`                                             | *string*                                           | :heavy_check_mark:                                 | Display name for the group                         | Engineering Team                                   |
| `Type`                                             | [operations.Type](../../models/operations/type.md) | :heavy_check_mark:                                 | Group type determining behavior and privileges     | standard                                           |
| `Description`                                      | **string*                                          | :heavy_minus_sign:                                 | Optional description of the group's purpose        | All engineering department members                 |