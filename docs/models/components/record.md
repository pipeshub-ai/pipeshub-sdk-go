# Record


## Supported Types

### AppliedFilterNode

```go
record := components.CreateRecordAppliedFilterNode(components.AppliedFilterNode{/* values here */})
```

### RedactedSource

```go
record := components.CreateRecordRedactedSource(components.RedactedSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch record.Type {
	case components.RecordTypeAppliedFilterNode:
		// record.AppliedFilterNode is populated
	case components.RecordTypeRedactedSource:
		// record.RedactedSource is populated
	default:
		// Unknown type - use record.GetUnknownRaw() for raw JSON
}
```
