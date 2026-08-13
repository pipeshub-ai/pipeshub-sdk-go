// Package agui decodes the AG-UI Server-Sent Event vocabulary that every
// PipesHub conversation stream speaks.
//
// The generated SSE event types (components.ConversationStreamSSEEvent,
// components.AgentStreamSSEEvent, and four more) each carry an Event field of
// their own named string type and share no interface, so this package works on
// plain strings. Every call site converts with string(*ev.Event).
//
// This is the Go counterpart of integration-tests/helper/agui_sse.py in the
// pipeshub-ai repo; the two intentionally use the same names and rules.
package agui

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// AG-UI event names. These are the same strings as the generated enums, so
// agui.RunFinished == string(components.ConversationStreamSSEEventEventRunFinished).
const (
	RunStarted  = "RUN_STARTED"
	RunFinished = "RUN_FINISHED"
	RunError    = "RUN_ERROR"

	StepStarted  = "STEP_STARTED"
	StepFinished = "STEP_FINISHED"

	TextMessageStart   = "TEXT_MESSAGE_START"
	TextMessageContent = "TEXT_MESSAGE_CONTENT"
	TextMessageEnd     = "TEXT_MESSAGE_END"

	ReasoningStart          = "REASONING_START"
	ReasoningMessageStart   = "REASONING_MESSAGE_START"
	ReasoningMessageContent = "REASONING_MESSAGE_CONTENT"
	ReasoningMessageEnd     = "REASONING_MESSAGE_END"
	ReasoningEnd            = "REASONING_END"

	ToolCallStart  = "TOOL_CALL_START"
	ToolCallArgs   = "TOOL_CALL_ARGS"
	ToolCallEnd    = "TOOL_CALL_END"
	ToolCallResult = "TOOL_CALL_RESULT"

	StateDelta    = "STATE_DELTA"
	StateSnapshot = "STATE_SNAPSHOT"
	Custom        = "CUSTOM"
	Heartbeat     = "HEARTBEAT"
)

// ConversationCreated is the CUSTOM frame name that replaced the legacy
// "connected" event. Other names ride the same channel (tool_unavailable,
// ask_user_question), so callers must match on the name.
const ConversationCreated = "conversation_created"

// normalizedAnswerPath is the JSON-Patch path carrying the citation-rewritten
// answer on STATE_DELTA frames.
const normalizedAnswerPath = "/normalizedAnswer"

// Message is one turn of a persisted conversation.
type Message struct {
	ID          string `json:"_id"`
	MessageType string `json:"messageType"` // user_query | bot_response | tool_call | ...
	Role        string `json:"role"`
	Content     string `json:"content"`
}

// Conversation is the persisted conversation carried by RUN_FINISHED.
type Conversation struct {
	ID       string    `json:"_id"`
	Title    string    `json:"title"`
	Messages []Message `json:"messages"`
}

// LastBotMessage returns the most recent assistant turn.
func (c Conversation) LastBotMessage() (Message, bool) {
	for i := len(c.Messages) - 1; i >= 0; i-- {
		m := c.Messages[i]
		if m.MessageType != "bot_response" && m.Role != "assistant" {
			continue
		}
		if strings.TrimSpace(m.Content) != "" {
			return m, true
		}
	}
	return Message{}, false
}

// Result is the RUN_FINISHED `result` object — the same completion shape the
// legacy `complete` event carried at the top level, moved down one level.
type Result struct {
	Conversation Conversation `json:"conversation"`
	RecordsUsed  int          `json:"recordsUsed"`
	Answer       string       `json:"answer"`
}

// Error is a stream-level RUN_ERROR.
type Error struct {
	Message string
	Code    string
}

func (e *Error) Error() string {
	if e.Code == "" {
		return e.Message
	}
	return fmt.Sprintf("%s (%s)", e.Message, e.Code)
}

// Collector folds one AG-UI stream into a final answer. The zero value is ready
// to use; set Echo to mirror tokens as they arrive.
type Collector struct {
	// Echo receives TEXT_MESSAGE_CONTENT deltas as they arrive, for a live
	// typing effect. Nil disables echoing.
	Echo io.Writer

	// ConversationID and Title arrive on the CUSTOM/conversation_created frame,
	// before any answer token. On add-message and regenerate streams that frame
	// carries neither, because the caller already knows the conversation.
	ConversationID string
	Title          string

	// Streamed is the concatenation of every TEXT_MESSAGE_CONTENT delta.
	Streamed string

	// NormalizedAnswer is the citation-rewritten answer from STATE_DELTA. Unlike
	// Streamed it is replaced wholesale, not appended to.
	NormalizedAnswer string

	Result Result
	Done   bool
}

// Handle folds one frame into the collector. It reports done once the terminal
// RUN_FINISHED arrives, and returns *Error on a stream-level RUN_ERROR.
// Unknown event names are ignored, as the AG-UI contract requires.
func (c *Collector) Handle(event, data string) (bool, error) {
	switch event {
	case Custom:
		// Value stays raw until the name is known: other CUSTOM frames
		// (tool_unavailable, ask_user_question) carry unrelated shapes here.
		var f struct {
			Name  string          `json:"name"`
			Value json.RawMessage `json:"value"`
		}
		if err := json.Unmarshal([]byte(data), &f); err != nil {
			return false, fmt.Errorf("decode CUSTOM: %w", err)
		}
		if f.Name != ConversationCreated {
			return false, nil
		}
		var v struct {
			ConversationID string `json:"conversationId"`
			Title          string `json:"title"`
		}
		// Add-message and regenerate streams send only {message} here, because
		// the caller already knows which conversation it is streaming onto.
		if err := json.Unmarshal(f.Value, &v); err != nil {
			return false, fmt.Errorf("decode CUSTOM/%s: %w", ConversationCreated, err)
		}
		if v.ConversationID != "" {
			c.ConversationID = v.ConversationID
		}
		if v.Title != "" {
			c.Title = v.Title
		}

	case TextMessageContent:
		// `delta` is decoded here rather than in a shared envelope struct: it is
		// a string on this event but a JSON-Patch array on STATE_DELTA, and one
		// struct cannot hold both.
		var f struct {
			Delta string `json:"delta"`
		}
		if err := json.Unmarshal([]byte(data), &f); err != nil {
			return false, fmt.Errorf("decode TEXT_MESSAGE_CONTENT: %w", err)
		}
		c.Streamed += f.Delta
		if c.Echo != nil && f.Delta != "" {
			io.WriteString(c.Echo, f.Delta)
		}

	case StateDelta:
		var f struct {
			Delta []struct {
				Path  string `json:"path"`
				Value any    `json:"value"`
			} `json:"delta"`
		}
		if err := json.Unmarshal([]byte(data), &f); err != nil {
			return false, fmt.Errorf("decode STATE_DELTA: %w", err)
		}
		for _, op := range f.Delta {
			if op.Path != normalizedAnswerPath {
				continue
			}
			if s, ok := op.Value.(string); ok {
				c.NormalizedAnswer = s
			}
		}

	case RunFinished:
		var f struct {
			ParentRunID *string         `json:"parentRunId"`
			Result      json.RawMessage `json:"result"`
		}
		if err := json.Unmarshal([]byte(data), &f); err != nil {
			return false, fmt.Errorf("decode RUN_FINISHED: %w", err)
		}
		// Sub-agents emit their own RUN_FINISHED, which the gateway forwards
		// verbatim while swallowing the root one. Those child frames carry
		// parentRunId and no result, and arrive before the terminal frame, so
		// treating one as terminal ends the stream with an empty answer.
		if f.ParentRunID != nil || len(f.Result) == 0 {
			return false, nil
		}
		if err := json.Unmarshal(f.Result, &c.Result); err != nil {
			return false, fmt.Errorf("decode RUN_FINISHED result: %w", err)
		}
		if id := c.Result.Conversation.ID; id != "" {
			c.ConversationID = id
		}
		if t := c.Result.Conversation.Title; t != "" {
			c.Title = t
		}
		c.Done = true
		return true, nil

	case RunError:
		var f struct {
			ParentRunID *string `json:"parentRunId"`
			Message     string  `json:"message"`
			Code        string  `json:"code"`
		}
		if err := json.Unmarshal([]byte(data), &f); err != nil {
			return false, fmt.Errorf("decode RUN_ERROR: %w", err)
		}
		// Likewise: a child RUN_ERROR is one sub-agent failing, not the stream.
		if f.ParentRunID != nil {
			return false, nil
		}
		return false, &Error{Message: f.Message, Code: f.Code}
	}

	return false, nil
}

// Answer returns the assistant's reply, preferring the persisted message from
// RUN_FINISHED.
//
// Prefer this over reading Streamed directly. TEXT_MESSAGE_CONTENT carries raw
// append-only tokens, while citation references are rewritten as the answer is
// finalized, so the persisted message is the authoritative text.
func (c *Collector) Answer() string {
	if m, ok := c.Result.Conversation.LastBotMessage(); ok {
		return m.Content
	}
	for _, s := range []string{c.Result.Answer, c.NormalizedAnswer, c.Streamed} {
		if strings.TrimSpace(s) != "" {
			return s
		}
	}
	return ""
}
