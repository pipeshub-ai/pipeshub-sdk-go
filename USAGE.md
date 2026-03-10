<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New()

	res, err := s.UserAccount.InitAuth(ctx, components.InitAuthRequest{
		Email: "user@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.InitAuthResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->