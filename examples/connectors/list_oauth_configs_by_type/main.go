// List the OAuth app configs of one connector type.
// Needs PIPESHUB_BASE_URL, PIPESHUB_BEARER_AUTH, PIPESHUB_CONNECTOR_TYPE.
//
// Usage (from examples/): go run ./connectors/list_oauth_configs_by_type <path-to-.env>
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
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./connectors/list_oauth_configs_by_type <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	baseURL := strings.TrimRight(os.Getenv("PIPESHUB_BASE_URL"), "/") + "/api/v1"
	client := pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithSecurity(components.Security{BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH"))}),
	)

	res, err := client.OAuthConfiguration.ListOAuthConfigsByType(context.Background(), operations.ListOAuthConfigsByTypeRequest{
		ConnectorType: os.Getenv("PIPESHUB_CONNECTOR_TYPE"),
	})
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.Object, "", "  ")
	fmt.Println(string(out))
}
