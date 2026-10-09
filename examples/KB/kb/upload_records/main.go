// Upload a file into a knowledge base root (streaming the upload events).
//
// Uploads the file at PIPESHUB_UPLOAD_PATH, or a small built-in text file
// when that variable is empty.
//
// Usage (from examples/):
//
//	go run ./KB/kb/upload_records <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./KB/kb/upload_records <path-to-.env>")
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

	fileName := "sdk-example.txt"
	content := []byte("Hello from the PipesHub Go SDK.\n")
	if p := os.Getenv("PIPESHUB_UPLOAD_PATH"); p != "" {
		b, err := os.ReadFile(p)
		if err != nil {
			log.Fatalf("read upload file: %v", err)
		}
		fileName, content = filepath.Base(p), b
	}

	ctx := context.Background()

	created, err := sdk.KnowledgeBase.CreateKnowledgeBase(ctx, operations.CreateKnowledgeBaseRequest{
		KbName: "Internal documents",
	})
	if err != nil {
		log.Fatalf("create knowledge base: %v", err)
	}
	kbID := created.GetKnowledgeBaseCreateResponse().GetID()

	res, err := sdk.KnowledgeBase.UploadRecords(ctx, kbID, operations.UploadRecordsRequestBody{
		Files: []operations.UploadRecordsFile{
			{FileName: fileName, Content: content},
		},
	}, nil)
	if err != nil {
		log.Fatalf("upload records: %v", err)
	}

	stream := res.UploadStreamSSEEvent
	if stream == nil {
		log.Fatal("upload records: no SSE stream returned")
	}
	defer stream.Close()

	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil {
			continue
		}
		data := ""
		if ev.Data != nil {
			data = *ev.Data
		}
		fmt.Printf("%s  %s\n", *ev.Event, data)
	}
	if err := stream.Err(); err != nil {
		log.Fatalf("upload stream: %v", err)
	}
}
