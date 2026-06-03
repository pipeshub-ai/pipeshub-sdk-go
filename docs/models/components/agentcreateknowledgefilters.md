# AgentCreateKnowledgeFilters


## Supported Types

### AgentKnowledgeFiltersParsed

```go
agentCreateKnowledgeFilters := components.CreateAgentCreateKnowledgeFiltersAgentKnowledgeFiltersParsed(components.AgentKnowledgeFiltersParsed{/* values here */})
```

### 

```go
agentCreateKnowledgeFilters := components.CreateAgentCreateKnowledgeFiltersStr(string{/* values here */})
```

### 

```go
agentCreateKnowledgeFilters := components.CreateAgentCreateKnowledgeFiltersArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch agentCreateKnowledgeFilters.Type {
	case components.AgentCreateKnowledgeFiltersTypeAgentKnowledgeFiltersParsed:
		// agentCreateKnowledgeFilters.AgentKnowledgeFiltersParsed is populated
	case components.AgentCreateKnowledgeFiltersTypeStr:
		// agentCreateKnowledgeFilters.Str is populated
	case components.AgentCreateKnowledgeFiltersTypeArrayOfAny:
		// agentCreateKnowledgeFilters.ArrayOfAny is populated
}
```
