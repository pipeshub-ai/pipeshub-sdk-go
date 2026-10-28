# MessageCitation


## Supported Types

### CitationReference

```go
messageCitation := components.CreateMessageCitationCitationReference(components.CitationReference{/* values here */})
```

### PopulatedCitationReference

```go
messageCitation := components.CreateMessageCitationPopulatedCitationReference(components.PopulatedCitationReference{/* values here */})
```

### RedactedSource

```go
messageCitation := components.CreateMessageCitationRedactedSource(components.RedactedSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messageCitation.Type {
	case components.MessageCitationTypeCitationReference:
		// messageCitation.CitationReference is populated
	case components.MessageCitationTypePopulatedCitationReference:
		// messageCitation.PopulatedCitationReference is populated
	case components.MessageCitationTypeRedactedSource:
		// messageCitation.RedactedSource is populated
	default:
		// Unknown type - use messageCitation.GetUnknownRaw() for raw JSON
}
```
