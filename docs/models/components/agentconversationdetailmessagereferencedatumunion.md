# AgentConversationDetailMessageReferenceDatumUnion


## Supported Types

### AgentConversationDetailMessageReferenceDatum

```go
agentConversationDetailMessageReferenceDatumUnion := components.CreateAgentConversationDetailMessageReferenceDatumUnionAgentConversationDetailMessageReferenceDatum(components.AgentConversationDetailMessageReferenceDatum{/* values here */})
```

### RedactedSource

```go
agentConversationDetailMessageReferenceDatumUnion := components.CreateAgentConversationDetailMessageReferenceDatumUnionRedactedSource(components.RedactedSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch agentConversationDetailMessageReferenceDatumUnion.Type {
	case components.AgentConversationDetailMessageReferenceDatumUnionTypeAgentConversationDetailMessageReferenceDatum:
		// agentConversationDetailMessageReferenceDatumUnion.AgentConversationDetailMessageReferenceDatum is populated
	case components.AgentConversationDetailMessageReferenceDatumUnionTypeRedactedSource:
		// agentConversationDetailMessageReferenceDatumUnion.RedactedSource is populated
	default:
		// Unknown type - use agentConversationDetailMessageReferenceDatumUnion.GetUnknownRaw() for raw JSON
}
```
