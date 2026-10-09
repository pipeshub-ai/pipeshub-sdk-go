// Create an agent conversation (streaming), then update its title.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/update_conversation_title <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"

	"enterprise_search/agui"
	"enterprise_search/auth"
)

const (
	// First user message when creating the conversation.
	firstMessage = "Who moved the cheese?"

	// Title applied via UpdateAgentConversationTitle.
	newTitle = "SDK example: updated title"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/update_conversation_title <path-to-.env>")
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

	convID, oldTitle, err := createAgentConversation(ctx, sdk, agentKey, firstMessage, envFilters())
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	log.Printf("conversation id: %s", convID)
	if oldTitle == "" {
		oldTitle = firstMessage
	}

	updatedTitle, err := updateAgentConversationTitle(ctx, sdk, agentKey, convID, newTitle)
	if err != nil {
		log.Fatalf("update title: %v", err)
	}

	log.Printf("old title: %q", oldTitle)
	log.Printf("new title: %q", updatedTitle)
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

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, query string, filters *components.Filters) (convID, oldTitle string, err error) {
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

	fmt.Printf("You: %s\n\nBot: ", query)

	c := agui.Collector{Echo: os.Stdout}
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
	fmt.Println()

	if c.ConversationID == "" {
		return "", "", fmt.Errorf("stream did not report a conversation id")
	}
	return c.ConversationID, c.Title, nil
}

func updateAgentConversationTitle(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, convID, title string) (string, error) {
	res, err := sdk.Agents.UpdateAgentConversationTitle(ctx, agentKey, convID, components.ConversationTitleUpdateRequest{
		Title: title,
	})
	if err != nil {
		return "", fmt.Errorf("update agent conversation title: %w", err)
	}
	conv := res.AgentConversationTitleUpdateResponse.GetConversation()
	updated := (&conv).GetTitle()
	if updated == nil {
		return "", fmt.Errorf("response missing conversation title")
	}
	return *updated, nil
}
