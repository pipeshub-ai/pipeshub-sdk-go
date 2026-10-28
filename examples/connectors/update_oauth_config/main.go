// Rename an OAuth app config.
// Needs PIPESHUB_BASE_URL, PIPESHUB_BEARER_AUTH, PIPESHUB_CONNECTOR_TYPE, PIPESHUB_OAUTH_CONFIG_ID.
//
// Usage (from examples/): go run ./connectors/update_oauth_config <path-to-.env>
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
		log.Fatal("usage: go run ./connectors/update_oauth_config <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	baseURL := strings.TrimRight(os.Getenv("PIPESHUB_BASE_URL"), "/") + "/api/v1"
	client := pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithSecurity(components.Security{BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH"))}),
	)

	res, err := client.OAuthConfiguration.UpdateOAuthConfig(context.Background(), os.Getenv("PIPESHUB_CONNECTOR_TYPE"), os.Getenv("PIPESHUB_OAUTH_CONFIG_ID"), operations.CreateUpdateOAuthConfigRequestBodyRequestBody1(
		operations.RequestBody1{OauthInstanceName: "Renamed OAuth app"},
	))
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.Object, "", "  ")
	fmt.Println(string(out))
}
