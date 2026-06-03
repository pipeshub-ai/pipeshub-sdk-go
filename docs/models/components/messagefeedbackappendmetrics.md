# MessageFeedbackAppendMetrics

Telemetry recorded server-side alongside the feedback. Always present
on append responses.



## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `TimeToFeedback`                                                                | *float64*                                                                       | :heavy_check_mark:                                                              | Milliseconds between message creation and feedback submission.<br/>Always present.<br/> |
| `UserAgent`                                                                     | **string*                                                                       | :heavy_minus_sign:                                                              | Value of the `User-Agent` request header captured server-side.                  |