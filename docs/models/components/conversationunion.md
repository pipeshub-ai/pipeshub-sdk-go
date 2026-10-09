# ConversationUnion


## Supported Types

### StoredAgentConversation

```go
conversationUnion := components.CreateConversationUnionStoredAgentConversation(components.StoredAgentConversation{/* values here */})
```

### AgentConversationDeleteResponseConversation

```go
conversationUnion := components.CreateConversationUnionAgentConversationDeleteResponseConversation(components.AgentConversationDeleteResponseConversation{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch conversationUnion.Type {
	case components.ConversationUnionTypeStoredAgentConversation:
		// conversationUnion.StoredAgentConversation is populated
	case components.ConversationUnionTypeAgentConversationDeleteResponseConversation:
		// conversationUnion.AgentConversationDeleteResponseConversation is populated
	default:
		// Unknown type - use conversationUnion.GetUnknownRaw() for raw JSON
}
```
