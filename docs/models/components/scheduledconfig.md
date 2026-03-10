# ScheduledConfig

Configuration for scheduled sync strategy


## Fields

| Field                                   | Type                                    | Required                                | Description                             | Example                                 |
| --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| `IntervalMinutes`                       | **int64*                                | :heavy_minus_sign:                      | Sync interval in minutes                | 60                                      |
| `CronExpression`                        | **string*                               | :heavy_minus_sign:                      | Cron expression for advanced scheduling | 0 */6 * * *                             |
| `Timezone`                              | **string*                               | :heavy_minus_sign:                      | Timezone for scheduled sync             | America/New_York                        |