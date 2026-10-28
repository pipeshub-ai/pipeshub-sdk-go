# GetKnowledgeHubChildNodesSortBy

Field to sort results by. Omitted → default `updatedAt`.
Any other value is answered `400` ("Invalid sort_by: ...").


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
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