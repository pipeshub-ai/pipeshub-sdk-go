# AuthConfig

Organization authentication configuration. Supports 1-3 authentication steps for multi-factor authentication.
**Validation Rules:**
- Minimum 1 step, maximum 3 steps
- Each step must have unique order
- No duplicate methods within the same step
- No method can appear in multiple steps



## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `AuthMethods`                                                | [][components.AuthStep](../../models/components/authstep.md) | :heavy_check_mark:                                           | List of authentication steps in order                        |