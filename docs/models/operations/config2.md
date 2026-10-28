# Config2

OAuth app fields to change, merged over the saved ones. Keys follow
the connector type's OAuth fields; a masked secret sent back as
returned is ignored, so the saved value is kept.



## Fields

| Field                  | Type                   | Required               | Description            |
| ---------------------- | ---------------------- | ---------------------- | ---------------------- |
| `ClientID`             | `*string`              | :heavy_minus_sign:     | N/A                    |
| `ClientSecret`         | `*string`              | :heavy_minus_sign:     | N/A                    |
| `TenantID`             | `*string`              | :heavy_minus_sign:     | N/A                    |
| `AdditionalProperties` | map[string]`any`       | :heavy_minus_sign:     | N/A                    |