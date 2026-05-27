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

	kbIDs, err := listKnowledgeBaseIDs(ctx, sdk)
	if err != nil {
		log.Fatal(err)
	}
	if len(kbIDs) == 0 {
		log.Fatal("no knowledge bases found")
	}

	chatMode := "internal_search"
	filters := &components.Filters{Kb: kbIDs}

	firstQuery := "Every year Asana performs what exercise?"
	convID, err := askFirst(ctx, sdk, firstQuery, chatMode, filters)
	if err != nil {
		log.Fatalf("first turn: %v", err)
	}

	followUp := "Can you give me more details on that?"
	if err := askFollowUp(ctx, sdk, convID, followUp, filters); err != nil {
		log.Fatalf("follow-up turn: %v", err)
	}
}

func askFirst(ctx context.Context, sdk *pipeshub.Pipeshub, query, chatMode string, filters *components.Filters) (string, error) {
	res, err := sdk.Conversations.StreamChat(ctx, components.CreateConversationRequest{
		Query:    query,
		ChatMode: &chatMode,
		Filters:  filters,
	})
	if err != nil {
		return "", fmt.Errorf("stream chat: %w", err)
	}
	if res == nil || res.AssistantStreamSSEEvent == nil {
		return "", fmt.Errorf("no SSE stream returned")
	}
	stream := res.AssistantStreamSSEEvent
	defer stream.Close()

	fmt.Printf("You: %s\n\nBot: ", query)

	var convID string
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.AssistantStreamSSEEventEventConnected:
			var payload struct {
				ConversationID string `json:"conversationId"`
			}
			if err := json.Unmarshal([]byte(*ev.Data), &payload); err == nil {
				convID = payload.ConversationID
			}
		case components.AssistantStreamSSEEventEventComplete:
			answer, id, err := decodeComplete(*ev.Data)
			if err != nil {
				return "", err
			}
			if convID == "" {
				convID = id
			}
			fmt.Println(answer)
			if convID == "" {
				return "", fmt.Errorf("complete event missing conversation id")
			}
			return convID, nil
		case components.AssistantStreamSSEEventEventError:
			return "", fmt.Errorf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		return "", fmt.Errorf("stream: %w", err)
	}
	return "", fmt.Errorf("stream ended without complete event")
}

func askFollowUp(ctx context.Context, sdk *pipeshub.Pipeshub, convID, query string, filters *components.Filters) error {
	res, err := sdk.Conversations.AddMessageStream(ctx, convID, components.AddMessageRequest{
		Query:   query,
		Filters: filters,
	})
	if err != nil {
		return fmt.Errorf("add message stream: %w", err)
	}
	if res == nil || res.AssistantMessageStreamSSEEvent == nil {
		return fmt.Errorf("no SSE stream returned")
	}
	stream := res.AssistantMessageStreamSSEEvent
	defer stream.Close()

	fmt.Printf("\nYou: %s\n\nBot: ", query)

	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.AssistantMessageStreamSSEEventEventComplete:
			answer, _, err := decodeComplete(*ev.Data)
			if err != nil {
				return err
			}
			fmt.Println(answer)
			return nil
		case components.AssistantMessageStreamSSEEventEventError:
			return fmt.Errorf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		return fmt.Errorf("stream: %w", err)
	}
	return fmt.Errorf("stream ended without complete event")
}

func decodeComplete(data string) (answer, convID string, err error) {
	var payload struct {
		Conversation struct {
			ID       string `json:"_id"`
			Messages []struct {
				MessageType string `json:"messageType"`
				Content     string `json:"content"`
			} `json:"messages"`
		} `json:"conversation"`
	}
	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		return "", "", fmt.Errorf("decode complete: %w", err)
	}
	for i := len(payload.Conversation.Messages) - 1; i >= 0; i-- {
		m := payload.Conversation.Messages[i]
		if m.MessageType == "bot_response" {
			return m.Content, payload.Conversation.ID, nil
		}
	}
	return "", "", fmt.Errorf("no bot response in complete event")
}

func listKnowledgeBaseIDs(ctx context.Context, sdk *pipeshub.Pipeshub) ([]string, error) {
	orgRes, err := sdk.Organizations.GetCurrentOrganization(ctx)
	if err != nil {
		return nil, fmt.Errorf("get current organization: %w", err)
	}
	if orgRes == nil || orgRes.Organization == nil || orgRes.Organization.ID == nil || *orgRes.Organization.ID == "" {
		return nil, fmt.Errorf("get current organization: missing organization id")
	}
	parentID := "knowledgeBase_" + *orgRes.Organization.ID

	kbsRes, err := sdk.KnowledgeHub.GetKnowledgeHubChildNodes(ctx, operations.GetKnowledgeHubChildNodesRequest{
		ParentType: operations.ParentTypeApp,
		ParentID:   parentID,
	})
	if err != nil {
		return nil, fmt.Errorf("list knowledge bases: %w", err)
	}
	if kbsRes == nil || kbsRes.KnowledgeHubNodesResponse == nil {
		return nil, fmt.Errorf("list knowledge bases: empty response")
	}
	items := kbsRes.KnowledgeHubNodesResponse.GetItems()
	ids := make([]string, 0, len(items))
	for _, kb := range items {
		ids = append(ids, kb.ID)
	}
	return ids, nil
}
