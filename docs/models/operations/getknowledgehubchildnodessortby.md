# GetKnowledgeHubChildNodesSortBy

Field to sort results by. Omitted → default `updatedAt`.
Unknown value → silently falls back to `name`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.GetKnowledgeHubChildNodesSortByName
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `GetKnowledgeHubChildNodesSortByName`      | name                                       |
| `GetKnowledgeHubChildNodesSortByCreatedAt` | createdAt                                  |
| `GetKnowledgeHubChildNodesSortByUpdatedAt` | updatedAt                                  |
| `GetKnowledgeHubChildNodesSortBySize`      | size                                       |
| `GetKnowledgeHubChildNodesSortByType`      | type                                       |