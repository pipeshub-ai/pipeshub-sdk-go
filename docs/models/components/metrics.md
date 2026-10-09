# Metrics

Optional telemetry captured alongside the feedback


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `TimeToFeedback`                                   | `*float64`                                         | :heavy_minus_sign:                                 | Time from response delivery to feedback submission |
| `UserInteractionTime`                              | `*float64`                                         | :heavy_minus_sign:                                 | Total time the user spent reviewing the response   |
| `FeedbackSessionID`                                | `*string`                                          | :heavy_minus_sign:                                 | N/A                                                |
| `UserAgent`                                        | `*string`                                          | :heavy_minus_sign:                                 | N/A                                                |
| `Platform`                                         | `*string`                                          | :heavy_minus_sign:                                 | N/A                                                |