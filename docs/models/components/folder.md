# Folder


## Fields

| Field                         | Type                          | Required                      | Description                   |
| ----------------------------- | ----------------------------- | ----------------------------- | ----------------------------- |
| `ID`                          | **string*                     | :heavy_minus_sign:            | Unique folder identifier      |
| `Name`                        | *string*                      | :heavy_check_mark:            | Name of the folder            |
| `ParentID`                    | **string*                     | :heavy_minus_sign:            | Parent folder or KB ID        |
| `KbID`                        | *string*                      | :heavy_check_mark:            | Knowledge base ID             |
| `OrgID`                       | *string*                      | :heavy_check_mark:            | Organization ID               |
| `CreatedAtTimestamp`          | **int64*                      | :heavy_minus_sign:            | Creation timestamp            |
| `UpdatedAtTimestamp`          | **int64*                      | :heavy_minus_sign:            | Last update timestamp         |
| `IsDeleted`                   | **bool*                       | :heavy_minus_sign:            | Whether the folder is deleted |