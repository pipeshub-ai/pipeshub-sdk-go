# Kb


## Supported Types

### AppliedFilterNode

```go
kb := components.CreateKbAppliedFilterNode(components.AppliedFilterNode{/* values here */})
```

### RedactedSource

```go
kb := components.CreateKbRedactedSource(components.RedactedSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch kb.Type {
	case components.KbTypeAppliedFilterNode:
		// kb.AppliedFilterNode is populated
	case components.KbTypeRedactedSource:
		// kb.RedactedSource is populated
	default:
		// Unknown type - use kb.GetUnknownRaw() for raw JSON
}
```
