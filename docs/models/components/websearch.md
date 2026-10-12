# WebSearch


## Supported Types

### AgentCreateResponseAgentWebSearch

```go
webSearch := components.CreateWebSearchAgentCreateResponseAgentWebSearch(components.AgentCreateResponseAgentWebSearch{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch webSearch.Type {
	case components.WebSearchTypeAgentCreateResponseAgentWebSearch:
		// webSearch.AgentCreateResponseAgentWebSearch is populated
	default:
		// Unknown type - use webSearch.GetUnknownRaw() for raw JSON
}
```
