# CreatePatResponse

Response body for `POST /personal-access-tokens` (`pat.controller.ts` `createToken`).


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Message`                                                            | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  | Personal access token created successfully                           |
| `Token`                                                              | [components.PatWithSecret](../../models/components/patwithsecret.md) | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |