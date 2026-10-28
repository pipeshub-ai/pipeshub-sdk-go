# App


## Supported Types

### AppliedFilterNode

```go
app := components.CreateAppAppliedFilterNode(components.AppliedFilterNode{/* values here */})
```

### RedactedSource

```go
app := components.CreateAppRedactedSource(components.RedactedSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch app.Type {
	case components.AppTypeAppliedFilterNode:
		// app.AppliedFilterNode is populated
	case components.AppTypeRedactedSource:
		// app.RedactedSource is populated
	default:
		// Unknown type - use app.GetUnknownRaw() for raw JSON
}
```
