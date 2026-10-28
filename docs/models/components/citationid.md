# CitationID


## Supported Types

### PersistedSemanticSearchCitation

```go
citationID := components.CreateCitationIDPersistedSemanticSearchCitation(components.PersistedSemanticSearchCitation{/* values here */})
```

### RedactedSource

```go
citationID := components.CreateCitationIDRedactedSource(components.RedactedSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch citationID.Type {
	case components.CitationIDTypePersistedSemanticSearchCitation:
		// citationID.PersistedSemanticSearchCitation is populated
	case components.CitationIDTypeRedactedSource:
		// citationID.RedactedSource is populated
	default:
		// Unknown type - use citationID.GetUnknownRaw() for raw JSON
}
```
