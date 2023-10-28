package router

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/kbgod/illuminate"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestContext_GetState(t *testing.T) {
	ctx := new(Context)
	if ctx.GetState() != nil {
		t.Errorf("ctx.GetState() = %v; want <nil>", ctx.GetState())
	}
	state := "test"
	ctx.state = &state

	if ctxState := ctx.GetState(); ctxState == nil {
		t.Errorf("ctx.GetState() = %v; want test", ctxState)
	} else if *ctxState != "test" {
		t.Errorf("ctx.GetState() = %s; want test", *ctxState)
	}
}

func TestContext_SetState(t *testing.T) {
	ctx := new(Context)
	if ctx.state != nil {
		t.Errorf("ctx.state = %v; want <nil>", ctx.state)
	}
	ctx.SetState("test")

	if ctx.state == nil || *ctx.state != "test" {
		t.Errorf("ctx.state = %v; want test", ctx.state)
	}
}

func TestContext_Next(t *testing.T) {
	router := New(&illuminate.Bot{})
	var (
		firstHandlerCalled  bool
		secondHandlerCalled bool

		routeHandlerCalled       bool
		secondRouteHandlerCalled bool
	)
	router.Use(func(ctx *Context) error {
		firstHandlerCalled = true
		return ctx.Next()
	})
	router.Use(func(ctx *Context) error {
		secondHandlerCalled = true
		return ctx.Next()
	})
	router.On(Command("test"), func(ctx *Context) error {
		routeHandlerCalled = true
		return ctx.Next()
	}, func(ctx *Context) error {
		secondRouteHandlerCalled = true
		return ctx.Next()
	})

	if err := router.HandleUpdate(context.Background(), &illuminate.Update{
		Message: &illuminate.Message{
			Text: "/test",
		},
	}); err != nil {
		t.Errorf("router.HandleUpdate() = %v; want <nil>", err)
	}

	if !firstHandlerCalled {
		t.Errorf("firstHandlerCalled = %v; want true", firstHandlerCalled)
	}
	if !secondHandlerCalled {
		t.Errorf("secondHandlerCalled = %v; want true", secondHandlerCalled)
	}
	if !routeHandlerCalled {
		t.Errorf("routeHandlerCalled = %v; want true", routeHandlerCalled)
	}
	if !secondRouteHandlerCalled {
		t.Errorf("secondRouteHandlerCalled = %v; want true", secondRouteHandlerCalled)
	}
	firstHandlerCalled = false
	secondHandlerCalled = false
	routeHandlerCalled = false
	secondRouteHandlerCalled = false

	if err := router.HandleUpdate(context.Background(), &illuminate.Update{}); !errors.Is(err, ErrRouteNotFound) {
		t.Errorf("router.HandleUpdate() = %v; want <nil>", err)
	}

	if !firstHandlerCalled {
		t.Errorf("firstHandlerCalled = %v; want true", firstHandlerCalled)
	}
	if !secondHandlerCalled {
		t.Errorf("secondHandlerCalled = %v; want true", secondHandlerCalled)
	}
	if routeHandlerCalled {
		t.Errorf("routeHandlerCalled = %v; want false", routeHandlerCalled)
	}
	if secondRouteHandlerCalled {
		t.Errorf("secondRouteHandlerCalled = %v; want false", secondRouteHandlerCalled)
	}
}

func TestContext_Message(t *testing.T) {
	r := New(&illuminate.Bot{})
	ctx := newContext(nil, r, &illuminate.Update{
		CallbackQuery: &illuminate.CallbackQuery{
			Message: &illuminate.Message{
				MessageID: 1,
			},
		},
	})
	if ctx.Message() == nil || ctx.Message().MessageID != 1 {
		t.Errorf("ctx.Message()[CallbackQuery] = %v; want 1", ctx.Message())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		EditedMessage: &illuminate.Message{
			MessageID: 1,
		},
	})
	if ctx.Message() == nil || ctx.Message().MessageID != 1 {
		t.Errorf("ctx.Message()[EditedMessage] = %v; want 1", ctx.Message())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		ChannelPost: &illuminate.Message{
			MessageID: 1,
		},
	})
	if ctx.Message() == nil || ctx.Message().MessageID != 1 {
		t.Errorf("ctx.Message()[ChannelPost] = %v; want 1", ctx.Message())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		EditedChannelPost: &illuminate.Message{
			MessageID: 1,
		},
	})
	if ctx.Message() == nil || ctx.Message().MessageID != 1 {
		t.Errorf("ctx.Message()[EditedChannelPost] = %v; want 1", ctx.Message())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			MessageID: 1,
		},
	})
	if ctx.Message() == nil || ctx.Message().MessageID != 1 {
		t.Errorf("ctx.Message()[Message] = %v; want 1", ctx.Message())
	}
	ctx = newContext(nil, r, &illuminate.Update{})
	if ctx.Message() != nil {
		t.Errorf("ctx.Message() = %v; want <nil>", ctx.Message())
	}
}

func TestContext_Sender(t *testing.T) {
	r := New(&illuminate.Bot{})
	ctx := newContext(nil, r, &illuminate.Update{
		CallbackQuery: &illuminate.CallbackQuery{
			From: illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[CallbackQuery] = %v; want 1", ctx.Sender())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		EditedMessage: &illuminate.Message{
			From: &illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[EditedMessage] = %v; want 1", ctx.Sender())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		ChannelPost: &illuminate.Message{
			From: &illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[ChannelPost] = %v; want 1", ctx.Sender())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		EditedChannelPost: &illuminate.Message{
			From: &illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[EditedChannelPost] = %v; want 1", ctx.Sender())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			From: &illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[Message] = %v; want 1", ctx.Sender())
	}

	ctx = newContext(nil, r, &illuminate.Update{
		InlineQuery: &illuminate.InlineQuery{
			From: illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[InlineQuery] = %v; want 1", ctx.Sender())
	}

	ctx = newContext(nil, r, &illuminate.Update{
		ShippingQuery: &illuminate.ShippingQuery{
			From: illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[ShippingQuery] = %v; want 1", ctx.Sender())
	}

	ctx = newContext(nil, r, &illuminate.Update{
		PreCheckoutQuery: &illuminate.PreCheckoutQuery{
			From: illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[PreCheckoutQuery] = %v; want 1", ctx.Sender())
	}

	ctx = newContext(nil, r, &illuminate.Update{
		PollAnswer: &illuminate.PollAnswer{
			User: &illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[PollAnswer] = %v; want 1", ctx.Sender())
	}

	ctx = newContext(nil, r, &illuminate.Update{
		MyChatMember: &illuminate.ChatMemberUpdated{
			From: illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[MyChatMember] = %v; want 1", ctx.Sender())
	}

	ctx = newContext(nil, r, &illuminate.Update{
		ChatMember: &illuminate.ChatMemberUpdated{
			From: illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[ChatMember] = %v; want 1", ctx.Sender())
	}

	ctx = newContext(nil, r, &illuminate.Update{
		ChatJoinRequest: &illuminate.ChatJoinRequest{
			From: illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.Sender() == nil || ctx.Sender().ID != 1 {
		t.Errorf("ctx.Sender()[ChatJoinRequest] = %v; want 1", ctx.Sender())
	}

	ctx = newContext(nil, r, &illuminate.Update{})
	if ctx.Sender() != nil {
		t.Errorf("ctx.Sender() = %v; want <nil>", ctx.Sender())
	}
}

func TestContext_Chat(t *testing.T) {
	r := New(&illuminate.Bot{})
	ctx := newContext(nil, r, &illuminate.Update{
		CallbackQuery: &illuminate.CallbackQuery{
			Message: &illuminate.Message{
				Chat: illuminate.Chat{
					ID: 1,
				},
			},
		},
	})
	if ctx.Chat() == nil || ctx.Chat().ID != 1 {
		t.Errorf("ctx.Chat()[CallbackQuery] = %v; want 1", ctx.Chat())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		EditedMessage: &illuminate.Message{
			Chat: illuminate.Chat{
				ID: 1,
			},
		},
	})
	if ctx.Chat() == nil || ctx.Chat().ID != 1 {
		t.Errorf("ctx.Chat()[EditedMessage] = %v; want 1", ctx.Chat())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		ChannelPost: &illuminate.Message{
			Chat: illuminate.Chat{
				ID: 1,
			},
		},
	})
	if ctx.Chat() == nil || ctx.Chat().ID != 1 {
		t.Errorf("ctx.Chat()[ChannelPost] = %v; want 1", ctx.Chat())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		EditedChannelPost: &illuminate.Message{
			Chat: illuminate.Chat{
				ID: 1,
			},
		},
	})
	if ctx.Chat() == nil || ctx.Chat().ID != 1 {
		t.Errorf("ctx.Chat()[EditedChannelPost] = %v; want 1", ctx.Chat())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			Chat: illuminate.Chat{
				ID: 1,
			},
		},
	})
	if ctx.Chat() == nil || ctx.Chat().ID != 1 {
		t.Errorf("ctx.Chat()[Message] = %v; want 1", ctx.Chat())
	}

	ctx = newContext(nil, r, &illuminate.Update{
		MyChatMember: &illuminate.ChatMemberUpdated{
			Chat: illuminate.Chat{
				ID: 1,
			},
		},
	})
	if ctx.Chat() == nil || ctx.Chat().ID != 1 {
		t.Errorf("ctx.Chat()[MyChatMember] = %v; want 1", ctx.Chat())
	}

	ctx = newContext(nil, r, &illuminate.Update{
		ChatMember: &illuminate.ChatMemberUpdated{
			Chat: illuminate.Chat{
				ID: 1,
			},
		},
	})
	if ctx.Chat() == nil || ctx.Chat().ID != 1 {
		t.Errorf("ctx.Chat()[ChatMember] = %v; want 1", ctx.Chat())
	}

	ctx = newContext(nil, r, &illuminate.Update{
		ChatJoinRequest: &illuminate.ChatJoinRequest{
			Chat: illuminate.Chat{
				ID: 1,
			},
		},
	})
	if ctx.Chat() == nil || ctx.Chat().ID != 1 {
		t.Errorf("ctx.Chat()[ChatJoinRequest] = %v; want 1", ctx.Chat())
	}

	ctx = newContext(nil, r, &illuminate.Update{})
	if ctx.Chat() != nil {
		t.Errorf("ctx.Chat() = %v; want <nil>", ctx.Chat())
	}
}

func TestContext_ChatID(t *testing.T) {
	r := New(&illuminate.Bot{})
	ctx := newContext(nil, r, &illuminate.Update{
		ChannelPost: &illuminate.Message{
			Chat: illuminate.Chat{
				ID: 1,
			},
		},
	})
	if ctx.ChatID().PeerID() != "1" {
		t.Errorf("ctx.ChatID()[ChannelPost] = %v; want 1", ctx.ChatID())
	}
	ctx = newContext(nil, r, &illuminate.Update{
		PreCheckoutQuery: &illuminate.PreCheckoutQuery{
			From: illuminate.User{
				ID: 1,
			},
		},
	})
	if ctx.ChatID().PeerID() != "1" {
		t.Errorf("ctx.ChatID()[ChatMember] = %v; want 1", ctx.ChatID())
	}

	ctx = newContext(nil, r, &illuminate.Update{})
	if ctx.ChatID().PeerID() != "0" {
		t.Errorf("ctx.ChatID() = %v; want 0", ctx.ChatID())
	}
}

func TestContext_CommandArgs(t *testing.T) {
	r := New(&illuminate.Bot{})
	ctx := newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			Text: "/test arg1 arg2",
		},
	})
	if len(ctx.CommandArgs()) != 2 {
		t.Errorf("len(ctx.CommandArgs()) = %d; want 2", len(ctx.CommandArgs()))
	}
	if ctx.CommandArgs()[0] != "arg1" {
		t.Errorf("ctx.CommandArgs()[0] = %s; want arg1", ctx.CommandArgs()[0])
	}
	if ctx.CommandArgs()[1] != "arg2" {
		t.Errorf("ctx.CommandArgs()[1] = %s; want arg2", ctx.CommandArgs()[1])
	}

	ctx = newContext(nil, r, &illuminate.Update{
		Message: &illuminate.Message{
			Text: "/test",
		},
	})
	if len(ctx.CommandArgs()) != 0 {
		t.Errorf("len(ctx.CommandArgs()) = %d; want 0", len(ctx.CommandArgs()))
	}

	ctx = newContext(nil, r, &illuminate.Update{})
	if ctx.CommandArgs() != nil {
		t.Errorf("ctx.CommandArgs() = %v; want <nil>", ctx.CommandArgs())
	}
}

type fakeHttpClient struct {
}

func (fakeHttpClient) Do(req *http.Request) (*http.Response, error) {
	sendMessage := map[string]string{}
	_ = json.NewDecoder(req.Body).Decode(&sendMessage)
	chatID, _ := strconv.ParseInt(sendMessage["chat_id"], 10, 64)
	replyToMessageID, _ := strconv.ParseInt(sendMessage["reply_to_message_id"], 10, 64)

	if req.URL.String() == "https://api.telegram.org/bot123:test/sendMessage" {
		message := &illuminate.Message{
			Chat: illuminate.Chat{
				ID: illuminate.ChatID(chatID),
			},
			MessageID: 123,
			Text:      sendMessage["text"],
			ReplyToMessage: &illuminate.Message{
				MessageID: replyToMessageID,
			},
		}
		msgBytes, _ := json.Marshal(message)
		respBytes, _ := json.Marshal(&illuminate.Response{
			Ok:     true,
			Result: msgBytes,
		})
		return &http.Response{
			Body: io.NopCloser(bytes.NewBuffer(respBytes)),
		}, nil
	}
	return nil, errors.New("hui")
}

func TestContext_Reply(t *testing.T) {
	cl := fakeHttpClient{}
	r := New(illuminate.NewBot(illuminate.WithToken("123:test"), illuminate.WithHttpDoer(cl)))
	ctx := newContext(context.Background(), r, &illuminate.Update{
		Message: &illuminate.Message{
			Chat: illuminate.Chat{
				ID: 1,
			},
		},
	})
	m, err := ctx.Reply("test", &illuminate.SendMessageOpts{
		ReplyToMessageID: 225,
	})
	if err != nil {
		t.Errorf("ctx.Reply() = %v; want <nil>", err)
	} else if m == nil {
		t.Errorf("ctx.Reply() = %v; want not <nil>", m)
	} else if m.MessageID != 123 {
		t.Errorf("ctx.Reply() = %d; want 123", m.MessageID)
	} else if m.ReplyToMessage == nil {
		t.Errorf("ctx.Reply() = %v; want not <nil>", m.ReplyToMessage)
	} else if m.ReplyToMessage.MessageID != 225 {
		t.Errorf("ctx.Reply() = %d; want 225", m.ReplyToMessage.MessageID)
	} else if m.Text != "test" {
		t.Errorf("ctx.Reply() = %s; want test", m.Text)
	} else if m.Chat.ID != 1 {
		t.Errorf("ctx.Reply() = %d; want 1", m.Chat.ID)
	}
}
