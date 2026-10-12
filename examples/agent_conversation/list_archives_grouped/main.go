// Create two agent conversations (streaming), archive both, then list archived
// conversations grouped by agent.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/list_archives_grouped <path-to-.env>
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

	// Second conversation uses a different opening query.
	secondMessage = "Can you give me more details on that?"
)

type archivedConv struct {
	id    string
	title string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/list_archives_grouped <path-to-.env>")
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
	filters := envFilters()

	var created []archivedConv

	for i, query := range []string{firstMessage, secondMessage} {
		fmt.Printf("Creating conversation %d (waiting for response...)...\n", i+1)
		convID, title, err := createAgentConversation(ctx, sdk, agentKey, query, filters)
		if err != nil {
			log.Fatalf("create conversation %d: %v", i+1, err)
		}
		if title == "" {
			title = query
		}
		if err := archiveAgentConversation(ctx, sdk, agentKey, convID); err != nil {
			log.Fatalf("archive conversation %d: %v", i+1, err)
		}
		created = append(created, archivedConv{id: convID, title: title})
	}

	fmt.Println("\nCreated and archived:")
	for i, c := range created {
		fmt.Printf("  %d. %s — %q\n", i+1, c.id, c.title)
	}

	res, err := sdk.Agents.ListAgentArchivedConversationsGrouped(ctx, pipeshub.Int64(1), pipeshub.Int64(20))
	if err != nil {
		log.Fatalf("list archived conversations grouped: %v", err)
	}
	groups := res.AgentArchivedGroupsResponse.GetGroups()
	var agentGroup *components.AgentArchivedConversationGroup
	for i := range groups {
		if groups[i].GetAgentKey() == agentKey {
			agentGroup = &groups[i]
			break
		}
	}

	fmt.Printf("\nArchived conversations for this agent (grouped list):\n")
	if agentGroup == nil {
		fmt.Println("  (no group found for this agentKey)")
		return
	}

	createdIDs := make(map[string]string, len(created))
	for _, c := range created {
		createdIDs[c.id] = c.title
	}

	matched := 0
	for _, conv := range agentGroup.GetConversations() {
		id := conv.GetID()
		if id == nil || *id == "" {
			continue
		}
		title := ""
		if t := conv.GetTitle(); t != nil {
			title = *t
		}
		if _, ok := createdIDs[*id]; ok {
			matched++
			fmt.Printf("  - %q (%s)\n", title, *id)
		}
	}

	if matched == 0 {
		fmt.Println("  (none of the conversations we created appear in this page of results)")
	} else if matched == len(created) {
		fmt.Printf("\nBoth conversations created in this run appear under agent %s.\n", agentKey)
	} else {
		fmt.Printf("\n%d of %d conversations created in this run appear under agent %s.\n", matched, len(created), agentKey)
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
	if _, err := sdk.Agents.ArchiveAgentConversation(ctx, agentKey, convID); err != nil {
		return fmt.Errorf("archive agent conversation: %w", err)
	}
	return nil
}
