// Command knowledgebase streams one answer scoped to a single knowledge base,
// using Filters.Kb.
//
// Usage: go run . <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/agui"
	"enterprise_search/auth"
)

const defaultKnowledgeBaseName = "SDK-test"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run . <path-to-.env>")
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

	knowledgeBaseName := os.Getenv("PIPESHUB_KB_NAME")
	if knowledgeBaseName == "" {
		knowledgeBaseName = defaultKnowledgeBaseName
	}

	ctx := context.Background()

	kbsRes, err := sdk.KnowledgeBase.ListKnowledgeBases(ctx, operations.ListKnowledgeBasesRequest{
		Search: &knowledgeBaseName,
	})
	if err != nil {
		log.Fatalf("list knowledge bases: %v", err)
	}
	var kbID string
	for _, kb := range kbsRes.GetAllKnowledgeBaseResponseSchema.GetKnowledgeBases() {
		if kb.Name == knowledgeBaseName {
			kbID = kb.ID
			break
		}
	}
	if kbID == "" {
		log.Fatalf("knowledge base %q not found", knowledgeBaseName)
	}

	query := "Who moved the cheese?"

	res, err := sdk.Conversations.StreamChat(ctx, components.ConversationStreamRequest{
		Query:    query,
		ChatMode: components.ConversationStreamRequestChatModeInternalSearch,
		Filters:  &components.Filters{Kb: []string{kbID}},
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

	c := agui.Collector{Echo: os.Stdout}
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		done, err := c.Handle(string(*ev.Event), *ev.Data)
		if err != nil {
			log.Fatalf("stream: %v", err)
		}
		if done {
			break
		}
	}
	if err := stream.Err(); err != nil {
		log.Fatalf("stream: %v", err)
	}
	if !c.Done {
		log.Fatal("stream ended without RUN_FINISHED")
	}

	fmt.Println()
	// Citation references are rewritten as the answer is finalized, so the
	// persisted message can differ from the streamed tokens.
	if answer := c.Answer(); answer != c.Streamed {
		fmt.Printf("\nFinal answer:\n%s\n", answer)
	}
	fmt.Printf("\nconversation: %s\n", c.ConversationID)
}
