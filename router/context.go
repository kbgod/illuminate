package router

import (
	"context"
	"strings"

	"github.com/kbgod/illuminate"
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
		switch m := ctx.Update.CallbackQuery.Message.(type) {
		case illuminate.Message:
			return &m
		case *illuminate.Message:
			return m
		case illuminate.InaccessibleMessage:
			return &illuminate.Message{
				Chat:      m.GetChat(),
				MessageId: m.MessageId,
				Date:      m.Date,
			}
		case *illuminate.InaccessibleMessage:
			return &illuminate.Message{
				Chat:      m.GetChat(),
				MessageId: m.MessageId,
				Date:      m.Date,
			}
		}
		return nil
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

func (ctx *Context) ChatID() int64 {
	if c := ctx.Chat(); c != nil {
		return c.Id
	}

	if s := ctx.Sender(); s != nil {
		return s.Id
	}

	// impossible
	return 0
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
	return ctx.Bot.SendMessage(ctx.ChatID(), text, opt)
}

// ReplyVoid sends message without returning result
func (ctx *Context) ReplyVoid(text string, opts ...*illuminate.SendMessageOpts) error {
	_, err := ctx.Reply(text, opts...)
	return err
}

// ReplyWithMenu sends message with menu
func (ctx *Context) ReplyWithMenu(
	text string, menu illuminate.IMenu, opts ...*illuminate.SendMessageOpts,
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
	text string, menu illuminate.IMenu, opts ...*illuminate.SendMessageOpts,
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
	return ctx.Bot.AnswerCallbackQuery(ctx.Update.CallbackQuery.Id, opt)
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
	return ctx.Bot.DeleteMessage(ctx.ChatID(), ctx.Message().MessageId, opt)
}

// DeleteMessageVoid deletes message which is in update without returning result
func (ctx *Context) DeleteMessageVoid(opts ...*illuminate.DeleteMessageOpts) error {
	_, err := ctx.DeleteMessage(opts...)
	return err
}

func (ctx *Context) EditMessageText(text string, opts ...*illuminate.EditMessageTextOpts) (*illuminate.Message, bool, error) {
	if ctx.parseMode != nil {
		if len(opts) == 0 {
			opts = append(opts, &illuminate.EditMessageTextOpts{
				ParseMode: *ctx.parseMode,
			})
		} else {
			opts[0].ParseMode = *ctx.parseMode
		}
	}
	var opt *illuminate.EditMessageTextOpts
	if len(opts) > 0 {
		opt = opts[0]
		opt.ChatId = ctx.ChatID()
		opt.MessageId = ctx.Message().MessageId
	} else {
		opt = &illuminate.EditMessageTextOpts{
			ChatId:    ctx.ChatID(),
			MessageId: ctx.Message().MessageId,
		}
	}

	return ctx.Bot.EditMessageText(text, opt)
}

func (ctx *Context) EditMessageTextVoid(text string, opts ...*illuminate.EditMessageTextOpts) error {
	_, _, err := ctx.EditMessageText(text, opts...)
	return err
}

func (ctx *Context) ReplyEmojiReaction(emoji ...string) (bool, error) {
	reactions := make([]illuminate.ReactionType, 0, len(emoji))
	for _, e := range emoji {
		reactions = append(reactions, illuminate.ReactionTypeEmoji{Emoji: e})
	}
	return ctx.Bot.SetMessageReaction(ctx.ChatID(), ctx.Message().MessageId, &illuminate.SetMessageReactionOpts{
		Reaction: reactions,
	})
}

func (ctx *Context) ReplyEmojiReactionVoid(emoji ...string) error {
	_, err := ctx.ReplyEmojiReaction(emoji...)
	return err
}

func (ctx *Context) ReplyEmojiBigReaction(emoji ...string) (bool, error) {
	reactions := make([]illuminate.ReactionType, 0, len(emoji))
	for _, e := range emoji {
		reactions = append(reactions, illuminate.ReactionTypeEmoji{Emoji: e})
	}
	return ctx.Bot.SetMessageReaction(ctx.ChatID(), ctx.Message().MessageId, &illuminate.SetMessageReactionOpts{
		Reaction: reactions,
		IsBig:    true,
	})
}

func (ctx *Context) ReplyEmojiBigReactionVoid(emoji ...string) error {
	_, err := ctx.ReplyEmojiBigReaction(emoji...)
	return err
}
