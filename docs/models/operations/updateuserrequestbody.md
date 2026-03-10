# UpdateUserRequestBody

Request payload


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `FullName`                                                | **string*                                                 | :heavy_minus_sign:                                        | Full display name                                         | John Smith                                                |
| `FirstName`                                               | **string*                                                 | :heavy_minus_sign:                                        | First name only                                           | John                                                      |
| `LastName`                                                | **string*                                                 | :heavy_minus_sign:                                        | Last name only                                            | Smith                                                     |
| `Email`                                                   | **string*                                                 | :heavy_minus_sign:                                        | Email address (admin only)                                | john.smith@company.com                                    |
| `Mobile`                                                  | **string*                                                 | :heavy_minus_sign:                                        | Mobile phone with country code                            | +15551234567                                              |
| `Designation`                                             | **string*                                                 | :heavy_minus_sign:                                        | Job title or role                                         | Senior Software Engineer                                  |
| `Address`                                                 | [*components.Address](../../models/components/address.md) | :heavy_minus_sign:                                        | N/A                                                       |                                                           |