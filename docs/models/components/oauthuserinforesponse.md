# OAuthUserInfoResponse

OpenID Connect UserInfo Response.
Contains claims about the authenticated user.



## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `UserID`                                   | *string*                                   | :heavy_check_mark:                         | User ID                                    |
| `Name`                                     | **string*                                  | :heavy_minus_sign:                         | Full name                                  |
| `GivenName`                                | **string*                                  | :heavy_minus_sign:                         | First name                                 |
| `FamilyName`                               | **string*                                  | :heavy_minus_sign:                         | Last name                                  |
| `Email`                                    | **string*                                  | :heavy_minus_sign:                         | Email address                              |
| `EmailVerified`                            | **bool*                                    | :heavy_minus_sign:                         | Whether email has been verified            |
| `Picture`                                  | **string*                                  | :heavy_minus_sign:                         | Profile picture URL                        |
| `UpdatedAt`                                | **int64*                                   | :heavy_minus_sign:                         | Last profile update timestamp (Unix epoch) |