package main

import (
	"context"
	"github.com/kbgod/illuminate"
	zerologAdapter "github.com/kbgod/illuminate/log/adapter/zerolog"
	"github.com/kbgod/illuminate/plugin"
	"github.com/kbgod/illuminate/router"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger = zerolog.New(
	zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stderr
		w.TimeFormat = time.RFC3339
	}),
).With().Timestamp().Logger()

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-exit
		cancel()
	}()
	logger.WithContext(ctx)
	bot := illuminate.NewBot(
		illuminate.WithToken(os.Getenv("BOT_TOKEN")),
		illuminate.WithLogger(zerologAdapter.NewAdapter(&logger)),
	)
	bofInfo, err := bot.GetMe(ctx, nil)
	if err != nil {
		panic(err)
	}
	logger.Info().Str("username", string(bofInfo.Username)).Msg("bot authorized successfully")

	app := router.New(bot)
	app.Use(plugin.RecoveryMiddleware(zerologAdapter.NewAdapter(&logger)))
	app.OnStart(func(ctx *router.Context) error {
		return plugin.Void(ctx.Reply("Hello!"))
	})
	app.OnCommand("fatal", func(ctx *router.Context) error {
		var a *int
		*a = 1

		return nil
	})
	app.OnMessage(func(ctx *router.Context) error {
		return plugin.Void(ctx.Reply("Undefined command!"))
	})

	updates := bot.GetUpdatesChan(ctx)
	runWorkerPool(ctx, 100, app, updates)

	<-ctx.Done()

	logger.Info().Str("username", string(bofInfo.Username)).Msg("bot stopped")
}

func runWorkerPool(ctx context.Context, size int, router *router.Router, updates <-chan illuminate.Update) {
	for i := 0; i < size; i++ {
		go func(id int) {
			for update := range updates {
				u := update
				_ = router.HandleUpdate(ctx, &u)
			}
		}(i)
	}
}
