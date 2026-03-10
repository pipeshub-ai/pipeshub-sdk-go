# ProgressUnion

Job progress information (can be number 0-100 or object with percentage/current/total)


## Supported Types

### 

```go
progressUnion := components.CreateProgressUnionNumber(float64{/* values here */})
```

### Progress

```go
progressUnion := components.CreateProgressUnionProgress(components.Progress{/* values here */})
```

