// Create an RSS connector instance (team scope, manual sync) for one feed.
// Needs PIPESHUB_BASE_URL, PIPESHUB_BEARER_AUTH.
//
// Usage (from examples/): go run ./connectors/create_connector_instance <path-to-.env>
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
		log.Fatal("usage: go run ./connectors/create_connector_instance <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	baseURL := strings.TrimRight(os.Getenv("PIPESHUB_BASE_URL"), "/") + "/api/v1"
	client := pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithSecurity(components.Security{BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH"))}),
	)

	res, err := client.ConnectorInstances.CreateConnectorInstance(context.Background(), components.CreateConnectorRequest{
		ConnectorType: "RSS",
		InstanceName:  "RSS quick start",
		Scope:         components.ConnectorScopeTeam,
		AuthType:      pipeshub.Pointer(components.ConnectorAuthTypeNone),
		Config: &components.CreateConnectorRequestConfig{
			Auth: &components.ConnectorAuthConfig{},
			Sync: &components.ConnectorSyncConfig{
				SelectedStrategy: pipeshub.Pointer(components.ConnectorSyncConfigSelectedStrategyManual),
				// RSS's own sync fields come from its registry schema; they travel as additional properties.
				AdditionalProperties: map[string]any{"feed_urls": "http://feeds.arstechnica.com/arstechnica/index/"},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.Object, "", "  ")
	fmt.Println(string(out))
}
