# AgentConversationDetailMessageCitationUnion


## Supported Types

### AgentConversationDetailMessageCitation

```go
agentConversationDetailMessageCitationUnion := components.CreateAgentConversationDetailMessageCitationUnionAgentConversationDetailMessageCitation(components.AgentConversationDetailMessageCitation{/* values here */})
```

### RedactedSource

```go
agentConversationDetailMessageCitationUnion := components.CreateAgentConversationDetailMessageCitationUnionRedactedSource(components.RedactedSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch agentConversationDetailMessageCitationUnion.Type {
	case components.AgentConversationDetailMessageCitationUnionTypeAgentConversationDetailMessageCitation:
		// agentConversationDetailMessageCitationUnion.AgentConversationDetailMessageCitation is populated
	case components.AgentConversationDetailMessageCitationUnionTypeRedactedSource:
		// agentConversationDetailMessageCitationUnion.RedactedSource is populated
	default:
		// Unknown type - use agentConversationDetailMessageCitationUnion.GetUnknownRaw() for raw JSON
}
```
