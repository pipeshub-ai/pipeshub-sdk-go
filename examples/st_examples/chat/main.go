// Command chat runs a two-turn conversation: StreamChat to open it, then
// AddMessageStream to follow up on the same conversation id.
//
// Usage: go run . <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/agui"
	"enterprise_search/auth"
)

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

	ctx := context.Background()

	limit := int64(100)
	kbsRes, err := sdk.KnowledgeBase.ListKnowledgeBases(ctx, operations.ListKnowledgeBasesRequest{
		Limit: &limit,
	})
	if err != nil {
		log.Fatalf("list knowledge bases: %v", err)
	}
	items := kbsRes.GetAllKnowledgeBaseResponseSchema.GetKnowledgeBases()
	kbIDs := make([]string, 0, len(items))
	for _, kb := range items {
		kbIDs = append(kbIDs, kb.ID)
	}
	if len(kbIDs) == 0 {
		log.Fatal("no knowledge bases found")
	}

	filters := &components.Filters{Kb: kbIDs}

	firstQuery := "Every year Asana performs what exercise?"
	convID, err := askFirst(ctx, sdk, firstQuery, filters)
	if err != nil {
		log.Fatalf("first turn: %v", err)
	}

	followUp := "Can you give me more details on that?"
	if err := askFollowUp(ctx, sdk, convID, followUp, filters); err != nil {
		log.Fatalf("follow-up turn: %v", err)
	}
}

func askFirst(ctx context.Context, sdk *pipeshub.Pipeshub, query string, filters *components.Filters) (string, error) {
	res, err := sdk.Conversations.StreamChat(ctx, components.ConversationStreamRequest{
		Query:    query,
		ChatMode: components.ConversationStreamRequestChatModeInternalSearch,
		Filters:  filters,
	})
	if err != nil {
		return "", fmt.Errorf("stream chat: %w", err)
	}
	if res == nil || res.ConversationStreamSSEEvent == nil {
		return "", fmt.Errorf("no SSE stream returned")
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
			return "", err
		}
		if done {
			break
		}
	}
	if err := stream.Err(); err != nil {
		return "", fmt.Errorf("stream: %w", err)
	}
	if !c.Done {
		return "", fmt.Errorf("stream ended without RUN_FINISHED")
	}
	fmt.Println()

	if c.ConversationID == "" {
		return "", fmt.Errorf("stream did not report a conversation id")
	}
	return c.ConversationID, nil
}

func askFollowUp(ctx context.Context, sdk *pipeshub.Pipeshub, convID, query string, filters *components.Filters) error {
	// The follow-up carries its own chatMode; it is required on this route and
	// should match the mode the conversation was opened with.
	res, err := sdk.Conversations.AddMessageStream(ctx, convID, components.ConversationMessageStreamRequest{
		Query:    query,
		ChatMode: components.ConversationMessageStreamRequestChatModeInternalSearch,
		Filters:  filters,
	})
	if err != nil {
		return fmt.Errorf("add message stream: %w", err)
	}
	if res == nil || res.ConversationMessageStreamSSEEvent == nil {
		return fmt.Errorf("no SSE stream returned")
	}
	stream := res.ConversationMessageStreamSSEEvent
	defer stream.Close()

	fmt.Printf("\nYou: %s\n\nBot: ", query)

	c := agui.Collector{Echo: os.Stdout}
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		done, err := c.Handle(string(*ev.Event), *ev.Data)
		if err != nil {
			return err
		}
		if done {
			break
		}
	}
	if err := stream.Err(); err != nil {
		return fmt.Errorf("stream: %w", err)
	}
	if !c.Done {
		return fmt.Errorf("stream ended without RUN_FINISHED")
	}
	fmt.Println()

	return nil
}
