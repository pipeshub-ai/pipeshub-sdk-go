# WebSearchProviderType

Supported web search provider. crawl4ai only reads pages (it backs up the built-in page reader) and cannot be the default; firecrawl both searches and reads pages. searxng, firecrawl and crawl4ai take a baseUrl; firecrawl defaults to its cloud API.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.WebSearchProviderTypeDuckduckgo

// Open enum: custom values can be created with a direct type cast
custom := components.WebSearchProviderType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `WebSearchProviderTypeDuckduckgo` | duckduckgo                        |
| `WebSearchProviderTypeSerper`     | serper                            |
| `WebSearchProviderTypeTavily`     | tavily                            |
| `WebSearchProviderTypeExa`        | exa                               |
| `WebSearchProviderTypeBrave`      | brave                             |
| `WebSearchProviderTypeSearxng`    | searxng                           |
| `WebSearchProviderTypeFirecrawl`  | firecrawl                         |
| `WebSearchProviderTypeCrawl4ai`   | crawl4ai                          |