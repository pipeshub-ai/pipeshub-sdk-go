# User

User account in an organization


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ID`                                                      | **string*                                                 | :heavy_minus_sign:                                        | Unique user identifier (MongoDB ObjectId)                 |
| `Slug`                                                    | **string*                                                 | :heavy_minus_sign:                                        | Unique slug for the user                                  |
| `OrgID`                                                   | *string*                                                  | :heavy_check_mark:                                        | Organization ID the user belongs to                       |
| `FullName`                                                | **string*                                                 | :heavy_minus_sign:                                        | Full name of the user                                     |
| `FirstName`                                               | **string*                                                 | :heavy_minus_sign:                                        | First name                                                |
| `LastName`                                                | **string*                                                 | :heavy_minus_sign:                                        | Last name                                                 |
| `MiddleName`                                              | **string*                                                 | :heavy_minus_sign:                                        | Middle name                                               |
| `Email`                                                   | *string*                                                  | :heavy_check_mark:                                        | Email address (unique, lowercase)                         |
| `Mobile`                                                  | **string*                                                 | :heavy_minus_sign:                                        | Mobile number (10-15 digits with optional +)              |
| `HasLoggedIn`                                             | **bool*                                                   | :heavy_minus_sign:                                        | Whether user has logged in at least once                  |
| `Designation`                                             | **string*                                                 | :heavy_minus_sign:                                        | Job title/designation                                     |
| `Address`                                                 | [*components.Address](../../models/components/address.md) | :heavy_minus_sign:                                        | N/A                                                       |
| `DataCollectionConsent`                                   | **bool*                                                   | :heavy_minus_sign:                                        | Whether user has consented to data collection             |
| `IsDeleted`                                               | **bool*                                                   | :heavy_minus_sign:                                        | Soft delete flag                                          |
| `DeletedBy`                                               | **string*                                                 | :heavy_minus_sign:                                        | ID of user who deleted this user                          |
| `V`                                                       | **int64*                                                  | :heavy_minus_sign:                                        | Document version (MongoDB)                                |
| `CreatedAt`                                               | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | Creation timestamp (ISO 8601)                             |
| `UpdatedAt`                                               | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | Last update timestamp (ISO 8601)                          |