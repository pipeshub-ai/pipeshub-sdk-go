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

const knowledgeBaseName = "SDK-test"

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

	kbID, err := findKnowledgeBaseID(ctx, sdk, knowledgeBaseName)
	if err != nil {
		log.Fatal(err)
	}

	query := "Who moved the cheese?"
	chatMode := "internal_search"

	res, err := sdk.Conversations.StreamChat(ctx, components.CreateConversationRequest{
		Query:    query,
		ChatMode: &chatMode,
		Filters:  &components.Filters{Kb: []string{kbID}},
	})
	if err != nil {
		log.Fatalf("conversation: %v", err)
	}
	if res.AssistantStreamSSEEvent == nil {
		log.Fatal("no SSE stream returned")
	}
	stream := res.AssistantStreamSSEEvent
	defer stream.Close()

	fmt.Printf("You: %s\n\nBot: ", query)

	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.AssistantStreamSSEEventEventComplete:
			var payload struct {
				Conversation struct {
					Messages []struct {
						MessageType string `json:"messageType"`
						Content     string `json:"content"`
					} `json:"messages"`
				} `json:"conversation"`
			}
			if err := json.Unmarshal([]byte(*ev.Data), &payload); err != nil {
				log.Fatalf("decode complete: %v", err)
			}
			for _, m := range payload.Conversation.Messages {
				if m.MessageType == "bot_response" {
					fmt.Println(m.Content)
					return
				}
			}
			log.Fatal("no bot response in complete event")
		case components.AssistantStreamSSEEventEventError:
			log.Fatalf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		log.Fatalf("stream: %v", err)
	}
}

func findKnowledgeBaseID(ctx context.Context, sdk *pipeshub.Pipeshub, name string) (string, error) {
	parentID, err := getKnowledgeBaseParentID(ctx, sdk)
	if err != nil {
		return "", err
	}

	kbsRes, err := sdk.KnowledgeHub.GetKnowledgeHubChildNodes(ctx, operations.GetKnowledgeHubChildNodesRequest{
		ParentType: operations.ParentTypeApp,
		ParentID:   parentID,
	})
	if err != nil {
		return "", fmt.Errorf("list knowledge bases: %w", err)
	}
	if kbsRes == nil || kbsRes.KnowledgeHubNodesResponse == nil {
		return "", fmt.Errorf("list knowledge bases: empty response")
	}

	fmt.Println("[debug] available knowledge bases:")
	for _, kb := range kbsRes.KnowledgeHubNodesResponse.GetItems() {
		fmt.Printf("  - %q (id=%s)\n", kb.Name, kb.ID)
		if kb.Name == name {
			return kb.ID, nil
		}
	}

	return "", fmt.Errorf("knowledge base %q not found for parent %q", name, parentID)
}

// Before getting the knowledge bases available for the org, we need to find out the org id
// Parent ID is the org id prefixed with "knowledgeBase_"

func getKnowledgeBaseParentID(ctx context.Context, sdk *pipeshub.Pipeshub) (string, error) {
	orgRes, err := sdk.Organizations.GetCurrentOrganization(ctx)
	if err != nil {
		return "", fmt.Errorf("get current organization: %w", err)
	}
	if orgRes == nil || orgRes.Organization == nil || orgRes.Organization.ID == nil || *orgRes.Organization.ID == "" {
		return "", fmt.Errorf("get current organization: missing organization id")
	}

	return "knowledgeBase_" + *orgRes.Organization.ID, nil
}
