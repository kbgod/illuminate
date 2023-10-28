package router

import "testing"

func TestRoute_Name(t *testing.T) {
	route := new(Route)
	route.Name("test")

	if route.name != "test" {
		t.Errorf("route.Name() = %s; want test", route.name)
	}
}

func TestRoute_GetName(t *testing.T) {
	route := new(Route)
	route.Name("test")

	if route.GetName() != "test" {
		t.Errorf("route.GetName() = %s; want test", route.GetName())
	}
}

func TestRoute_GetState(t *testing.T) {
	route := new(Route)
	if route.state != nil {
		t.Errorf("route.GetState() = %v; want <nil>", route.GetState())
	}
	state := "test"
	route.state = &state

	if routeState := route.GetState(); routeState == nil {
		t.Errorf("route.GetState() = %v; want test", routeState)
	} else if *routeState != "test" {
		t.Errorf("route.GetState() = %s; want test", *routeState)
	}
}

func TestRoute_GetFormattedState(t *testing.T) {
	route := new(Route)
	if route.GetFormattedState() != "<nil>" {
		t.Errorf("route.GetFormattedState() = %s; want <nil>", route.GetFormattedState())
	}
	state := "test"
	route.state = &state

	if route.GetFormattedState() != "test" {
		t.Errorf("route.GetFormattedState() = %s; want test", route.GetFormattedState())
	}
}

func TestRoute_GetHandlersCount(t *testing.T) {
	route := new(Route)
	if route.GetHandlersCount() != 0 {
		t.Errorf("route.GetHandlersCount() = %d; want 0", route.GetHandlersCount())
	}
	route.handlers = append(route.handlers, func(ctx *Context) error {
		return nil
	})
	if route.GetHandlersCount() != 1 {
		t.Errorf("route.GetHandlersCount() = %d; want 1", route.GetHandlersCount())
	}
}

func TestNewRoute(t *testing.T) {
	route := newRoute(nil, nil)
	if route == nil {
		t.Errorf("newRoute(nil, nil) = <nil>; want not <nil>")
	}
	if route.filter != nil {
		t.Errorf("newRoute(nil, nil).filter = %v; want <nil>", route.filter)
	}
	if route.state != nil {
		t.Errorf("newRoute(nil, nil).state = %v; want <nil>", route.state)
	}
	if len(route.handlers) != 0 {
		t.Errorf("newRoute(nil, nil).handlers = %d; want 0", len(route.handlers))
	}

	route = newRoute(nil, nil, func(ctx *Context) error {
		return nil
	})
	if route == nil {
		t.Errorf("newRoute(nil, nil, func(ctx *Context) error { return nil }) = <nil>; want not <nil>")
	}
	if route.filter != nil {
		t.Errorf("newRoute(nil, nil, func(ctx *Context) error { return nil }).filter = %v; want <nil>", route.filter)
	}
	if route.state != nil {
		t.Errorf("newRoute(nil, nil, func(ctx *Context) error { return nil }).state = %v; want <nil>", route.state)
	}
	if len(route.handlers) != 1 {
		t.Errorf("newRoute(nil, nil, func(ctx *Context) error { return nil }).handlers = %d; want 1", len(route.handlers))
	}
}
