<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New()

	res, err := s.OAuthProvider.OauthToken(ctx, components.OAuthTokenRequest{
		GrantType: components.GrantTypeClientCredentials,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.OAuthTokenResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->