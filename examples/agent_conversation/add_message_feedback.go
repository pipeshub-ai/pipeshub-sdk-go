// Create an agent conversation (streaming, KB filter only), then submit feedback on the last bot response.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/add_message_feedback.go <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"

	"enterprise_search/auth"
)

const (
	// Stable key for the agent that owns the conversation.
	agentKey = "02a7d998-d21b-4015-aaf7-5cda765c1012"

	// Knowledge-base / record-group id (Filters.Kb).
	kbID = "8747da12-4724-4a95-ac92-827b88d79647"

	// First user message when creating the conversation.
	firstMessage = "What are some latest tech news?"

	// Positive free-text feedback (no negative comment — omit that field).
	positiveFeedbackComment = "The answer stayed on topic and covered the main points without filler. " +
		"Citations pointed to relevant sources I could verify, and the explanation was structured so " +
		"each section built on the last—easy to follow on a first read."
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/add_message_feedback.go <path-to-.env>")
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
		Kb: []string{kbID},
	}

	convID, botMessageID, answer, err := createAgentConversation(ctx, sdk, firstMessage, filters)
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	log.Printf("conversation id: %s", convID)
	log.Printf("last bot message id: %s", botMessageID)
	fmt.Printf("\nBot response (%d chars):\n%s\n", len(answer), answer)

	if err := submitMessageFeedback(ctx, sdk, convID, botMessageID); err != nil {
		log.Fatalf("submit feedback: %v", err)
	}
	log.Println("feedback submitted successfully")
}

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, query string, filters *components.Filters) (convID, botMessageID, answer string, err error) {
	chatMode := components.AgentStreamCreateConversationRequestChatModeAuto

	res, err := sdk.Agents.StreamAgentConversation(ctx, agentKey, components.AgentStreamCreateConversationRequest{
		Query:    query,
		Filters:  filters,
		ChatMode: chatMode.ToPointer(),
	})
	if err != nil {
		return "", "", "", fmt.Errorf("stream agent conversation: %w", err)
	}
	if res.AgentStreamSSEEvent == nil {
		return "", "", "", fmt.Errorf("no SSE stream returned")
	}
	stream := res.AgentStreamSSEEvent
	defer stream.Close()

	fmt.Printf("You: %s\n\nBot: ", query)

	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.AgentStreamSSEEventEventComplete:
			answer, convID, botMessageID, err = decodeConversationComplete(*ev.Data)
			if err != nil {
				return "", "", "", err
			}
			fmt.Println(answer)
			if convID == "" {
				return "", "", "", fmt.Errorf("complete event missing conversation id")
			}
			if botMessageID == "" {
				return "", "", "", fmt.Errorf("complete event missing bot message id")
			}
			return convID, botMessageID, answer, nil
		case components.AgentStreamSSEEventEventError:
			return "", "", "", fmt.Errorf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		return "", "", "", fmt.Errorf("stream: %w", err)
	}
	return "", "", "", fmt.Errorf("stream ended without complete event")
}

func submitMessageFeedback(ctx context.Context, sdk *pipeshub.Pipeshub, convID, messageID string) error {
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
	if res.MessageFeedbackUpdateResponse == nil {
		return fmt.Errorf("no feedback response returned")
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
	log.Printf("feedback metrics timeToFeedback(ms): %.0f", metrics.GetTimeToFeedback())
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

func decodeConversationComplete(data string) (answer, convID, botMessageID string, err error) {
	var payload struct {
		Conversation struct {
			ID       string `json:"_id"`
			Messages []struct {
				ID          string `json:"_id"`
				MessageType string `json:"messageType"`
				Content     string `json:"content"`
			} `json:"messages"`
		} `json:"conversation"`
	}
	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		return "", "", "", fmt.Errorf("decode complete: %w", err)
	}
	convID = payload.Conversation.ID
	for i := len(payload.Conversation.Messages) - 1; i >= 0; i-- {
		m := payload.Conversation.Messages[i]
		if m.MessageType == "bot_response" {
			return m.Content, convID, m.ID, nil
		}
	}
	return "", convID, "", fmt.Errorf("no bot_response in complete event")
}
