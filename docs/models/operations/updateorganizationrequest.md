# UpdateOrganizationRequest

Request payload


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `RegisteredName`                                          | **string*                                                 | :heavy_minus_sign:                                        | Official registered/legal name                            | Acme Corporation Inc.                                     |
| `ShortName`                                               | **string*                                                 | :heavy_minus_sign:                                        | Short display name for UI                                 | Acme Corp                                                 |
| `PhoneNumber`                                             | **string*                                                 | :heavy_minus_sign:                                        | Contact phone number (international format)               | +15551234567                                              |
| `PermanentAddress`                                        | [*components.Address](../../models/components/address.md) | :heavy_minus_sign:                                        | N/A                                                       |                                                           |