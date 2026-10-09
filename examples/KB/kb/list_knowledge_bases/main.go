// List knowledge bases (capping the page size with Limit, max 100).
//
// Usage (from examples/):
//
//	go run ./KB/kb/list_knowledge_bases <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./KB/kb/list_knowledge_bases <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewClient(
		os.Getenv("PIPESHUB_TEST_USER_EMAIL"),
		os.Getenv("PIPESHUB_TEST_USER_PASSWORD"),
	)
	if err != nil {
		log.Fatal(err)
	}

	res, err := sdk.KnowledgeBase.ListKnowledgeBases(context.Background(), operations.ListKnowledgeBasesRequest{
		Limit: pipeshub.Pointer(int64(10)),
	})
	if err != nil {
		log.Fatalf("list knowledge bases: %v", err)
	}

	printJSON(res.GetAllKnowledgeBaseResponseSchema)
}

func printJSON(v any) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(data))
}
