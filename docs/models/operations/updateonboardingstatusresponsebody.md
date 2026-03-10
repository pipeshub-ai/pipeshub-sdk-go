# UpdateOnboardingStatusResponseBody

Onboarding status updated successfully


## Fields

| Field                                  | Type                                   | Required                               | Description                            | Example                                |
| -------------------------------------- | -------------------------------------- | -------------------------------------- | -------------------------------------- | -------------------------------------- |
| `Success`                              | **bool*                                | :heavy_minus_sign:                     | N/A                                    | true                                   |
| `Message`                              | **string*                              | :heavy_minus_sign:                     | N/A                                    | Step 'invite_team' marked as complete  |
| `IsOnboardingComplete`                 | **bool*                                | :heavy_minus_sign:                     | Whether all onboarding is now complete | false                                  |
| `NextStep`                             | **string*                              | :heavy_minus_sign:                     | Next step to complete (null if done)   | connect_integrations                   |