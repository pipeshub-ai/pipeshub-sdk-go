# OrgAuthConfigCreateRequest

Request to create initial organization auth configuration


## Fields

| Field                         | Type                          | Required                      | Description                   |
| ----------------------------- | ----------------------------- | ----------------------------- | ----------------------------- |
| `ContactEmail`                | `string`                      | :heavy_check_mark:            | Organization contact email    |
| `RegisteredName`              | `string`                      | :heavy_check_mark:            | Organization registered name  |
| `AdminFullName`               | `string`                      | :heavy_check_mark:            | Admin user full name          |
| `SendEmail`                   | `*bool`                       | :heavy_minus_sign:            | Whether to send welcome email |