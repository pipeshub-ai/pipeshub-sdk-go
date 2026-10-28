// Change a connector's sync strategy (filters can go in the same call).
// Needs PIPESHUB_BASE_URL, PIPESHUB_CLIENT_ID, PIPESHUB_CLIENT_SECRET, PIPESHUB_CONNECTOR_ID.
//
// Usage (from examples/): go run ./connectors/update_connector_filters_sync_config <path-to-.env>
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
		log.Fatal("usage: go run ./connectors/update_connector_filters_sync_config <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	baseURL := strings.TrimRight(os.Getenv("PIPESHUB_BASE_URL"), "/") + "/api/v1"
	tokenURL := os.Getenv("PIPESHUB_TOKEN_URL") // optional
	if tokenURL == "" {
		tokenURL = baseURL + "/oauth2/token"
	}
	client := pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithSecurity(components.Security{Oauth2: &components.SchemeOauth2{
			ClientID:     os.Getenv("PIPESHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("PIPESHUB_CLIENT_SECRET"),
			TokenURL:     tokenURL,
		}}),
	)

	res, err := client.ConnectorConfiguration.UpdateConnectorFiltersSyncConfig(context.Background(), os.Getenv("PIPESHUB_CONNECTOR_ID"), components.CreateUpdateConnectorFiltersSyncRequestUnionUpdateConnectorFiltersSyncRequest1(components.UpdateConnectorFiltersSyncRequest1{
		Sync: components.ConnectorSyncConfig{
			SelectedStrategy: pipeshub.Pointer(components.ConnectorSyncConfigSelectedStrategyManual),
		},
	}))
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.Object, "", "  ")
	fmt.Println(string(out))
}
