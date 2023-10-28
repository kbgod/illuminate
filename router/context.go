package router

import (
	"context"
	"github.com/kbgod/illuminate"
	"strings"
)

type Context struct {
	state        *string
	router       *Router
	route        *Route
	indexRoute   int
	indexHandler int

	Context context.Context
	Update  *illuminate.Update
	Bot     *illuminate.Bot
}

func newContext(ctx context.Context, router *Router, update *illuminate.Update) *Context {
	return &Context{
		Context:      ctx,
		indexHandler: -1,
		indexRoute:   -1,
		Update:       update,
		router:       router,
		Bot:          router.bot,
	}
}

func (ctx *Context) GetState() *string {
	return ctx.state
}

func (ctx *Context) SetState(state string) {
	ctx.state = &state
}

func (ctx *Context) Next() error {
	var err error
	ctx.indexHandler++
	if ctx.route == nil && ctx.indexHandler < len(ctx.router.handlers) {
		err = ctx.router.handlers[ctx.indexHandler](ctx)
	} else if ctx.route != nil && ctx.indexHandler < len(ctx.route.handlers) {
		return ctx.route.handlers[ctx.indexHandler](ctx)
	} else if ctx.route == nil {
		err = ctx.router.next(ctx)
	}
	return err
}

// HELPER GETTERS

func (ctx *Context) Message() *illuminate.Message {
	if m := firstNotNil(
		ctx.Update.Message,
		ctx.Update.EditedMessage,
		ctx.Update.ChannelPost,
		ctx.Update.EditedChannelPost,
	); m != nil {
		return m
	}
	if ctx.Update.CallbackQuery != nil && ctx.Update.CallbackQuery.Message != nil {
		return ctx.Update.CallbackQuery.Message
	}

	return nil
}

func (ctx *Context) Sender() *illuminate.User {
	switch {
	case ctx.Update.CallbackQuery != nil:
		return &ctx.Update.CallbackQuery.From
	case ctx.Message() != nil:
		return ctx.Message().From
	case ctx.Update.InlineQuery != nil:
		return &ctx.Update.InlineQuery.From
	case ctx.Update.ShippingQuery != nil:
		return &ctx.Update.ShippingQuery.From
	case ctx.Update.PreCheckoutQuery != nil:
		return &ctx.Update.PreCheckoutQuery.From
	case ctx.Update.PollAnswer != nil:
		return ctx.Update.PollAnswer.User
	case ctx.Update.MyChatMember != nil:
		return &ctx.Update.MyChatMember.From
	case ctx.Update.ChatMember != nil:
		return &ctx.Update.ChatMember.From
	case ctx.Update.ChatJoinRequest != nil:
		return &ctx.Update.ChatJoinRequest.From
	default:
		return nil
	}
}

func (ctx *Context) Chat() *illuminate.Chat {
	if m := ctx.Message(); m != nil {
		return &m.Chat
	} else if ctx.Update.MyChatMember != nil {
		return &ctx.Update.MyChatMember.Chat
	} else if ctx.Update.ChatMember != nil {
		return &ctx.Update.ChatMember.Chat
	} else if ctx.Update.ChatJoinRequest != nil {
		return &ctx.Update.ChatJoinRequest.Chat
	} else {
		return nil
	}
}

func (ctx *Context) ChatID() illuminate.PeerID {
	if c := ctx.Chat(); c != nil {
		return c.ID
	}

	if s := ctx.Sender(); s != nil {
		return s.ID
	}

	// impossible
	return illuminate.ChatID(0)
}

func (ctx *Context) CommandArgs() []string {
	if ctx.Update.Message == nil {
		return nil
	}
	args := strings.Split(ctx.Update.Message.Text, " ")
	if len(args) > 1 {
		return args[1:]
	}
	return nil
}

// HELPER FUNCTIONS

func (ctx *Context) Reply(text string, opts ...*illuminate.SendMessageOpts) (*illuminate.Message, error) {
	var opt *illuminate.SendMessageOpts
	if len(opts) > 0 {
		opt = opts[0]
	}
	return ctx.Bot.SendMessage(ctx.Context, ctx.Update.Message.Chat.ID, text, opt)
}
