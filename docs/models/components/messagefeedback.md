# MessageFeedback

Comprehensive feedback on an AI response. Feedback helps improve
the AI's performance and response quality over time.



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `IsHelpful`                                                                  | **bool*                                                                      | :heavy_minus_sign:                                                           | Overall helpfulness rating                                                   |
| `Ratings`                                                                    | [*components.Ratings](../../models/components/ratings.md)                    | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Categories`                                                                 | [][components.Category](../../models/components/category.md)                 | :heavy_minus_sign:                                                           | Categories of issues identified                                              |
| `Comments`                                                                   | [*components.Comments](../../models/components/comments.md)                  | :heavy_minus_sign:                                                           | N/A                                                                          |
| `CitationFeedback`                                                           | [][components.CitationFeedback](../../models/components/citationfeedback.md) | :heavy_minus_sign:                                                           | Feedback on individual citations                                             |
| `FollowUpQuestionsHelpful`                                                   | **bool*                                                                      | :heavy_minus_sign:                                                           | Were the suggested follow-up questions helpful                               |