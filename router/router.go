package router

import (
	"context"
	"errors"
	"github.com/kbgod/illuminate"
)

var ErrRouteNotFound = errors.New("route not found")

type Handler func(*Context) error

type Router struct {
	state    *string
	parent   *Router
	bot      *illuminate.Bot
	routes   []*Route
	handlers []Handler
}

func New(bot *illuminate.Bot) *Router {
	return &Router{
		bot: bot,
	}
}

func (r *Router) next(ctx *Context) error {
	for ctx.indexRoute < len(r.routes)-1 {
		ctx.indexRoute++
		route := r.routes[ctx.indexRoute]
		if route.filter(ctx) && (route.state == nil || (ctx.state != nil && *route.state == *ctx.state)) {
			ctx.route = route
			ctx.indexHandler = -1
			return ctx.Next()
		}
	}

	return ErrRouteNotFound
}

func (r *Router) Use(middlewares ...Handler) {
	r.handlers = append(r.handlers, middlewares...)
}

func (r *Router) UseState(state string, handlers ...Handler) *Router {
	return &Router{
		parent:   r,
		state:    &state,
		bot:      r.bot,
		routes:   r.routes,
		handlers: handlers,
	}
}

func (r *Router) Group(handlers ...Handler) *Router {
	return &Router{
		parent:   r,
		state:    r.state,
		bot:      r.bot,
		routes:   r.routes,
		handlers: handlers,
	}
}

func (r *Router) GetRoutes() []*Route {
	return r.routes
}

func (r *Router) On(filter RouteFilter, handlers ...Handler) *Route {
	var route *Route
	if r.parent != nil {
		route = newRoute(filter, r.state, append(r.handlers, handlers...)...)
	} else {
		route = newRoute(filter, r.state, handlers...)
	}
	r.addRoute(route)
	return route
}

func (r *Router) addRoute(route *Route) {
	if r.parent != nil {
		r.parent.addRoute(route)
	} else {
		r.routes = append(r.routes, route)
	}
}

func (r *Router) HandleUpdate(ctx context.Context, update *illuminate.Update) error {
	return newContext(ctx, r, update).Next()
}
