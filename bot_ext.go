package illuminate

import (
	"time"
)

type GetUpdatesChanOpts struct {
	*GetUpdatesOpts
	Buffer       int
	ErrorHandler func(error)
}

func (bot *Bot) GetUpdatesChan(opts *GetUpdatesChanOpts) <-chan Update {
	defaultOpts := &GetUpdatesChanOpts{
		Buffer: 100,
		GetUpdatesOpts: &GetUpdatesOpts{
			Timeout: 600,
		},
	}
	if opts == nil {
		opts = defaultOpts
	}
	ch := make(chan Update, opts.Buffer)
	go func() {
		for {
			updates, err := bot.GetUpdates(opts.GetUpdatesOpts)
			if err != nil {
				if opts.ErrorHandler != nil {
					opts.ErrorHandler(err)
				}
				time.Sleep(time.Second * 3)
				continue
			}

			for _, update := range updates {
				if update.UpdateId >= opts.GetUpdatesOpts.Offset {
					opts.GetUpdatesOpts.Offset = update.UpdateId + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

func (bot *Bot) GetChannelAdministrators(username string, opts *GetChatAdministratorsOpts) ([]ChatMember, error) {
	v := map[string]string{}
	v["chat_id"] = username

	var reqOpts *RequestOpts
	if opts != nil {
		reqOpts = opts.RequestOpts
	}

	r, err := bot.Request("getChatAdministrators", v, nil, reqOpts)
	if err != nil {
		return nil, err
	}

	return unmarshalChatMemberArray(r)
}
