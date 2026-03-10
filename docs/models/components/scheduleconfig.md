# ScheduleConfig

Schedule configuration for crawling jobs. The structure varies based on <code>scheduleType</code>.<br><br>
<b>Schedule Type Configurations:</b><br>
<ul>
<li><b>hourly:</b> <code>minute</code>, <code>interval</code> (optional)</li>
<li><b>daily:</b> <code>hour</code>, <code>minute</code></li>
<li><b>weekly:</b> <code>daysOfWeek</code>, <code>hour</code>, <code>minute</code></li>
<li><b>monthly:</b> <code>dayOfMonth</code>, <code>hour</code>, <code>minute</code></li>
<li><b>custom:</b> <code>cronExpression</code>, <code>description</code> (optional)</li>
<li><b>once:</b> <code>scheduledTime</code></li>
</ul>



## Supported Types

### HourlyScheduleConfig

```go
scheduleConfig := components.CreateScheduleConfigHourly(components.HourlyScheduleConfig{/* values here */})
```

### DailyScheduleConfig

```go
scheduleConfig := components.CreateScheduleConfigDaily(components.DailyScheduleConfig{/* values here */})
```

### WeeklyScheduleConfig

```go
scheduleConfig := components.CreateScheduleConfigWeekly(components.WeeklyScheduleConfig{/* values here */})
```

### MonthlyScheduleConfig

```go
scheduleConfig := components.CreateScheduleConfigMonthly(components.MonthlyScheduleConfig{/* values here */})
```

### CustomScheduleConfig

```go
scheduleConfig := components.CreateScheduleConfigCustom(components.CustomScheduleConfig{/* values here */})
```

### OnceScheduleConfig

```go
scheduleConfig := components.CreateScheduleConfigOnce(components.OnceScheduleConfig{/* values here */})
```

