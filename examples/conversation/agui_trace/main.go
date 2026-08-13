// Command agui_trace prints every AG-UI frame a conversation stream emits,
// instead of only accumulating the answer text.
//
// Use it to see what the protocol actually exposes — reasoning turns, tool
// calls and their results, sub-agent step boundaries, and the parentRunId that
// separates a child run's lifecycle frames from the stream's own.
//
// Usage: go run . <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
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

	query := "What changed in our documentation this week, and why?"

	// Agent mode is the one that emits STEP_* and nested sub-agent runs, so it
	// produces the widest range of frames.
	res, err := client.Conversations.StreamChat(context.Background(), components.ConversationStreamRequest{
		Query:    query,
		ChatMode: components.ConversationStreamRequestChatModeAgent,
	})
	if err != nil {
		log.Fatalf("conversation: %v", err)
	}
	if res == nil || res.ConversationStreamSSEEvent == nil {
		log.Fatal("no SSE stream returned")
	}
	stream := res.ConversationStreamSSEEvent
	defer stream.Close()

	fmt.Printf("You: %s\n\n", query)

	// The collector folds the stream into an answer while trace() reports each
	// frame; both see every event.
	var c agui.Collector
	counts := map[string]int{}

	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		event, data := string(*ev.Event), *ev.Data
		counts[event]++
		trace(event, data)

		done, err := c.Handle(event, data)
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

	fmt.Printf("\n--- frame counts ---\n")
	for _, name := range []string{
		agui.RunStarted, agui.RunFinished, agui.StepStarted, agui.StepFinished,
		agui.TextMessageStart, agui.TextMessageContent, agui.TextMessageEnd,
		agui.ReasoningMessageContent, agui.ToolCallStart, agui.ToolCallArgs,
		agui.ToolCallResult, agui.StateDelta, agui.Custom, agui.Heartbeat,
	} {
		if n := counts[name]; n > 0 {
			fmt.Printf("  %-26s %d\n", name, n)
		}
	}

	fmt.Printf("\n--- answer ---\n%s\n", c.Answer())
	fmt.Printf("\nconversation: %s\n", c.ConversationID)
}

// frame holds the union of fields the trace reports. `delta` is deliberately
// absent: it is a string on TEXT_MESSAGE_CONTENT but a JSON-Patch array on
// STATE_DELTA, so it cannot be decoded here.
type frame struct {
	ParentRunID   *string `json:"parentRunId"`
	StepName      string  `json:"stepName"`
	ToolCallName  string  `json:"toolCallName"`
	DisplayName   string  `json:"displayName"`
	ArgsSummary   string  `json:"argsSummary"`
	Status        string  `json:"status"`
	ResultSummary string  `json:"resultSummary"`
	Name          string  `json:"name"`
	Message       string  `json:"message"`
	Code          string  `json:"code"`
}

func trace(event, data string) {
	var f frame
	if err := json.Unmarshal([]byte(data), &f); err != nil {
		fmt.Printf("[?    ] %-26s <undecodable payload>\n", event)
		return
	}

	// Sub-agent frames carry parentRunId. Their RUN_FINISHED and RUN_ERROR are
	// not terminal for the stream.
	scope := "root "
	if f.ParentRunID != nil {
		scope = "child"
	}

	detail := ""
	switch event {
	case agui.StepStarted, agui.StepFinished:
		detail = f.StepName
	case agui.ToolCallStart:
		detail = f.ToolCallName
		if f.DisplayName != "" {
			detail = fmt.Sprintf("%s (%s)", f.ToolCallName, f.DisplayName)
		}
	case agui.ToolCallArgs:
		detail = f.ArgsSummary
	case agui.ToolCallResult:
		detail = strings.TrimSpace(f.Status + " " + f.ResultSummary)
	case agui.Custom:
		detail = f.Name
	case agui.RunError:
		detail = fmt.Sprintf("%s (%s)", f.Message, f.Code)
	case agui.TextMessageContent, agui.ReasoningMessageContent:
		// Content frames arrive per token; reporting each one would drown the
		// trace, so just mark them.
		detail = "…"
	}

	if detail == "" {
		fmt.Printf("[%s] %s\n", scope, event)
		return
	}
	fmt.Printf("[%s] %-26s %s\n", scope, event, detail)
}
