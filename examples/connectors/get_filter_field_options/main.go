// List the values you can pick for one filter field (a channel, space, folder...).
// Needs PIPESHUB_BASE_URL, PIPESHUB_BEARER_AUTH, PIPESHUB_CONNECTOR_ID, PIPESHUB_FILTER_KEY.
//
// Usage (from examples/): go run ./connectors/get_filter_field_options <path-to-.env>
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
		log.Fatal("usage: go run ./connectors/get_filter_field_options <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	baseURL := strings.TrimRight(os.Getenv("PIPESHUB_BASE_URL"), "/") + "/api/v1"
	client := pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithSecurity(components.Security{BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH"))}),
	)

	res, err := client.ConnectorFilters.GetFilterFieldOptions(context.Background(), operations.GetFilterFieldOptionsRequest{
		ConnectorID: os.Getenv("PIPESHUB_CONNECTOR_ID"),
		FilterKey:   os.Getenv("PIPESHUB_FILTER_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.Object, "", "  ")
	fmt.Println(string(out))
}
