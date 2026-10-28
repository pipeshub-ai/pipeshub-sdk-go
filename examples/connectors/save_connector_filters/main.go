// Save a connector instance's filters (keys come from its schema).
// Needs PIPESHUB_BASE_URL, PIPESHUB_CLIENT_ID, PIPESHUB_CLIENT_SECRET, PIPESHUB_CONNECTOR_ID.
//
// Usage (from examples/): go run ./connectors/save_connector_filters <path-to-.env>
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
		log.Fatal("usage: go run ./connectors/save_connector_filters <path-to-.env>")
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

	res, err := client.ConnectorFilters.SaveConnectorFilters(context.Background(), os.Getenv("PIPESHUB_CONNECTOR_ID"), components.SaveConnectorFiltersRequest{
		Filters: map[string]any{
			"sync": map[string]any{
				"modified": map[string]any{"operator": "last_30_days"},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.Object, "", "  ")
	fmt.Println(string(out))
}
