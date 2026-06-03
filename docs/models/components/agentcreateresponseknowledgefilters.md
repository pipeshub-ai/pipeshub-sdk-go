# AgentCreateResponseKnowledgeFilters


## Supported Types

### 

```go
agentCreateResponseKnowledgeFilters := components.CreateAgentCreateResponseKnowledgeFiltersMapOfAny(map[string]any{/* values here */})
```

### 

```go
agentCreateResponseKnowledgeFilters := components.CreateAgentCreateResponseKnowledgeFiltersStr(string{/* values here */})
```

### 

```go
agentCreateResponseKnowledgeFilters := components.CreateAgentCreateResponseKnowledgeFiltersArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch agentCreateResponseKnowledgeFilters.Type {
	case components.AgentCreateResponseKnowledgeFiltersTypeMapOfAny:
		// agentCreateResponseKnowledgeFilters.MapOfAny is populated
	case components.AgentCreateResponseKnowledgeFiltersTypeStr:
		// agentCreateResponseKnowledgeFilters.Str is populated
	case components.AgentCreateResponseKnowledgeFiltersTypeArrayOfAny:
		// agentCreateResponseKnowledgeFilters.ArrayOfAny is populated
	default:
		// Unknown type - use agentCreateResponseKnowledgeFilters.GetUnknownRaw() for raw JSON
}
```
