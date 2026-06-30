# ReIndexRecordGroupResponseSchema

Response returned by POST /knowledgeBase/reindex/record-group/{recordGroupId}.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Success`          | *bool*             | :heavy_check_mark: | N/A                |
| `Message`          | *string*           | :heavy_check_mark: | N/A                |
| `RecordGroupID`    | *string*           | :heavy_check_mark: | N/A                |
| `Depth`            | *int64*            | :heavy_check_mark: | N/A                |
| `Connector`        | **string*          | :heavy_minus_sign: | N/A                |
| `EventPublished`   | *bool*             | :heavy_check_mark: | N/A                |