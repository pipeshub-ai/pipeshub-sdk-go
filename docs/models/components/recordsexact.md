# RecordsExact


## Supported Types

### AppliedFilterNode

```go
recordsExact := components.CreateRecordsExactAppliedFilterNode(components.AppliedFilterNode{/* values here */})
```

### RedactedSource

```go
recordsExact := components.CreateRecordsExactRedactedSource(components.RedactedSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch recordsExact.Type {
	case components.RecordsExactTypeAppliedFilterNode:
		// recordsExact.AppliedFilterNode is populated
	case components.RecordsExactTypeRedactedSource:
		// recordsExact.RedactedSource is populated
	default:
		// Unknown type - use recordsExact.GetUnknownRaw() for raw JSON
}
```
