package illuminate

import "time"

type GetUpdatesChanOpts struct {
	*GetUpdatesOpts
	Buffer int
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
				time.Sleep(time.Second * 3)
				continue
			}

			for _, update := range updates {
				if update.UpdateID >= opts.GetUpdatesOpts.Offset {
					opts.GetUpdatesOpts.Offset = update.UpdateID + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}
