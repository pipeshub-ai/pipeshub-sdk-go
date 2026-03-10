# Team


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ID`                                                         | **string*                                                    | :heavy_minus_sign:                                           | Unique team identifier                                       |
| `Name`                                                       | *string*                                                     | :heavy_check_mark:                                           | Team name                                                    |
| `Description`                                                | **string*                                                    | :heavy_minus_sign:                                           | Team description                                             |
| `OrgID`                                                      | **string*                                                    | :heavy_minus_sign:                                           | Organization ID                                              |
| `UserRoles`                                                  | [][components.UserRole](../../models/components/userrole.md) | :heavy_minus_sign:                                           | Users and their roles in the team                            |
| `CreatedAt`                                                  | [*time.Time](https://pkg.go.dev/time#Time)                   | :heavy_minus_sign:                                           | Creation timestamp (ISO 8601)                                |
| `UpdatedAt`                                                  | [*time.Time](https://pkg.go.dev/time#Time)                   | :heavy_minus_sign:                                           | Last update timestamp (ISO 8601)                             |