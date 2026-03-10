# SetPlatformSettingsRequest

Request body for Update platform settings


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `FileUploadMaxSizeBytes`                                                 | *int64*                                                                  | :heavy_check_mark:                                                       | Maximum file upload size in bytes                                        | 31457280                                                                 |
| `FeatureFlags`                                                           | map[string]*bool*                                                        | :heavy_check_mark:                                                       | Key-value map of feature flags. Set to true to enable, false to disable. | {<br/>"ENABLE_BETA_CONNECTORS": false<br/>}                              |