package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

const connectorName = "abc news"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run . <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	client, err := auth.NewClient(
		os.Getenv("PIPESHUB_TEST_USER_EMAIL"),
		os.Getenv("PIPESHUB_TEST_USER_PASSWORD"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	nodes, err := client.KnowledgeHub.GetKnowledgeHubRootNodes(ctx, operations.GetKnowledgeHubRootNodesRequest{})
	if err != nil {
		log.Fatalf("get knowledge hub root nodes: %v", err)
	}
	if nodes == nil || nodes.KnowledgeHubNodesResponse == nil {
		log.Fatal("get knowledge hub root nodes: empty response")
	}
	var connectorID string
	for _, n := range nodes.KnowledgeHubNodesResponse.GetItems() {
		if n.Name == connectorName && n.Origin == components.OriginConnector {
			connectorID = n.ID
			break
		}
	}
	if connectorID == "" {
		log.Fatalf("connector %q not found", connectorName)
	}

	res, err := client.SemanticSearch.Search(ctx, components.SemanticSearchRequest{
		Query:   "What are some latest news about the stock market?",
		Filters: &components.Filters{Apps: []string{connectorID}},
	})
	if err != nil {
		log.Fatalf("search: %v", err)
	}

	for i, searchResult := range res.SemanticSearchExecuteResponse.SearchResponse.SearchResults {
		name, _ := searchResult.Metadata.RecordName.GetOrZero()
		id, _ := searchResult.Metadata.RecordID.GetOrZero()
		chunk, _ := searchResult.Content.GetOrZero()
		fmt.Printf("─── Result %d ──────────────────────────────────────────────\n", i+1)
		fmt.Printf("  Record:  %s\n", name)
		fmt.Printf("  ID:      %s\n", id)
		fmt.Printf("  Chunk:   %s\n\n", chunk)
	}
}
