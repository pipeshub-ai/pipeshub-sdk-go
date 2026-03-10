# BulkInviteUsersResponseBody

Bulk invitation processed


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Success`                                                  | **bool*                                                    | :heavy_minus_sign:                                         | N/A                                                        | true                                                       |
| `Message`                                                  | **string*                                                  | :heavy_minus_sign:                                         | N/A                                                        | Bulk invitation completed                                  |
| `Invited`                                                  | **int64*                                                   | :heavy_minus_sign:                                         | Number of successful invitations sent                      | 8                                                          |
| `Skipped`                                                  | **int64*                                                   | :heavy_minus_sign:                                         | Number of existing users skipped                           | 1                                                          |
| `Failed`                                                   | **int64*                                                   | :heavy_minus_sign:                                         | Number of failed invitations                               | 1                                                          |
| `Failures`                                                 | [][operations.Failure](../../models/operations/failure.md) | :heavy_minus_sign:                                         | Details of failed invitations                              |                                                            |