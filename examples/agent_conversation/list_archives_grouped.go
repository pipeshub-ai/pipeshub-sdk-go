// Create two agent conversations (streaming), archive both, then list archived
// conversations grouped by agent.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/list_archives_grouped.go <path-to-.env>
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

	"enterprise_search/auth"
)

const (
	agentKey    = "02a7d998-d21b-4015-aaf7-5cda765c1012"
	kbID        = "8747da12-4724-4a95-ac92-827b88d79647"
	connectorID = "aeab9ddc-fb9b-47c8-ad98-bd4744e19555"

	firstMessage  = "What are some latest tech news?"
	secondMessage = "Summarize recent AI industry news in three bullets."
)

type archivedConv struct {
	id    string
	title string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/list_archives_grouped.go <path-to-.env>")
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
	filters := &components.Filters{
		Kb:   []string{kbID},
		Apps: []string{connectorID},
	}

	var created []archivedConv

	for i, query := range []string{firstMessage, secondMessage} {
		fmt.Printf("Creating conversation %d (waiting for response...)...\n", i+1)
		convID, title, err := createAgentConversation(ctx, sdk, query, filters)
		if err != nil {
			log.Fatalf("create conversation %d: %v", i+1, err)
		}
		if title == "" {
			title = query
		}
		if err := archiveAgentConversation(ctx, sdk, convID); err != nil {
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
	if res.AgentArchivedGroupsResponse == nil {
		log.Fatal("no grouped archive response returned")
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

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, query string, filters *components.Filters) (convID, title string, err error) {
	chatMode := components.AgentStreamCreateConversationRequestChatModeAuto

	res, err := sdk.Agents.StreamAgentConversation(ctx, agentKey, components.AgentStreamCreateConversationRequest{
		Query:    query,
		Filters:  filters,
		ChatMode: chatMode.ToPointer(),
	})
	if err != nil {
		return "", "", fmt.Errorf("stream agent conversation: %w", err)
	}
	if res.AgentStreamSSEEvent == nil {
		return "", "", fmt.Errorf("no SSE stream returned")
	}
	stream := res.AgentStreamSSEEvent
	defer stream.Close()

	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.AgentStreamSSEEventEventComplete:
			id, t, err := decodeConversationComplete(*ev.Data)
			if err != nil {
				return "", "", err
			}
			if id == "" {
				return "", "", fmt.Errorf("complete event missing conversation id")
			}
			return id, t, nil
		case components.AgentStreamSSEEventEventError:
			return "", "", fmt.Errorf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		return "", "", fmt.Errorf("stream: %w", err)
	}
	return "", "", fmt.Errorf("stream ended without complete event")
}

func archiveAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, convID string) error {
	res, err := sdk.Agents.ArchiveAgentConversation(ctx, agentKey, convID)
	if err != nil {
		return fmt.Errorf("archive agent conversation: %w", err)
	}
	if res.AgentConversationArchiveResponse == nil {
		return fmt.Errorf("no archive response returned")
	}
	return nil
}

func decodeConversationComplete(data string) (convID, title string, err error) {
	var payload struct {
		Conversation struct {
			ID       string  `json:"_id"`
			Title    *string `json:"title"`
			Messages []struct {
				MessageType string `json:"messageType"`
				Content     string `json:"content"`
			} `json:"messages"`
		} `json:"conversation"`
	}
	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		return "", "", fmt.Errorf("decode complete: %w", err)
	}
	convID = payload.Conversation.ID
	if payload.Conversation.Title != nil {
		title = *payload.Conversation.Title
	}
	for _, m := range payload.Conversation.Messages {
		if title == "" && m.MessageType == "user_query" {
			title = m.Content
		}
	}
	return convID, title, nil
}
