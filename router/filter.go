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
