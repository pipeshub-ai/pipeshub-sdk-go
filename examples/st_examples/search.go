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

	orgRes, err := client.Organizations.GetCurrentOrganization(ctx)
	if err != nil {
		log.Fatalf("get current organization: %v", err)
	}
	if orgRes == nil || orgRes.Organization == nil || orgRes.Organization.ID == nil || *orgRes.Organization.ID == "" {
		log.Fatal("get current organization: missing organization id")
	}
	parentID := "knowledgeBase_" + *orgRes.Organization.ID

	kbsRes, err := client.KnowledgeHub.GetKnowledgeHubChildNodes(ctx, operations.GetKnowledgeHubChildNodesRequest{
		ParentType: operations.ParentTypeApp,
		ParentID:   parentID,
	})
	if err != nil {
		log.Fatalf("list knowledge bases: %v", err)
	}
	if kbsRes == nil || kbsRes.KnowledgeHubNodesResponse == nil {
		log.Fatal("list knowledge bases: empty response")
	}
	items := kbsRes.KnowledgeHubNodesResponse.GetItems()
	kbIDs := make([]string, 0, len(items))
	for _, kb := range items {
		kbIDs = append(kbIDs, kb.ID)
	}
	if len(kbIDs) == 0 {
		log.Fatal("no knowledge bases found")
	}

	res, err := client.SemanticSearch.Search(ctx, components.SemanticSearchRequest{
		Query:   "What is SoundThinking?",
		Filters: &components.Filters{Kb: kbIDs},
	})
	if err != nil {
		log.Fatalf("search: %v", err)
	}

	if res == nil || res.SemanticSearchExecuteResponse == nil || res.SemanticSearchExecuteResponse.SearchResponse == nil {
		log.Fatal("search: empty response")
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
