# CreateOAuthAppResponse

Response body for `POST /oauth-clients` (`oauth.app.controller.ts` `createApp`).
The new app (including one-time `clientSecret`) is nested under `app`.



## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Message`                                                                      | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            | OAuth app created successfully                                                 |
| `App`                                                                          | [components.OAuthAppWithSecret](../../models/components/oauthappwithsecret.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |