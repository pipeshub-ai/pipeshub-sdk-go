# GetUploadLimitsResponseBody

Upload limits retrieved


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `MaxFilesPerRequest`                       | **int64*                                   | :heavy_minus_sign:                         | Maximum number of files per upload request | 1000                                       |
| `MaxFileSizeBytes`                         | **int64*                                   | :heavy_minus_sign:                         | Maximum file size in bytes (default 30MB)  | 31457280                                   |