# AuthStep

A single step in multi-factor authentication flow


## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `Order`                                                                                              | *int64*                                                                                              | :heavy_check_mark:                                                                                   | Order of the authentication step (1-3, must be unique across steps)                                  |
| `AllowedMethods`                                                                                     | [][components.AuthMethod](../../models/components/authmethod.md)                                     | :heavy_check_mark:                                                                                   | List of allowed authentication methods for this step. User can choose any one method from this list. |