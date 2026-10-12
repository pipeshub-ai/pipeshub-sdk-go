// Move a record from a source folder into a destination folder.
//
// Usage (from examples/):
//
//	go run ./KB/records/move_record <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./KB/records/move_record <path-to-.env>")
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

	source, err := sdk.KnowledgeBase.CreateFolder(ctx, kbID, operations.CreateFolderRequestBody{
		FolderName: "Source",
	}, nil)
	if err != nil {
		log.Fatalf("create source folder: %v", err)
	}
	sourceID := source.GetFolderCreateResponseSchema().GetID()

	recordID := uploadSampleRecord(ctx, sdk, kbID, &sourceID)

	destination, err := sdk.KnowledgeBase.CreateFolder(ctx, kbID, operations.CreateFolderRequestBody{
		FolderName: "Destination",
	}, nil)
	if err != nil {
		log.Fatalf("create destination folder: %v", err)
	}
	destinationID := destination.GetFolderCreateResponseSchema().GetID()

	res, err := sdk.KnowledgeBase.MoveRecord(ctx, kbID, recordID, components.KnowledgeBaseMoveRecordRequestBody{
		NewParentID: &destinationID,
	})
	if err != nil {
		log.Fatalf("move record: %v", err)
	}

	printJSON(res.KnowledgeBaseMoveRecordResponse)
}

// uploadSampleRecord uploads a small text file and returns the new record's ID.
func uploadSampleRecord(ctx context.Context, sdk *pipeshub.Pipeshub, kbID string, folderID *string) string {
	res, err := sdk.KnowledgeBase.UploadRecords(ctx, kbID, operations.UploadRecordsRequestBody{
		Files: []operations.UploadRecordsFile{
			{FileName: "sdk-example.txt", Content: []byte("Hello from the PipesHub Go SDK.\n")},
		},
	}, folderID)
	if err != nil {
		log.Fatalf("upload record: %v", err)
	}
	stream := res.UploadStreamSSEEvent
	if stream == nil {
		log.Fatal("upload record: no SSE stream returned")
	}
	defer stream.Close()

	recordID := ""
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil || *ev.Event != components.UploadStreamSSEEventEventFileSucceeded {
			continue
		}
		var payload struct {
			RecordID string `json:"recordId"`
			ID       string `json:"id"`
			Record   struct {
				ID string `json:"id"`
			} `json:"record"`
		}
		_ = json.Unmarshal([]byte(*ev.Data), &payload)
		for _, id := range []string{payload.RecordID, payload.ID, payload.Record.ID} {
			if id != "" {
				recordID = id
				break
			}
		}
	}
	if err := stream.Err(); err != nil {
		log.Fatalf("upload stream: %v", err)
	}
	if recordID == "" {
		log.Fatal("upload stream did not report a record ID")
	}
	return recordID
}

func printJSON(v any) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(data))
}
