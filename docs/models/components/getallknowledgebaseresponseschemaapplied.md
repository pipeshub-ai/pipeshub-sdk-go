# GetAllKnowledgeBaseResponseSchemaApplied

Active filters. Empty `{}` when defaults. Keys use snake_case for sort
fields (backend convention in kb_service.py).



## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Search`                                                                       | **string*                                                                      | :heavy_minus_sign:                                                             | N/A                                                                            |
| `Permissions`                                                                  | [][components.AppliedPermission](../../models/components/appliedpermission.md) | :heavy_minus_sign:                                                             | N/A                                                                            |
| `SortBy`                                                                       | [*components.SortBy](../../models/components/sortby.md)                        | :heavy_minus_sign:                                                             | N/A                                                                            |
| `SortOrder`                                                                    | [*components.AppliedSortOrder](../../models/components/appliedsortorder.md)    | :heavy_minus_sign:                                                             | N/A                                                                            |