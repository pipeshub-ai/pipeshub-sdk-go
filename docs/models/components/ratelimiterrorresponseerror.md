# RateLimitErrorResponseError


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Code`                                     | `*string`                                  | :heavy_minus_sign:                         | Error code.                                |
| `Message`                                  | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `RetryAfter`                               | optionalnullable.OptionalNullable[`int64`] | :heavy_minus_sign:                         | Suggested seconds to wait before retrying. |