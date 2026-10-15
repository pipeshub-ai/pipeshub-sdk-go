// Create a knowledge base, upload two files, delete both records so they move
// to the trash, bring them back in one call with RestoreRecords, then delete
// the knowledge base.
//
// Requires the workspace Labs setting "Move Deleted Records to the Trash"
// (ENABLE_SOFT_DELETE). With it off, deleted records cannot be restored.
//
// Uploads the file at PIPESHUB_UPLOAD_PATH (or a small built-in text file
// when that variable is empty) plus a generated text note.
//
// Usage (from examples/):
//
//	go run ./KB/trash/restore_records <path-to-.env>
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
		log.Fatal("usage: go run ./KB/trash/restore_records <path-to-.env>")
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
	note := operations.UploadRecordsFile{
		FileName: "go-sdk-restore-note.txt",
		Content:  []byte(fmt.Sprintf("Restore example note written at %s.\n", time.Now().Format(time.RFC3339))),
	}

	kbID, err := createKnowledgeBase(ctx, sdk)
	if err != nil {
		log.Fatal(err)
	}
	ids, err := uploadFiles(ctx, sdk, kbID, []operations.UploadRecordsFile{file, note})
	if err != nil {
		log.Fatal(err)
	}
	for _, id := range ids {
		if _, err := sdk.KnowledgeBase.DeleteRecord(ctx, id); err != nil {
			log.Fatalf("delete record: %v", err)
		}
	}
	fmt.Printf("Uploaded and deleted %d records\n", len(ids))

	res, err := sdk.KnowledgeBase.RestoreRecords(ctx, components.RestoreRecordsRequestBody{
		RecordIds: ids,
	})
	if err != nil {
		log.Fatalf("restore records: %v", err)
	}
	if res.RestoreRecordsResponseSchema == nil {
		log.Fatal("restore records: empty response")
	}
	body := res.RestoreRecordsResponseSchema

	// Each id succeeds or fails on its own; check Results, not just Success.
	fmt.Printf("\nRestored %d, failed %d\n", body.RestoredCount, body.FailedCount)
	for _, r := range body.Results {
		if r.Success {
			fmt.Printf("  ok      %s\n", r.RecordID)
			continue
		}
		fmt.Printf("  failed  %s  %s %s\n", r.RecordID, ptr(r.Reason), ptr(r.Message))
	}

	if _, err := sdk.KnowledgeBase.DeleteKnowledgeBase(ctx, kbID); err != nil {
		log.Fatalf("delete knowledge base: %v", err)
	}
	fmt.Println("\nDeleted the knowledge base.")
}

func ptr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
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
