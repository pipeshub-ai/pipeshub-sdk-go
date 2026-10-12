# CitationUnion


## Supported Types

### CitationReference

```go
citationUnion := components.CreateCitationUnionCitationReference(components.CitationReference{/* values here */})
```

### PopulatedCitationReference

```go
citationUnion := components.CreateCitationUnionPopulatedCitationReference(components.PopulatedCitationReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch citationUnion.Type {
	case components.CitationUnionTypeCitationReference:
		// citationUnion.CitationReference is populated
	case components.CitationUnionTypePopulatedCitationReference:
		// citationUnion.PopulatedCitationReference is populated
	default:
		// Unknown type - use citationUnion.GetUnknownRaw() for raw JSON
}
```
