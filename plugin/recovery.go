package plugin

import (
	"github.com/kbgod/illuminate/log"
	"github.com/kbgod/illuminate/router"
	"runtime/debug"
)

func RecoveryMiddleware(log log.Logger) router.Handler {
	return func(ctx *router.Context) error {
		defer func() {
			if r := recover(); r != nil {
				log.Error(nil, "fatal error", map[string]any{
					"panic": r,
				})
				debug.PrintStack()
			}
		}()
		return ctx.Next()
	}
}
