# MessageReferenceDatumUnion


## Supported Types

### MessageReferenceDatum

```go
messageReferenceDatumUnion := components.CreateMessageReferenceDatumUnionMessageReferenceDatum(components.MessageReferenceDatum{/* values here */})
```

### RedactedSource

```go
messageReferenceDatumUnion := components.CreateMessageReferenceDatumUnionRedactedSource(components.RedactedSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messageReferenceDatumUnion.Type {
	case components.MessageReferenceDatumUnionTypeMessageReferenceDatum:
		// messageReferenceDatumUnion.MessageReferenceDatum is populated
	case components.MessageReferenceDatumUnionTypeRedactedSource:
		// messageReferenceDatumUnion.RedactedSource is populated
	default:
		// Unknown type - use messageReferenceDatumUnion.GetUnknownRaw() for raw JSON
}
```
