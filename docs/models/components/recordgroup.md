# RecordGroup


## Supported Types

### AppliedFilterNode

```go
recordGroup := components.CreateRecordGroupAppliedFilterNode(components.AppliedFilterNode{/* values here */})
```

### RedactedSource

```go
recordGroup := components.CreateRecordGroupRedactedSource(components.RedactedSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch recordGroup.Type {
	case components.RecordGroupTypeAppliedFilterNode:
		// recordGroup.AppliedFilterNode is populated
	case components.RecordGroupTypeRedactedSource:
		// recordGroup.RedactedSource is populated
	default:
		// Unknown type - use recordGroup.GetUnknownRaw() for raw JSON
}
```
