// Command connector streams one answer scoped to a single connector, using
// Filters.Apps.
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

const defaultConnectorName = "ABC News RSS"

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

	connectorName := os.Getenv("PIPESHUB_CONNECTOR_NAME")
	if connectorName == "" {
		connectorName = defaultConnectorName
	}

	ctx := context.Background()

	connectorID, err := findConnectorIDByName(ctx, client, connectorName)
	if err != nil {
		log.Fatal(err)
	}

	query := "What are some latest news from stock market?"

	res, err := client.Conversations.StreamChat(ctx, components.ConversationStreamRequest{
		Query:    query,
		ChatMode: components.ConversationStreamRequestChatModeInternalSearch,
		Filters:  &components.Filters{Apps: []string{connectorID}},
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

func findConnectorIDByName(ctx context.Context, sdk *pipeshub.Pipeshub, name string) (string, error) {
	res, err := sdk.KnowledgeHub.GetKnowledgeHubRootNodes(ctx, operations.GetKnowledgeHubRootNodesRequest{})
	if err != nil {
		return "", fmt.Errorf("get knowledge hub root nodes: %w", err)
	}
	if res == nil || res.KnowledgeHubNodesResponse == nil {
		return "", fmt.Errorf("get knowledge hub root nodes: empty response")
	}

	for _, n := range res.KnowledgeHubNodesResponse.GetItems() {
		if n.Name == name && n.Origin == components.KnowledgeHubNodeOriginConnector {
			return n.ID, nil
		}
	}

	return "", fmt.Errorf("connector %q not found", name)
}
