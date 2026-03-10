# CreateTeamRequest

Request payload


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Name`                                                       | *string*                                                     | :heavy_check_mark:                                           | Team display name (must be unique in org)                    | Engineering Team                                             |
| `Description`                                                | **string*                                                    | :heavy_minus_sign:                                           | Team description and purpose                                 | Core engineering team for product development                |
| `UserRoles`                                                  | [][components.UserRole](../../models/components/userrole.md) | :heavy_minus_sign:                                           | Optional initial members with roles                          |                                                              |