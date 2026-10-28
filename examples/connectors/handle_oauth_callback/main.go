// Finish the OAuth flow with the code and state the provider redirected back with.
// Needs PIPESHUB_BASE_URL, PIPESHUB_BEARER_AUTH, PIPESHUB_OAUTH_CODE, PIPESHUB_OAUTH_STATE.
//
// Usage (from examples/): go run ./connectors/handle_oauth_callback <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./connectors/handle_oauth_callback <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	baseURL := strings.TrimRight(os.Getenv("PIPESHUB_BASE_URL"), "/") + "/api/v1"
	client := pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithSecurity(components.Security{BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH"))}),
	)

	res, err := client.ConnectorOAuth.HandleOAuthCallback(context.Background(), os.Getenv("PIPESHUB_OAUTH_CODE"), os.Getenv("PIPESHUB_OAUTH_STATE"), nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.Object, "", "  ")
	fmt.Println(string(out))
}
