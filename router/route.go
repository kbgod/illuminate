package router

type Route struct {
	name     string
	filter   RouteFilter
	state    *string
	handlers []Handler
}

func (route *Route) Name(name string) *Route {
	route.name = name
	return route
}

func (route *Route) GetName() string {
	return route.name
}

func (route *Route) GetState() *string {
	return route.state
}

func (route *Route) GetFormattedState() string {
	if route.state == nil {
		return "<nil>"
	}

	return *route.state
}

func (route *Route) GetHandlersCount() int {
	return len(route.handlers)
}

func newRoute(filter RouteFilter, state *string, handler ...Handler) *Route {
	return &Route{
		filter:   filter,
		state:    state,
		handlers: handler,
	}
}
