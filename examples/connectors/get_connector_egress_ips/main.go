// List the IP addresses PipesHub connects from (for allow-lists).
// Needs PIPESHUB_BASE_URL, PIPESHUB_BEARER_AUTH.
//
// Usage (from examples/): go run ./connectors/get_connector_egress_ips <path-to-.env>
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
		log.Fatal("usage: go run ./connectors/get_connector_egress_ips <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	baseURL := strings.TrimRight(os.Getenv("PIPESHUB_BASE_URL"), "/") + "/api/v1"
	client := pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithSecurity(components.Security{BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH"))}),
	)

	res, err := client.ConnectorRegistry.GetConnectorEgressIps(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.Object, "", "  ")
	fmt.Println(string(out))
}
