// Create a knowledge base, upload a file, fetch the new record with
// GetRecordByID, then delete the knowledge base.
//
// Uploads the file at PIPESHUB_UPLOAD_PATH, or a small built-in text file
// when that variable is empty.
//
// Usage (from examples/):
//
//	go run ./KB/records/get_record_by_id <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"enterprise_search/agui"
	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./KB/records/get_record_by_id <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	file, err := sampleFile()
	if err != nil {
		log.Fatal(err)
	}
	kbID, err := createKnowledgeBase(ctx, sdk)
	if err != nil {
		log.Fatal(err)
	}
	ids, err := uploadFiles(ctx, sdk, kbID, []operations.UploadRecordsFile{file})
	if err != nil {
		log.Fatal(err)
	}
	recordID := ids[0]

	// convertTo is optional; pass nil to get the record as stored.
	res, err := sdk.KnowledgeBase.GetRecordByID(ctx, recordID, nil)
	if err != nil {
		log.Fatalf("get record: %v", err)
	}
	if res.GetRecordByIDResponseSchema == nil {
		log.Fatal("get record: empty response")
	}
	body := res.GetRecordByIDResponseSchema
	r := body.Record

	fmt.Printf("Record:          %s\n", r.ID)
	fmt.Printf("Name:            %s\n", r.RecordName)
	fmt.Printf("Type:            %s\n", r.RecordType)
	fmt.Printf("Origin:          %s (%s)\n", r.Origin, r.ConnectorName)
	fmt.Printf("Indexing status: %s\n", r.IndexingStatus)
	if r.MimeType != nil && r.SizeInBytes != nil {
		fmt.Printf("File:            %s, %d bytes\n", *r.MimeType, *r.SizeInBytes)
	}
	fmt.Printf("Created:         %s\n", time.UnixMilli(r.CreatedAtTimestamp).Format(time.RFC1123))
	if body.KnowledgeBase != nil {
		fmt.Printf("Knowledge base:  %s (%s)\n", body.KnowledgeBase.Name, body.KnowledgeBase.ID)
	}
	fmt.Printf("Permissions:     %d\n", len(body.Permissions))

	if _, err := sdk.KnowledgeBase.DeleteKnowledgeBase(ctx, kbID); err != nil {
		log.Fatalf("delete knowledge base: %v", err)
	}
	fmt.Println("\nDeleted the knowledge base.")
}

// uploadFiles uploads files into the KB root and returns the new record ids,
// read from each file:succeeded event.
func uploadFiles(ctx context.Context, sdk *pipeshub.Pipeshub, kbID string, files []operations.UploadRecordsFile) ([]string, error) {
	res, err := sdk.KnowledgeBase.UploadRecords(ctx, kbID, operations.UploadRecordsRequestBody{Files: files}, nil)
	if err != nil {
		return nil, fmt.Errorf("upload records: %w", err)
	}
	stream := res.UploadStreamSSEEvent
	if stream == nil {
		return nil, fmt.Errorf("upload records: no SSE stream returned")
	}
	defer stream.Close()

	var ids []string
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.UploadStreamSSEEventEventFileSucceeded:
			// data is a JSON-encoded string; decode just the field we need.
			var payload struct {
				RecordID string `json:"recordId"`
			}
			if err := json.Unmarshal([]byte(agui.Raw(ev.Data)), &payload); err != nil {
				return nil, fmt.Errorf("decode %s: %w", *ev.Event, err)
			}
			if payload.RecordID != "" {
				ids = append(ids, payload.RecordID)
			}
		case components.UploadStreamSSEEventEventFileFailed, components.UploadStreamSSEEventEventError:
			return nil, fmt.Errorf("upload records: %s %s", *ev.Event, agui.Raw(ev.Data))
		}
	}
	if err := stream.Err(); err != nil {
		return nil, fmt.Errorf("upload stream: %w", err)
	}
	if len(ids) != len(files) {
		return nil, fmt.Errorf("upload records: %d of %d files reported a record id", len(ids), len(files))
	}
	return ids, nil
}

// sampleFile is the file at PIPESHUB_UPLOAD_PATH, or a small built-in text
// file when that variable is empty.
func sampleFile() (operations.UploadRecordsFile, error) {
	p := os.Getenv("PIPESHUB_UPLOAD_PATH")
	if p == "" {
		return operations.UploadRecordsFile{FileName: "sdk-example.txt", Content: []byte("Hello from the PipesHub Go SDK.\n")}, nil
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return operations.UploadRecordsFile{}, fmt.Errorf("read upload file: %w", err)
	}
	return operations.UploadRecordsFile{FileName: filepath.Base(p), Content: b}, nil
}

// createKnowledgeBase creates a fresh knowledge base and returns its id.
func createKnowledgeBase(ctx context.Context, sdk *pipeshub.Pipeshub) (string, error) {
	res, err := sdk.KnowledgeBase.CreateKnowledgeBase(ctx, operations.CreateKnowledgeBaseRequest{
		KbName: fmt.Sprintf("Go SDK example KB %d", time.Now().Unix()),
	})
	if err != nil {
		return "", fmt.Errorf("create knowledge base: %w", err)
	}
	id := res.GetKnowledgeBaseCreateResponse().GetID()
	if id == "" {
		return "", fmt.Errorf("create knowledge base: response has no id")
	}
	return id, nil
}
