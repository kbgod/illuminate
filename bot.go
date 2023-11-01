package illuminate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kbgod/illuminate/log"
	"github.com/kbgod/illuminate/log/adapter/std"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

const DefaultAPIHost = "https://api.telegram.org"

type BotOption func(*Bot)

type HttpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func WithAPIHost(host string) BotOption {
	return func(b *Bot) {
		b.apiHost = host
	}
}

func WithHttpDoer(client HttpDoer) BotOption {
	return func(b *Bot) {
		b.client = client
	}
}

func WithToken(token string) BotOption {
	return func(b *Bot) {
		b.token = token
	}
}

func WithDebugRequests(debug bool) BotOption {
	return func(b *Bot) {
		b.debugRequests = debug
	}
}

func WithLogger(log log.Logger) BotOption {
	return func(b *Bot) {
		b.log = log
	}
}

type Bot struct {
	apiHost string
	token   string
	client  HttpDoer

	debugRequests bool
	log           log.Logger

	Info *User
}

func NewBot(opts ...BotOption) *Bot {
	bot := &Bot{
		apiHost: DefaultAPIHost,
		client:  http.DefaultClient,
		log:     std.NewAdapter(std.LevelDebug),
	}
	for _, opt := range opts {
		opt(bot)
	}

	return bot
}

func (bot *Bot) CallMethod(
	ctx context.Context,
	method string,
	params map[string]string,
	data map[string]NamedReader,
) (json.RawMessage, error) {
	b := &bytes.Buffer{}

	var contentType string
	// Check if there are any files to upload. If yes, use multipart; else, use JSON.
	if len(data) > 0 {
		var err error
		contentType, err = fillBuffer(b, params, data)
		if err != nil {
			return nil, fmt.Errorf("failed to fill buffer with parameters and file data: %w", err)
		}
	} else {
		contentType = "application/json"
		err := json.NewEncoder(b).Encode(params)
		if err != nil {
			return nil, fmt.Errorf("failed to encode parameters as JSON: %w", err)
		}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, bot.apiHost+"/bot"+bot.token+"/"+method, b)
	if err != nil {
		return nil, fmt.Errorf("failed to build POST request to %s: %w", method, err)
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute POST request to %s: %w", method, err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var r Response
	if err = json.Unmarshal(respBody, &r); err != nil {
		return nil, fmt.Errorf(
			"failed to decode body [%s] of POST request to %s: %w", string(respBody), method, err,
		)
	}

	if bot.debugRequests {
		fields := map[string]any{
			"method": method,
			"params": params,
		}
		if !r.Ok {
			fields["error_code"] = r.ErrorCode
			fields["description"] = r.Description
			fields["parameters"] = r.Parameters
		} else {
			fields["result"] = string(r.Result)
		}
		bot.log.Debug("call method", fields)
	}
	if !r.Ok {
		return nil, &TelegramError{
			Method:         method,
			Params:         params,
			Code:           r.ErrorCode,
			Description:    r.Description,
			ResponseParams: r.Parameters,
		}
	}
	return r.Result, nil
}

type GetUpdatesChanConfig struct {
	bufferSize        int
	getUpdatesOptions *GetUpdatesOpts
}
type GetUpdatesChanOption func(*GetUpdatesChanConfig)

func WithGetUpdatesChanBufferSize(bufferSize int) GetUpdatesChanOption {
	return func(cfg *GetUpdatesChanConfig) {
		cfg.bufferSize = bufferSize
	}
}

func WithGetUpdatesOpts(getUpdatesOptions *GetUpdatesOpts) GetUpdatesChanOption {
	return func(cfg *GetUpdatesChanConfig) {
		cfg.getUpdatesOptions = getUpdatesOptions
	}
}

// GetUpdatesChan starts and returns a channel for getting updates.
func (bot *Bot) GetUpdatesChan(ctx context.Context, opts ...GetUpdatesChanOption) chan Update {
	cfg := &GetUpdatesChanConfig{
		bufferSize: 100,
		getUpdatesOptions: &GetUpdatesOpts{
			Timeout: 600,
		},
	}
	for _, opt := range opts {
		opt(cfg)
	}
	ch := make(chan Update, cfg.bufferSize)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
			}

			updates, err := bot.GetUpdates(ctx, cfg.getUpdatesOptions)
			if err != nil {
				bot.log.Error(err, "failed to get updates, retrying in 3 seconds...", nil)
				time.Sleep(time.Second * 3)

				continue
			}

			for _, update := range updates {
				if update.UpdateID >= cfg.getUpdatesOptions.Offset {
					cfg.getUpdatesOptions.Offset = update.UpdateID + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

func fillBuffer(b *bytes.Buffer, params map[string]string, data map[string]NamedReader) (string, error) {
	w := multipart.NewWriter(b)

	for k, v := range params {
		err := w.WriteField(k, v)
		if err != nil {
			return "", fmt.Errorf("failed to write multipart field %s with value %s: %w", k, v, err)
		}
	}

	for field, file := range data {
		fileName := file.Name()
		if fileName == "" {
			fileName = field
		}

		part, err := w.CreateFormFile(field, fileName)
		if err != nil {
			return "", fmt.Errorf("failed to create form file for field %s and fileName %s: %w", field, fileName, err)
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return "", fmt.Errorf("failed to copy file contents of field %s to form: %w", field, err)
		}
	}

	if err := w.Close(); err != nil {
		return "", fmt.Errorf("failed to close multipart form writer: %w", err)
	}

	return w.FormDataContentType(), nil
}
