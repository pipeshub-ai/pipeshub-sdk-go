# ApplicationJSONErrorResponseError


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Code`                                                              | *string*                                                            | :heavy_check_mark:                                                  | Machine-readable code (e.g. `HTTP_UNAUTHORIZED`, `HTTP_FORBIDDEN`). |
| `Message`                                                           | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `Metadata`                                                          | map[string]*any*                                                    | :heavy_minus_sign:                                                  | Optional; may appear in non-production for some errors.             |