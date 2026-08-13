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

	// Search across every accessible knowledge base, following pagination so
	// KBs beyond the first page are included too.
	limit := int64(100)
	var kbIDs []string
	for page := int64(1); ; page++ {
		kbsRes, err := client.KnowledgeBase.ListKnowledgeBases(ctx, operations.ListKnowledgeBasesRequest{
			Page:  &page,
			Limit: &limit,
		})
		if err != nil {
			log.Fatalf("list knowledge bases (page %d): %v", page, err)
		}
		schema := kbsRes.GetAllKnowledgeBaseResponseSchema
		for _, kb := range schema.GetKnowledgeBases() {
			kbIDs = append(kbIDs, kb.ID)
		}
		if !schema.GetPagination().HasNext {
			break
		}
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

	if res == nil || res.SemanticSearchExecuteResponse == nil {
		log.Fatal("search: empty response")
	}

	results := res.SemanticSearchExecuteResponse.SearchResponse.SearchResults
	if len(results) == 0 {
		log.Fatal("search: no results — is anything indexed in these knowledge bases?")
	}

	for i, searchResult := range results {
		// Metadata is optional on a hit; the generated getters are nil-safe,
		// direct field access is not.
		name, _ := searchResult.Metadata.GetRecordName().GetOrZero()
		id, _ := searchResult.Metadata.GetRecordID().GetOrZero()
		chunk, _ := searchResult.Content.GetOrZero()
		fmt.Printf("─── Result %d ──────────────────────────────────────────────\n", i+1)
		fmt.Printf("  Record:  %s\n", name)
		fmt.Printf("  ID:      %s\n", id)
		fmt.Printf("  Chunk:   %s\n\n", chunk)
	}
}
