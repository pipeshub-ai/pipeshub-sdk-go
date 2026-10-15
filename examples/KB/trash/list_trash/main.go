// Create a knowledge base, upload a file, delete the record so it moves to the
// trash, list the trash with ListTrash, then delete the knowledge base.
//
// Requires the workspace Labs setting "Move Deleted Records to the Trash"
// (ENABLE_SOFT_DELETE). With it off, deleted records are removed for good and
// the trash stays empty.
//
// Uploads the file at PIPESHUB_UPLOAD_PATH, or a small built-in text file
// when that variable is empty.
//
// Usage (from examples/):
//
//	go run ./KB/trash/list_trash <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./KB/trash/list_trash <path-to-.env>")
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
	if _, err := sdk.KnowledgeBase.DeleteRecord(ctx, ids[0]); err != nil {
		log.Fatalf("delete record: %v", err)
	}
	fmt.Printf("Uploaded and deleted record %s\n\n", ids[0])

	res, err := sdk.KnowledgeBase.ListTrash(ctx, kbID, pipeshub.Pointer(int64(1)), pipeshub.Pointer(int64(25)))
	if err != nil {
		log.Fatalf("list trash: %v", err)
	}
	if res.ListTrashResponseSchema == nil {
		log.Fatal("list trash: empty response")
	}
	body := res.ListTrashResponseSchema

	fmt.Printf("Trash holds %d item(s):\n", body.Pagination.TotalCount)
	for _, item := range body.Items {
		name, _ := item.Name.GetOrZero()
		by := ""
		if d, ok := item.DeletedBy.GetOrZero(); ok {
			by, _ = d.Email.GetOrZero()
		}
		fmt.Printf("  %s  %s\n", item.ID, name)
		fmt.Printf("      deleted %s by %s\n", time.UnixMilli(item.DeletedAtTimestamp).Format(time.DateTime), by)
	}
	if len(body.Items) == 0 {
		fmt.Println("  (empty) Is \"Move Deleted Records to the Trash\" turned on in Labs?")
	}

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
			if err := json.Unmarshal([]byte(*ev.Data), &payload); err != nil {
				return nil, fmt.Errorf("decode %s: %w", *ev.Event, err)
			}
			if payload.RecordID != "" {
				ids = append(ids, payload.RecordID)
			}
		case components.UploadStreamSSEEventEventFileFailed, components.UploadStreamSSEEventEventError:
			return nil, fmt.Errorf("upload records: %s %s", *ev.Event, *ev.Data)
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
