// Create an agent conversation (streaming), then archive and unarchive it.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/archive_unarchive <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"

	"enterprise_search/agui"
	"enterprise_search/auth"
)

// First user message when creating the conversation.
const firstMessage = "Who moved the cheese?"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/archive_unarchive <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	agentKey := os.Getenv("PIPESHUB_AGENT_KEY")
	if agentKey == "" {
		log.Fatal("PIPESHUB_AGENT_KEY is required")
	}

	sdk, err := auth.NewClient(
		os.Getenv("PIPESHUB_TEST_USER_EMAIL"),
		os.Getenv("PIPESHUB_TEST_USER_PASSWORD"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	convID, title, err := createAgentConversation(ctx, sdk, agentKey, firstMessage, envFilters())
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	if title == "" {
		title = firstMessage
	}

	fmt.Printf("Created conversation: %s\n", convID)
	fmt.Printf("Title: %q\n", title)

	if err := archiveAgentConversation(ctx, sdk, agentKey, convID); err != nil {
		log.Fatalf("archive conversation: %v", err)
	}
	if err := unarchiveAgentConversation(ctx, sdk, agentKey, convID); err != nil {
		log.Fatalf("unarchive conversation: %v", err)
	}
}

// envFilters builds knowledge filters from KB_ID and CONNECTOR_ID.
//
// It returns nil when neither is set, so the agent falls back to its stored
// knowledge config. An empty-but-non-nil Filters would instead be sent as
// {"apps":[],"kb":[]}, which means "no knowledge sources at all".
func envFilters() *components.Filters {
	kbID, connectorID := os.Getenv("KB_ID"), os.Getenv("CONNECTOR_ID")
	if kbID == "" && connectorID == "" {
		return nil
	}
	filters := &components.Filters{}
	if kbID != "" {
		filters.Kb = []string{kbID}
	}
	if connectorID != "" {
		filters.Apps = []string{connectorID}
	}
	return filters
}

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, query string, filters *components.Filters) (convID, title string, err error) {
	res, err := sdk.Agents.StreamAgentConversation(ctx, agentKey, components.AgentStreamCreateConversationRequest{
		Query:   query,
		Filters: filters,
		// Scoped agent conversations accept only "quick".
		ChatMode: components.AgentStreamCreateConversationRequestChatModeQuick,
	})
	if err != nil {
		return "", "", fmt.Errorf("stream agent conversation: %w", err)
	}
	if res == nil || res.AgentStreamSSEEvent == nil {
		return "", "", fmt.Errorf("no SSE stream returned")
	}
	stream := res.AgentStreamSSEEvent
	defer stream.Close()

	var c agui.Collector
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		done, err := c.Handle(string(*ev.Event), *ev.Data)
		if err != nil {
			return "", "", err
		}
		if done {
			break
		}
	}
	if err := stream.Err(); err != nil {
		return "", "", fmt.Errorf("stream: %w", err)
	}
	if !c.Done {
		return "", "", fmt.Errorf("stream ended without RUN_FINISHED")
	}
	if c.ConversationID == "" {
		return "", "", fmt.Errorf("stream did not report a conversation id")
	}
	return c.ConversationID, c.Title, nil
}

func archiveAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, convID string) error {
	res, err := sdk.Agents.ArchiveAgentConversation(ctx, agentKey, convID)
	if err != nil {
		return fmt.Errorf("archive agent conversation: %w", err)
	}
	at := res.AgentConversationArchiveResponse.GetArchivedAt()
	if at.IsZero() {
		fmt.Println("Archived (by you): conversation is now in archives")
		return nil
	}
	fmt.Printf("Archived (by you at %s): conversation is now in archives\n", at.Format(time.RFC1123))
	return nil
}

func unarchiveAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, convID string) error {
	res, err := sdk.Agents.UnarchiveAgentConversation(ctx, agentKey, convID)
	if err != nil {
		return fmt.Errorf("unarchive agent conversation: %w", err)
	}
	at := res.AgentConversationUnarchiveResponse.GetUnarchivedAt()
	if at.IsZero() {
		fmt.Println("Unarchived: conversation is back in your active list")
		return nil
	}
	fmt.Printf("Unarchived (at %s): conversation is back in your active list\n", at.Format(time.RFC1123))
	return nil
}
