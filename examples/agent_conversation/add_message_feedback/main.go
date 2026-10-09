// Create an agent conversation (streaming, KB filter only), then submit
// feedback on the last bot response.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/add_message_feedback <path-to-.env>
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

const (
	// First user message when creating the conversation.
	firstMessage = "Who moved the cheese?"

	// Positive free-text feedback (no negative comment — omit that field).
	positiveFeedbackComment = "The answer stayed on topic and covered the main points without filler. " +
		"Citations pointed to relevant sources I could verify, and the explanation was structured so " +
		"each section built on the last—easy to follow on a first read."
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/add_message_feedback <path-to-.env>")
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

	// Leaving filters nil lets the agent use its stored knowledge config. An
	// empty-but-non-nil Filters would instead mean "no knowledge sources at all".
	var filters *components.Filters
	if kbID := os.Getenv("KB_ID"); kbID != "" {
		filters = &components.Filters{Kb: []string{kbID}}
	}

	ctx := context.Background()

	convID, botMessageID, answer, err := createAgentConversation(ctx, sdk, agentKey, firstMessage, filters)
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	log.Printf("conversation id: %s", convID)
	log.Printf("last bot message id: %s", botMessageID)
	fmt.Printf("\nBot response (%d chars):\n%s\n", len(answer), answer)

	if err := submitMessageFeedback(ctx, sdk, agentKey, convID, botMessageID); err != nil {
		log.Fatalf("submit feedback: %v", err)
	}
	log.Println("feedback submitted successfully")
}

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, query string, filters *components.Filters) (convID, botMessageID, answer string, err error) {
	res, err := sdk.Agents.StreamAgentConversation(ctx, agentKey, components.AgentStreamCreateConversationRequest{
		Query:   query,
		Filters: filters,
		// Scoped agent conversations accept only "quick".
		ChatMode: components.AgentStreamCreateConversationRequestChatModeQuick,
	})
	if err != nil {
		return "", "", "", fmt.Errorf("stream agent conversation: %w", err)
	}
	if res == nil || res.AgentStreamSSEEvent == nil {
		return "", "", "", fmt.Errorf("no SSE stream returned")
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
			return "", "", "", err
		}
		if done {
			break
		}
	}
	if err := stream.Err(); err != nil {
		return "", "", "", fmt.Errorf("stream: %w", err)
	}
	if !c.Done {
		return "", "", "", fmt.Errorf("stream ended without RUN_FINISHED")
	}
	fmt.Println()

	if c.ConversationID == "" {
		return "", "", "", fmt.Errorf("stream did not report a conversation id")
	}
	msg, ok := c.Result.Conversation.LastBotMessage()
	if !ok {
		return "", "", "", fmt.Errorf("no bot response in RUN_FINISHED result")
	}
	return c.ConversationID, msg.ID, msg.Content, nil
}

func submitMessageFeedback(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, convID, messageID string) error {
	body := components.MessageFeedbackSubmitRequest{
		IsHelpful: pipeshub.Bool(true),
		Categories: []components.MessageFeedbackSubmitRequestCategory{
			components.MessageFeedbackSubmitRequestCategoryExcellentAnswer,
			components.MessageFeedbackSubmitRequestCategoryHelpfulCitations,
			components.MessageFeedbackSubmitRequestCategoryWellExplained,
		},
		Comments: &components.MessageFeedbackSubmitRequestComments{
			Positive: pipeshub.String(positiveFeedbackComment),
			// Negative omitted — only send positive sentiment for this example.
		},
	}

	res, err := sdk.Agents.UpdateAgentConversationMessageFeedback(ctx, agentKey, convID, messageID, body)
	if err != nil {
		return fmt.Errorf("update agent conversation message feedback: %w", err)
	}
	logFeedbackUpdateResponse(res.MessageFeedbackUpdateResponse)
	return nil
}

func logFeedbackUpdateResponse(resp *components.MessageFeedbackUpdateResponse) {
	log.Printf("feedback conversation id: %s", resp.GetConversationID())
	log.Printf("feedback message id: %s", resp.GetMessageID())

	fb := resp.GetFeedback()
	if helpful := fb.GetIsHelpful(); helpful != nil {
		log.Printf("feedback isHelpful: %v", *helpful)
	}
	if cats := fb.GetCategories(); len(cats) > 0 {
		names := make([]string, len(cats))
		for i, c := range cats {
			names[i] = string(c)
		}
		log.Printf("feedback categories: %v", names)
	}
	if comments := fb.GetComments(); comments != nil {
		if positive := comments.GetPositive(); positive != nil && *positive != "" {
			log.Printf("feedback comments.positive (%d chars): %s", len(*positive), *positive)
		}
	}
	if provider := fb.GetFeedbackProvider(); provider != "" {
		log.Printf("feedback provider: %s", provider)
	}
	if ts := fb.GetTimestamp(); ts != 0 {
		log.Printf("feedback timestamp (epoch ms): %d", ts)
	}

	metrics := fb.GetMetrics()
	if t := metrics.GetTimeToFeedback(); t != 0 {
		log.Printf("feedback metrics timeToFeedback(ms): %.0f", t)
	}
	if ua := metrics.GetUserAgent(); ua != nil && *ua != "" {
		log.Printf("feedback metrics userAgent: %s", *ua)
	}

	meta := resp.GetMeta()
	if rid := meta.GetRequestID(); rid != "" {
		log.Printf("feedback meta requestId: %s", rid)
	}
	if !meta.GetTimestamp().IsZero() {
		log.Printf("feedback meta timestamp: %s", meta.GetTimestamp().Format(time.RFC3339))
	}
	if dur := meta.GetDuration(); dur != 0 {
		log.Printf("feedback meta duration(ms): %d", dur)
	}
}
