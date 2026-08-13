// Get a knowledge base by ID.
//
// Usage (from examples/):
//
//	go run ./KB/kb/get_knowledge_base <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./KB/kb/get_knowledge_base <path-to-.env>")
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

	ctx := context.Background()

	created, err := sdk.KnowledgeBase.CreateKnowledgeBase(ctx, operations.CreateKnowledgeBaseRequest{
		KbName: "Internal documents",
	})
	if err != nil {
		log.Fatalf("create knowledge base: %v", err)
	}
	kbID := created.GetKnowledgeBaseCreateResponse().GetID()

	res, err := sdk.KnowledgeBase.GetKnowledgeBase(ctx, kbID)
	if err != nil {
		log.Fatalf("get knowledge base: %v", err)
	}

	printJSON(res.GetKnowledgeBaseByID)
}

func printJSON(v any) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(data))
}
