package illuminate

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// GetLink is a helper method to easily get the message link (It will return an empty string in case of private or group chat type).
func (m Message) GetLink() string {
	if m.Chat.Type == "private" || m.Chat.Type == "group" {
		return ""
	}
	if m.Chat.Username != "" {
		return fmt.Sprintf("https://t.me/%s/%d", m.Chat.Username, m.MessageId)
	}
	// Message links use raw chatIDs without the -100 prefix; this trims that prefix.
	rawChatID := strings.TrimPrefix(strconv.FormatInt(m.Chat.Id, 10), "-100")
	return fmt.Sprintf("https://t.me/c/%s/%d", rawChatID, m.MessageId)
}

// Reply is a helper function to easily call Bot.SendMessage as a reply to an existing message.
func (m Message) Reply(b *Bot, text string, opts *SendMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &SendMessageOpts{}
	}

	if opts.ReplyParameters == nil || opts.ReplyParameters.MessageId == 0 {
		opts.ReplyParameters = &ReplyParameters{MessageId: m.MessageId}
	}

	return b.SendMessage(m.Chat.Id, text, opts)
}

// SendMessage is a helper function to easily call Bot.SendMessage in a chat.
func (c Chat) SendMessage(b *Bot, text string, opts *SendMessageOpts) (*Message, error) {
	return b.SendMessage(c.Id, text, opts)
}

// Unban is a helper function to easily call Bot.UnbanChatMember in a chat.
func (c Chat) Unban(b *Bot, userID int64, opts *UnbanChatMemberOpts) (bool, error) {
	return b.UnbanChatMember(c.Id, userID, opts)
}

// Promote is a helper function to easily call Bot.PromoteChatMember in a chat.
func (c Chat) Promote(b *Bot, userID int64, opts *PromoteChatMemberOpts) (bool, error) {
	return b.PromoteChatMember(c.Id, userID, opts)
}

// URL gets the URL the file can be downloaded from.
func (f File) URL(b *Bot, opts *RequestOpts) string {
	return b.FileURL(b.Token, f.FilePath, opts)
}

// unmarshalMaybeInaccessibleMessage is a JSON unmarshal helper to marshal the right structs into a
// MaybeInaccessibleMessage interface based on the Date field.
// This method is manually maintained due to special-case handling on the Date field rather than a specific type field.
func unmarshalMaybeInaccessibleMessage(d json.RawMessage) (MaybeInaccessibleMessage, error) {
	if len(d) == 0 {
		return nil, nil
	}

	t := struct {
		Date int64
	}{}
	err := json.Unmarshal(d, &t)
	if err != nil {
		return nil, err
	}

	// As per the docs, date is always 0 for inaccessible messages:
	// https://core.telegram.org/bots/api#inaccessiblemessage
	if t.Date == 0 {
		s := InaccessibleMessage{}
		err := json.Unmarshal(d, &s)
		if err != nil {
			return nil, err
		}
		return s, nil
	}

	s := Message{}
	err = json.Unmarshal(d, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
