// Start a StreamChat with a client-generated runId, cancel it with
// CancelConversationStream as soon as the first answer token arrives, then
// read the stream to its end to see how the run stopped.
//
// Usage (from examples/):
//
//	go run ./conversation/cancel_conversation_stream <path-to-.env>
package main

import (
	"context"
	"crypto/rand"
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

// A long answer, so there is still something to cancel after the first token.
const query = "Write a detailed 500-word overview of every news story in my workspace."

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/cancel_conversation_stream <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// The runId is a client-generated UUID sent with the stream request;
	// without one the run cannot be cancelled.
	runID := newUUID()

	res, err := sdk.Conversations.StreamChat(ctx, components.ConversationStreamRequest{
		Query:    query,
		ChatMode: components.ConversationStreamRequestChatModeInternalSearch,
		RunID:    &runID,
	})
	if err != nil {
		log.Fatalf("stream chat: %v", err)
	}
	if res.ConversationStreamSSEEvent == nil {
		log.Fatal("stream chat: no SSE stream returned")
	}
	stream := res.ConversationStreamSSEEvent
	defer stream.Close()

	var convID string
	cancelled := false
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		// data is a JSON-encoded string whose shape depends on the event.
		data := []byte(*ev.Data)

		switch *ev.Event {
		case components.ConversationStreamSSEEventEventCustom:
			// CUSTOM/conversation_created carries the id before any answer token.
			// Other CUSTOM frames carry unrelated shapes in value, so keep it raw
			// until the name is known.
			var f struct {
				Name  string          `json:"name"`
				Value json.RawMessage `json:"value"`
			}
			if err := json.Unmarshal(data, &f); err != nil {
				log.Fatalf("decode CUSTOM: %v", err)
			}
			if f.Name != "conversation_created" {
				continue
			}
			var v struct {
				ConversationID string `json:"conversationId"`
			}
			if err := json.Unmarshal(f.Value, &v); err != nil {
				log.Fatalf("decode CUSTOM/conversation_created: %v", err)
			}
			convID = v.ConversationID
			fmt.Printf("Conversation: %s\n", convID)

		case components.ConversationStreamSSEEventEventTextMessageContent:
			if cancelled {
				continue
			}
			var f struct {
				Delta string `json:"delta"`
			}
			if err := json.Unmarshal(data, &f); err != nil {
				log.Fatalf("decode TEXT_MESSAGE_CONTENT: %v", err)
			}
			fmt.Printf("First token:  %q\n", f.Delta)
			if convID == "" {
				log.Fatal("first token arrived before the conversation id")
			}
			cancelRun(ctx, sdk, convID, runID)
			cancelled = true

		case components.ConversationStreamSSEEventEventRunFinished, components.ConversationStreamSSEEventEventRunError:
			var f struct {
				ParentRunID *string `json:"parentRunId"`
				Message     string  `json:"message"`
				Result      struct {
					Conversation struct {
						Status string `json:"status"`
					} `json:"conversation"`
				} `json:"result"`
			}
			if err := json.Unmarshal(data, &f); err != nil {
				log.Fatalf("decode %s: %v", *ev.Event, err)
			}
			// Sub-agent runs carry parentRunId; only the root run is terminal.
			if f.ParentRunID != nil {
				continue
			}
			fmt.Printf("Terminal:     %s\n", *ev.Event)
			if f.Result.Conversation.Status != "" {
				fmt.Printf("Status:       %s\n", f.Result.Conversation.Status)
			}
			if f.Message != "" {
				fmt.Printf("Message:      %s\n", f.Message)
			}
		}
	}
	if err := stream.Err(); err != nil {
		log.Fatalf("stream: %v", err)
	}
	if !cancelled {
		fmt.Println("The answer finished before any token arrived; nothing was cancelled.")
	}

	if convID != "" {
		if _, err := sdk.Conversations.DeleteConversationByID(ctx, convID); err != nil {
			log.Fatalf("delete conversation: %v", err)
		}
		fmt.Println("Deleted the conversation.")
	}
}

// cancelRun asks the server to stop the run. {cancelled: false} is not an
// error: it means the run had already finished.
func cancelRun(ctx context.Context, sdk *pipeshub.Pipeshub, convID, runID string) {
	res, err := sdk.Conversations.CancelConversationStream(ctx, convID, operations.CancelConversationStreamRequestBody{
		RunID: runID,
	})
	if err != nil {
		log.Fatalf("cancel conversation stream: %v", err)
	}
	ok := res.Object != nil && res.Object.Cancelled != nil && *res.Object.Cancelled
	fmt.Printf("Cancel:       cancelled=%v\n", ok)
}

// newUUID returns a random (version 4) UUID.
func newUUID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		log.Fatalf("generate run id: %v", err)
	}
	b[6] = b[6]&0x0f | 0x40 // version 4
	b[8] = b[8]&0x3f | 0x80 // RFC 4122 variant
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
