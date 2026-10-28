# AgentCreateKnowledgeFilters

The source's limit, as an object or its JSON text. An empty string or an empty
array means no limit. A non-empty array, `null`, or text that is not the JSON of
an object is refused with `400` (`VALIDATION_ERROR`, "Knowledge filters must be
an object, with recordGroups and records as lists of ids").



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
