# CreateKBPermissionRequestBody

Request payload


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `UserIds`                                                                              | []*string*                                                                             | :heavy_minus_sign:                                                                     | User IDs to grant permission                                                           |
| `TeamIds`                                                                              | []*string*                                                                             | :heavy_minus_sign:                                                                     | Team IDs to grant permission                                                           |
| `Role`                                                                                 | [operations.CreateKBPermissionRole](../../models/operations/createkbpermissionrole.md) | :heavy_check_mark:                                                                     | Permission role to grant                                                               |