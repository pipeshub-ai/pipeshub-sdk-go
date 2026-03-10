# CreateOrganizationRequest

Request payload


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `AccountType`                                                    | [operations.AccountType](../../models/operations/accounttype.md) | :heavy_check_mark:                                               | Type of organization account                                     | business                                                         |
| `ShortName`                                                      | **string*                                                        | :heavy_minus_sign:                                               | Short display name for the organization                          | Acme                                                             |
| `RegisteredName`                                                 | **string*                                                        | :heavy_minus_sign:                                               | Official registered name (for business accounts)                 | Acme Corporation Inc.                                            |
| `ContactEmail`                                                   | *string*                                                         | :heavy_check_mark:                                               | Primary contact email (also used as admin email)                 | admin@acme.com                                                   |
| `AdminFullName`                                                  | *string*                                                         | :heavy_check_mark:                                               | Full name of the first admin user                                | John Smith                                                       |
| `Password`                                                       | *string*                                                         | :heavy_check_mark:                                               | Password for the admin account (min 8 chars)                     | SecurePassword123!                                               |