package main

import (
	"context"
	"github.com/kbgod/illuminate"
	"github.com/kbgod/illuminate/plugin"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	bot := illuminate.NewBot(illuminate.WithToken(os.Getenv("BOT_TOKEN")))

	var lastUpdateID int64
	for {
		updates, err := bot.GetUpdates(ctx, &illuminate.GetUpdatesOpts{
			Offset: lastUpdateID + 1,
		})
		if err != nil {
			log.Println(err)
		}
		for _, upd := range updates {
			lastUpdateID = upd.UpdateID
			go handleUpdate(ctx, bot, upd)
		}

	}
}

func handleUpdate(ctx context.Context, bot *illuminate.Bot, upd illuminate.Update) {
	menu := plugin.NewInlineMenu()
	btns := []illuminate.InlineKeyboardButton{
		plugin.CallbackBtn("1", "1"),
		plugin.CallbackBtn("2", "2"),
		plugin.CallbackBtn("3", "3"),
		plugin.CallbackBtn("1", "1"),
		plugin.CallbackBtn("2", "2"),
		plugin.CallbackBtn("3", "3"),
		plugin.CallbackBtn("1", "1"),
		plugin.CallbackBtn("2", "2"),
		plugin.CallbackBtn("3", "3"),
		plugin.CallbackBtn("1", "1"),
		plugin.CallbackBtn("2", "2"),
		plugin.CallbackBtn("3", "3"),
	}
	menu.Fill(3, btns...)
	_, err := bot.SendMessage(ctx, upd.Message.Chat.ID, upd.Message.Text, &illuminate.SendMessageOpts{
		ReplyMarkup: menu,
	})
	if err != nil {
		log.Println(err)
	}
}
