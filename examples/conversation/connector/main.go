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

const connectorName = "ABC News RSS"

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

	ctx := context.Background()

	connectorID, err := findConnectorIDByName(ctx, client, connectorName)
	if err != nil {
		log.Fatal(err)
	}

	query := "What are some latest news from stock market?"
	chatMode := "internal_search"

	res, err := client.Conversations.StreamChat(ctx, components.CreateConversationRequest{
		Query:    query,
		ChatMode: &chatMode,
		Filters:  &components.Filters{Apps: []string{connectorID}},
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

func findConnectorIDByName(ctx context.Context, sdk *pipeshub.Pipeshub, name string) (string, error) {
	res, err := sdk.KnowledgeHub.GetKnowledgeHubRootNodes(ctx, operations.GetKnowledgeHubRootNodesRequest{})
	if err != nil {
		return "", fmt.Errorf("get knowledge hub root nodes: %w", err)
	}
	if res == nil || res.KnowledgeHubNodesResponse == nil {
		return "", fmt.Errorf("get knowledge hub root nodes: empty response")
	}

	for _, n := range res.KnowledgeHubNodesResponse.GetItems() {
		if n.Name == name && n.Origin == components.OriginConnector {
			return n.ID, nil
		}
	}

	return "", fmt.Errorf("connector %q not found", name)
}
