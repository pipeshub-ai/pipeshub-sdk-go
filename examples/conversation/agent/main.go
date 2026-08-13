// Command agent streams one answer through the universal agent loop —
// chatMode "agent" on the ordinary /conversations/stream route, with no agent
// key involved.
//
// This mode is what Tools and AgentCapabilities apply to; they are ignored by
// internal_search and web_search. It is also the mode that spawns sub-agents,
// so this is the example where child runs carrying parentRunId actually appear.
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

	client, err := auth.NewClient(
		os.Getenv("PIPESHUB_TEST_USER_EMAIL"),
		os.Getenv("PIPESHUB_TEST_USER_PASSWORD"),
	)
	if err != nil {
		log.Fatal(err)
	}

	query := "Summarise what our team shipped recently and what is still open."

	res, err := client.Conversations.StreamChat(context.Background(), components.ConversationStreamRequest{
		Query:    query,
		ChatMode: components.ConversationStreamRequestChatModeAgent,

		// Tools is an allow-list of fully-qualified action names. Omitting the
		// field entirely offers every configured tool; an empty (non-nil) slice
		// offers none.
		Tools: []string{"jira.create_issue", "confluence.search_content"},

		// Each capability defaults to enabled when its field is absent, so this
		// runs the agent with internal and deep search on, web search off.
		AgentCapabilities: &components.AgentCapabilities{
			WebSearch: pipeshub.Bool(false),
		},
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
	if answer := c.Answer(); answer != c.Streamed {
		fmt.Printf("\nFinal answer:\n%s\n", answer)
	}
	fmt.Printf("\nconversation: %s\n", c.ConversationID)
	fmt.Printf("records used: %d\n", c.Result.RecordsUsed)
}
