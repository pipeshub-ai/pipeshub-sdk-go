// Queue a reindex of a connector record group; here only records whose indexing failed.
// Needs PIPESHUB_BASE_URL, PIPESHUB_CLIENT_ID, PIPESHUB_CLIENT_SECRET, PIPESHUB_RECORD_GROUP_ID.
//
// Usage (from examples/): go run ./KB/records/reindex_record_group <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./KB/records/reindex_record_group <path-to-.env>")
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

	res, err := client.KnowledgeBase.ReindexRecordGroup(context.Background(), os.Getenv("PIPESHUB_RECORD_GROUP_ID"), &components.ReindexRecordGroupRequestBody{
		StatusFilters: []components.IndexingStatusFilter{components.IndexingStatusFilterFailed},
	})
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.ReIndexRecordGroupResponseSchema, "", "  ")
	fmt.Println(string(out))
}
