# RegenerateOAuthAppSecretResponse

Response body for `POST /oauth-clients/{appId}/regenerate-secret` (`regenerateSecret`).



## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Message`                                                          | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                | Client secret regenerated successfully                             |
| `ClientID`                                                         | *string*                                                           | :heavy_check_mark:                                                 | OAuth client ID (unchanged)                                        |                                                                    |
| `ClientSecret`                                                     | *string*                                                           | :heavy_check_mark:                                                 | New client secret (store securely; previous secret is invalidated) |                                                                    |