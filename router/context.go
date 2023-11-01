package router

import (
	"context"
	"github.com/kbgod/illuminate"
	"github.com/kbgod/illuminate/plugin"
	"strings"
)

type Context struct {
	state        *string
	router       *Router
	route        *Route
	indexRoute   int
	indexHandler int

	parseMode *string
	Context   context.Context
	Update    *illuminate.Update
	Bot       *illuminate.Bot
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

func (ctx *Context) SetParseMode(parseMode string) {
	ctx.parseMode = &parseMode
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

// Reply sends message to the chat from update
func (ctx *Context) Reply(text string, opts ...*illuminate.SendMessageOpts) (*illuminate.Message, error) {
	if ctx.parseMode != nil {
		if len(opts) == 0 {
			opts = append(opts, &illuminate.SendMessageOpts{
				ParseMode: *ctx.parseMode,
			})
		} else {
			opts[0].ParseMode = *ctx.parseMode
		}
	}
	var opt *illuminate.SendMessageOpts
	if len(opts) > 0 {
		opt = opts[0]
	}
	return ctx.Bot.SendMessage(ctx.Context, ctx.ChatID(), text, opt)
}

// ReplyVoid sends message without returning result
func (ctx *Context) ReplyVoid(text string, opts ...*illuminate.SendMessageOpts) error {
	_, err := ctx.Reply(text, opts...)
	return err
}

// ReplyWithMenu sends message with menu
func (ctx *Context) ReplyWithMenu(
	text string, menu plugin.IMenu, opts ...*illuminate.SendMessageOpts,
) (*illuminate.Message, error) {
	if len(opts) == 0 {
		opts = append(opts, &illuminate.SendMessageOpts{

			ReplyMarkup: menu.Unwrap(),
		})
	}
	return ctx.Reply(text, opts...)
}

// ReplyWithMenuVoid sends message with menu without returning result
func (ctx *Context) ReplyWithMenuVoid(
	text string, menu plugin.IMenu, opts ...*illuminate.SendMessageOpts,
) error {
	_, err := ctx.ReplyWithMenu(text, menu, opts...)
	return err
}

// Answer sends answer to callback query from update
func (ctx *Context) Answer(text string, opts ...*illuminate.AnswerCallbackQueryOpts) (bool, error) {
	if text != "" {
		if len(opts) == 0 {
			opts = append(opts, &illuminate.AnswerCallbackQueryOpts{
				Text: text,
			})
		} else {
			opts[0].Text = text
		}
	}
	var opt *illuminate.AnswerCallbackQueryOpts
	if len(opts) > 0 {
		opt = opts[0]
	}
	return ctx.Bot.AnswerCallbackQuery(ctx.Context, ctx.Update.CallbackQuery.ID, opt)
}

// AnswerVoid sends answer to callback query without returning result
func (ctx *Context) AnswerVoid(text string, opts ...*illuminate.AnswerCallbackQueryOpts) error {
	_, err := ctx.Answer(text, opts...)
	return err
}

// AnswerAlert sends answer to callback query from update with alert
func (ctx *Context) AnswerAlert(text string, opts ...*illuminate.AnswerCallbackQueryOpts) (bool, error) {
	if len(opts) == 0 {
		opts = append(opts, &illuminate.AnswerCallbackQueryOpts{
			ShowAlert: true,
		})
	} else {
		opts[0].ShowAlert = true
	}
	return ctx.Answer(text, opts...)
}

// AnswerAlertVoid sends answer to callback query with alert without returning result
func (ctx *Context) AnswerAlertVoid(text string, opts ...*illuminate.AnswerCallbackQueryOpts) error {
	_, err := ctx.AnswerAlert(text, opts...)
	return err
}

// DeleteMessage deletes message which is in update
func (ctx *Context) DeleteMessage(opts ...*illuminate.DeleteMessageOpts) (bool, error) {
	var opt *illuminate.DeleteMessageOpts
	if len(opts) > 0 {
		opt = opts[0]
	}
	return ctx.Bot.DeleteMessage(ctx.Context, ctx.ChatID(), ctx.Message().MessageID, opt)
}
