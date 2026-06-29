# ReIndexRecordResponseSchema

Response returned by POST /knowledgeBase/reindex/record/{recordId}.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Success`          | *bool*             | :heavy_check_mark: | N/A                |
| `Message`          | *string*           | :heavy_check_mark: | N/A                |
| `RecordID`         | **string*          | :heavy_minus_sign: | N/A                |
| `RecordName`       | **string*          | :heavy_minus_sign: | N/A                |
| `Connector`        | **string*          | :heavy_minus_sign: | N/A                |
| `EventPublished`   | *bool*             | :heavy_check_mark: | N/A                |
| `UserRole`         | **string*          | :heavy_minus_sign: | N/A                |
| `Depth`            | *int64*            | :heavy_check_mark: | N/A                |