# AgentFilters

Knowledge scope filter as stored on the graph edge. The Node `getAgent`
handler proxies this field unchanged from the AI service (only `agent.id`
is stripped). May be a JSON string (typical graph storage) or an object.
Prefer `filtersParsed` on GET for a guaranteed parsed object with the
same keys as the object branch below.



## Supported Types

### AgentKnowledgeFiltersParsed

```go
agentFilters := components.CreateAgentFiltersAgentKnowledgeFiltersParsed(components.AgentKnowledgeFiltersParsed{/* values here */})
```

### 

```go
agentFilters := components.CreateAgentFiltersStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch agentFilters.Type {
	case components.AgentFiltersTypeAgentKnowledgeFiltersParsed:
		// agentFilters.AgentKnowledgeFiltersParsed is populated
	case components.AgentFiltersTypeStr:
		// agentFilters.Str is populated
	default:
		// Unknown type - use agentFilters.GetUnknownRaw() for raw JSON
}
```
