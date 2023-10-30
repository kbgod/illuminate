package router

import "strings"

type RouteFilter func(*Context) bool

func Command(command string) RouteFilter {
	return func(ctx *Context) bool {
		if ctx.Update.Message == nil {
			return false
		}
		return strings.HasPrefix(ctx.Update.Message.Text, "/"+command)
	}
}

func AnyUpdate() RouteFilter {
	return func(ctx *Context) bool {
		return true
	}
}

func Message() RouteFilter {
	return func(ctx *Context) bool {
		return ctx.Update.Message != nil
	}
}

func CommandWithAt(command, username string) RouteFilter {
	return func(ctx *Context) bool {
		if ctx.Update.Message == nil {
			return false
		}
		return strings.HasPrefix(ctx.Update.Message.Text, "/"+command+"@"+username)
	}
}

func TextContains(text string) RouteFilter {
	return func(ctx *Context) bool {
		if ctx.Update.Message == nil {
			return false
		}
		return strings.Contains(ctx.Update.Message.Text, text)
	}
}

func TextPrefix(text string) RouteFilter {
	return func(ctx *Context) bool {
		if ctx.Update.Message == nil {
			return false
		}
		return strings.HasPrefix(ctx.Update.Message.Text, text)
	}
}

func CallbackPrefix(text string) RouteFilter {
	return func(ctx *Context) bool {
		if ctx.Update.CallbackQuery == nil {
			return false
		}
		return strings.HasPrefix(ctx.Update.CallbackQuery.Data, text)
	}
}
