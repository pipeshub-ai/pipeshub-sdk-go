# CreateUserRequest

Request payload


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  | Example                                      |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `FullName`                                   | *string*                                     | :heavy_check_mark:                           | User's full display name                     | John Smith                                   |
| `Email`                                      | *string*                                     | :heavy_check_mark:                           | User's email address (must be unique)        | john.smith@company.com                       |
| `Mobile`                                     | **string*                                    | :heavy_minus_sign:                           | Mobile phone number with country code        | +15551234567                                 |
| `Designation`                                | **string*                                    | :heavy_minus_sign:                           | Job title or designation                     | Software Engineer                            |
| `SendInvite`                                 | **bool*                                      | :heavy_minus_sign:                           | Whether to send invitation email immediately |                                              |