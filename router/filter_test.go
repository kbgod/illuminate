package router

import (
	"github.com/kbgod/illuminate"
	"testing"
)

func TestCommandWithAt(t *testing.T) {
	r := New(&illuminate.Bot{})
	if !CommandWithAt("test", "testbot")(newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			Text: "/test@testbot",
		},
	})) {
		t.Error("CommandWithAt failed")
	}
	if CommandWithAt("test", "testbot")(newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			Text: "/invalid@testbot",
		},
	})) {
		t.Error("CommandWithAt (invalid command) failed")
	}
	if CommandWithAt("test", "testbot")(newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			Text: "/test@invalid",
		},
	})) {
		t.Error("CommandWithAt (invalid bot) failed")
	}
	if CommandWithAt("test", "testbot")(newContext(nil, r, &illuminate.Update{})) {
		t.Error("CommandWithAt (empty message) failed")
	}
}

func TestTextContains(t *testing.T) {
	r := New(&illuminate.Bot{})
	if !TextContains("test")(newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			Text: "test",
		},
	})) {
		t.Error("TextContains failed")
	}
	if !TextContains("test")(newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			Text: "test123",
		},
	})) {
		t.Error("TextContains failed")
	}
	if !TextContains("test")(newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			Text: "123test",
		},
	})) {
		t.Error("TextContains failed")
	}
	if TextContains("test")(newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			Text: "123",
		},
	})) {
		t.Error("TextContains (invalid text) failed")
	}
	if TextContains("test")(newContext(nil, r, &illuminate.Update{})) {
		t.Error("TextContains (empty message) failed")
	}
}
