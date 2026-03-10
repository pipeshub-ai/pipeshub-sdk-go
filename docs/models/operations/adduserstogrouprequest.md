# AddUsersToGroupRequest

Request payload


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `GroupID`                                                  | *string*                                                   | :heavy_check_mark:                                         | ID of the group to add users to                            | 507f1f77bcf86cd799439011                                   |
| `UserIds`                                                  | []*string*                                                 | :heavy_check_mark:                                         | Array of user IDs to add to the group                      | [<br/>"507f1f77bcf86cd799439012",<br/>"507f1f77bcf86cd799439013"<br/>] |