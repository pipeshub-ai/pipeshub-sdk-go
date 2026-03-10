# GetStorageConfigResponseBody

Storage configuration retrieved


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `StorageType`                                                     | [*operations.StorageType](../../models/operations/storagetype.md) | :heavy_minus_sign:                                                | Currently configured storage type                                 |
| `MountName`                                                       | **string*                                                         | :heavy_minus_sign:                                                | Mount point name (Local)                                          |
| `BaseURL`                                                         | **string*                                                         | :heavy_minus_sign:                                                | Base URL for files (Local)                                        |
| `AccessKeyID`                                                     | **string*                                                         | :heavy_minus_sign:                                                | AWS access key ID (S3)                                            |
| `SecretAccessKey`                                                 | **string*                                                         | :heavy_minus_sign:                                                | AWS secret access key (S3)                                        |
| `Region`                                                          | **string*                                                         | :heavy_minus_sign:                                                | AWS region (S3)                                                   |
| `BucketName`                                                      | **string*                                                         | :heavy_minus_sign:                                                | S3 bucket name (S3)                                               |
| `ContainerName`                                                   | **string*                                                         | :heavy_minus_sign:                                                | Container name (Azure Blob)                                       |
| `AccountName`                                                     | **string*                                                         | :heavy_minus_sign:                                                | Storage account name (Azure Blob)                                 |