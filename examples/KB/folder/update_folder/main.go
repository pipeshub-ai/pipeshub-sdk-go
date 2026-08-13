// Rename a folder in a knowledge base.
//
// Usage (from examples/):
//
//	go run ./KB/folder/update_folder <path-to-.env>
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
		log.Fatal("usage: go run ./KB/folder/update_folder <path-to-.env>")
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

	folder, err := sdk.KnowledgeBase.CreateFolder(ctx, kbID, operations.CreateFolderRequestBody{
		FolderName: "Reports",
	}, nil)
	if err != nil {
		log.Fatalf("create folder: %v", err)
	}
	folderID := folder.GetFolderCreateResponseSchema().GetID()

	res, err := sdk.KnowledgeBase.UpdateFolder(ctx, kbID, folderID, operations.UpdateFolderRequestBody{
		FolderName: "Reports (updated)",
	})
	if err != nil {
		log.Fatalf("update folder: %v", err)
	}

	printJSON(res.FolderUpdateResponseSchema)
}

func printJSON(v any) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(data))
}
