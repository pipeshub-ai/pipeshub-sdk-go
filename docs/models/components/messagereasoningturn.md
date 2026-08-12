# MessageReasoningTurn

One model turn's chain-of-thought. Persisted only when reasoning
persistence is enabled; the array is empty otherwise.



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `MessageID`        | **string*          | :heavy_minus_sign: | N/A                |
| `TurnIndex`        | **float64*         | :heavy_minus_sign: | N/A                |
| `Content`          | *string*           | :heavy_check_mark: | N/A                |