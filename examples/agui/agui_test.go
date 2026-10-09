package agui

import "testing"

// A sub-agent's RUN_FINISHED arrives before the terminal one and carries no
// result. Ending the stream there would report an empty answer.
func TestCollectorIgnoresChildTerminalFrames(t *testing.T) {
	var c Collector

	frames := []struct {
		event string
		data  string
	}{
		{Custom, `{"type":"CUSTOM","name":"conversation_created","value":{"conversationId":"conv-1","title":"Cheese"}}`},
		{Custom, `{"type":"CUSTOM","name":"tool_unavailable","value":"jira.create_issue"}`},
		{RunStarted, `{"type":"RUN_STARTED","runId":"run-1","threadId":"t-1"}`},
		{StepStarted, `{"type":"STEP_STARTED","stepName":"sub_agent:docs","runId":"run-1"}`},
		{RunFinished, `{"type":"RUN_FINISHED","runId":"child-1","parentRunId":"run-1"}`},
		{RunError, `{"type":"RUN_ERROR","runId":"child-2","parentRunId":"run-1","message":"sub-agent failed","code":"agent_error"}`},
		{TextMessageContent, `{"type":"TEXT_MESSAGE_CONTENT","messageId":"m1","delta":"Spencer "}`},
		{StateDelta, `{"type":"STATE_DELTA","delta":[{"op":"replace","path":"/normalizedAnswer","value":"Spencer Johnson [1]"}]}`},
		{TextMessageContent, `{"type":"TEXT_MESSAGE_CONTENT","messageId":"m1","delta":"Johnson"}`},
	}

	for _, f := range frames {
		done, err := c.Handle(f.event, f.data)
		if err != nil {
			t.Fatalf("Handle(%s): unexpected error: %v", f.event, err)
		}
		if done {
			t.Fatalf("Handle(%s): reported done before the root RUN_FINISHED", f.event)
		}
	}

	if c.ConversationID != "conv-1" {
		t.Errorf("ConversationID = %q, want %q", c.ConversationID, "conv-1")
	}
	if c.Streamed != "Spencer Johnson" {
		t.Errorf("Streamed = %q, want %q", c.Streamed, "Spencer Johnson")
	}
	if c.NormalizedAnswer != "Spencer Johnson [1]" {
		t.Errorf("NormalizedAnswer = %q, want %q", c.NormalizedAnswer, "Spencer Johnson [1]")
	}

	root := `{"type":"RUN_FINISHED","result":{"conversation":{"_id":"conv-1","title":"Cheese",` +
		`"messages":[{"_id":"m0","messageType":"user_query","content":"Who moved the cheese?"},` +
		`{"_id":"m1","messageType":"bot_response","content":"Spencer Johnson [R1]"}]},"recordsUsed":3}}`
	done, err := c.Handle(RunFinished, root)
	if err != nil {
		t.Fatalf("Handle(root RUN_FINISHED): unexpected error: %v", err)
	}
	if !done {
		t.Fatal("Handle(root RUN_FINISHED): want done")
	}

	// The persisted message wins over the raw tokens, since citation refs are
	// rewritten as the answer is finalized.
	if got := c.Answer(); got != "Spencer Johnson [R1]" {
		t.Errorf("Answer() = %q, want the persisted message", got)
	}
	if c.Result.RecordsUsed != 3 {
		t.Errorf("RecordsUsed = %d, want 3", c.Result.RecordsUsed)
	}
}

// A root RUN_ERROR is the one failure that ends the stream.
func TestCollectorRootErrorTerminates(t *testing.T) {
	var c Collector

	_, err := c.Handle(RunError, `{"type":"RUN_ERROR","message":"upstream exploded","code":"streaming_error"}`)
	if err == nil {
		t.Fatal("Handle(root RUN_ERROR): want an error")
	}
	streamErr, ok := err.(*Error)
	if !ok {
		t.Fatalf("Handle(root RUN_ERROR): got %T, want *agui.Error", err)
	}
	if streamErr.Message != "upstream exploded" || streamErr.Code != "streaming_error" {
		t.Errorf("got %+v, want message/code from the frame", streamErr)
	}
}

// Unknown event names must be ignored rather than treated as failures.
func TestCollectorIgnoresUnknownEvents(t *testing.T) {
	var c Collector

	done, err := c.Handle("SOME_FUTURE_EVENT", `{"type":"SOME_FUTURE_EVENT","whatever":[1,2,3]}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if done {
		t.Fatal("unknown event reported done")
	}
}
