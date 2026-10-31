<!-- Start SDK Example Usage [usage] -->
### Retrieve passages with citations

```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New(
		pipeshub.WithSecurity(components.Security{
			BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
		}),
	)

	res, err := s.SemanticSearch.Retrieve(ctx, components.RetrieveRequest{
		Query: "how do we roll back a canary?",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.RetrieveResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->