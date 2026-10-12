// Command internal_search streams one answer from the assistant's internal
// search path.
//
// This example decodes the AG-UI stream longhand, against the generated event
// constants and with the payload structs spelled out, so the wire protocol is
// visible. Every other conversation example uses the shared decoder in
// enterprise_search/agui instead.
//
// Usage: go run . <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"

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

	query := "Who moved the cheese?"

	res, err := client.Conversations.StreamChat(context.Background(), components.ConversationStreamRequest{
		Query:    query,
		ChatMode: components.ConversationStreamRequestChatModeInternalSearch,
	})
	if err != nil {
		log.Fatalf("conversation: %v", err)
	}
	if res == nil || res.ConversationStreamSSEEvent == nil {
		log.Fatal("no SSE stream returned")
	}
	stream := res.ConversationStreamSSEEvent
	defer stream.Close()

	fmt.Printf("You: %s\n\nBot: ", query)

	var conversationID, streamed, answer string

	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}

		switch *ev.Event {
		case components.ConversationStreamSSEEventEventCustom:
			var f struct {
				Name  string `json:"name"`
				Value struct {
					ConversationID string `json:"conversationId"`
				} `json:"value"`
			}
			if err := json.Unmarshal([]byte(*ev.Data), &f); err != nil {
				log.Fatalf("decode CUSTOM: %v", err)
			}
			// CUSTOM/conversation_created replaced the legacy `connected` event.
			// Other frames (tool_unavailable, ask_user_question) share this
			// channel, so matching on the name is required.
			if f.Name == "conversation_created" {
				conversationID = f.Value.ConversationID
			}

		case components.ConversationStreamSSEEventEventTextMessageContent:
			// The token field is `delta`. It is a string here but a JSON-Patch
			// array on STATE_DELTA, so the two cannot share one envelope struct.
			var f struct {
				Delta string `json:"delta"`
			}
			if err := json.Unmarshal([]byte(*ev.Data), &f); err != nil {
				log.Fatalf("decode TEXT_MESSAGE_CONTENT: %v", err)
			}
			streamed += f.Delta
			fmt.Print(f.Delta)

		case components.ConversationStreamSSEEventEventRunFinished:
			var f struct {
				ParentRunID *string `json:"parentRunId"`
				Result      struct {
					Conversation struct {
						ID       string `json:"_id"`
						Messages []struct {
							MessageType string `json:"messageType"`
							Content     string `json:"content"`
						} `json:"messages"`
					} `json:"conversation"`
				} `json:"result"`
			}
			if err := json.Unmarshal([]byte(*ev.Data), &f); err != nil {
				log.Fatalf("decode RUN_FINISHED: %v", err)
			}
			// Sub-agents emit their own RUN_FINISHED, which the gateway forwards
			// verbatim. Those carry parentRunId and no result, and arrive before
			// the terminal frame — treating one as terminal truncates the answer.
			if f.ParentRunID != nil {
				continue
			}
			if f.Result.Conversation.ID != "" {
				conversationID = f.Result.Conversation.ID
			}
			msgs := f.Result.Conversation.Messages
			for i := len(msgs) - 1; i >= 0; i-- {
				if msgs[i].MessageType == "bot_response" && msgs[i].Content != "" {
					answer = msgs[i].Content
					break
				}
			}
			report(conversationID, streamed, answer)
			return

		case components.ConversationStreamSSEEventEventRunError:
			var f struct {
				ParentRunID *string `json:"parentRunId"`
				Message     string  `json:"message"`
				Code        string  `json:"code"`
			}
			if err := json.Unmarshal([]byte(*ev.Data), &f); err != nil {
				log.Fatalf("decode RUN_ERROR: %v", err)
			}
			// Likewise, a child RUN_ERROR is one sub-agent failing, not the stream.
			if f.ParentRunID != nil {
				continue
			}
			log.Fatalf("stream error: %s (%s)", f.Message, f.Code)
		}
	}

	if err := stream.Err(); err != nil {
		log.Fatalf("stream: %v", err)
	}
	log.Fatal("stream ended without RUN_FINISHED")
}

func report(conversationID, streamed, answer string) {
	fmt.Println()
	// TEXT_MESSAGE_CONTENT carries raw append-only tokens, while citation
	// references are rewritten as the answer is finalized. The persisted message
	// is authoritative, so it can differ from what was streamed.
	if answer != "" && answer != streamed {
		fmt.Printf("\nFinal answer:\n%s\n", answer)
	}
	fmt.Printf("\nconversation: %s\n", conversationID)
}
