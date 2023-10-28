package plugin

import "github.com/kbgod/illuminate"

type Menu struct {
	illuminate.ReplyKeyboardMarkup
	rowIndex int
}

func WithMenuKeyboardResize(resize bool) MenuOption {
	return func(menu *Menu) {
		menu.ResizeKeyboard = resize
	}
}

type MenuOption func(*Menu)

func NewMenu(options ...MenuOption) *Menu {
	menu := &Menu{
		ReplyKeyboardMarkup: illuminate.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard:       make([][]illuminate.KeyboardButton, 1),
		},
	}
	for _, option := range options {
		option(menu)
	}
	return menu
}

func (m *Menu) Row(buttons ...illuminate.KeyboardButton) *Menu {
	if len(m.Keyboard[m.rowIndex]) == 0 {
		m.Keyboard[m.rowIndex] = buttons
	} else {
		m.Keyboard = append(m.Keyboard, buttons)
		m.rowIndex++
	}

	return m
}

func (m *Menu) TextRow(buttons ...string) *Menu {
	keyboardButtons := make([]illuminate.KeyboardButton, 0, len(buttons))
	for _, button := range buttons {
		keyboardButtons = append(keyboardButtons, illuminate.KeyboardButton{
			Text: button,
		})
	}
	m.Row(keyboardButtons...)

	return m
}

func (m *Menu) Fill(perLine int, buttons ...illuminate.KeyboardButton) *Menu {
	for i := 0; i < len(buttons); i += perLine {
		end := i + perLine
		if end > len(buttons) {
			end = len(buttons)
		}
		m.Row(buttons[i:end]...)
	}
	return m
}

func (m *Menu) TextFill(perLine int, buttons ...string) *Menu {
	keyboardButtons := make([]illuminate.KeyboardButton, 0, len(buttons))
	for _, button := range buttons {
		keyboardButtons = append(keyboardButtons, illuminate.KeyboardButton{
			Text: button,
		})
	}
	m.Fill(perLine, keyboardButtons...)
	return m
}

func (m *Menu) Btn(btn illuminate.KeyboardButton) *Menu {
	m.Keyboard[m.rowIndex] = append(m.Keyboard[m.rowIndex], btn)
	return m
}

func (m *Menu) TextBtn(text string) *Menu {
	m.Keyboard[m.rowIndex] = append(m.Keyboard[m.rowIndex], illuminate.KeyboardButton{
		Text: text,
	})
	return m
}

func (m *Menu) ContactBtn(text string) *Menu {
	m.Keyboard[m.rowIndex] = append(m.Keyboard[m.rowIndex], illuminate.KeyboardButton{
		Text:           text,
		RequestContact: true,
	})
	return m
}

func (m *Menu) LocationBtn(text string) *Menu {
	m.Keyboard[m.rowIndex] = append(m.Keyboard[m.rowIndex], illuminate.KeyboardButton{
		Text:            text,
		RequestLocation: true,
	})
	return m
}

func (m *Menu) WebAppBtn(text, url string) *Menu {
	m.Keyboard[m.rowIndex] = append(m.Keyboard[m.rowIndex], illuminate.KeyboardButton{
		Text: text,
		WebApp: &illuminate.WebAppInfo{
			Url: url,
		},
	})
	return m
}

func (m *Menu) RequestChatBtn(text string, req *illuminate.KeyboardButtonRequestChat) *Menu {
	m.Keyboard[m.rowIndex] = append(m.Keyboard[m.rowIndex], illuminate.KeyboardButton{
		Text:        text,
		RequestChat: req,
	})
	return m
}

func (m *Menu) RequestUserBtn(text string, req *illuminate.KeyboardButtonRequestUser) *Menu {
	m.Keyboard[m.rowIndex] = append(m.Keyboard[m.rowIndex], illuminate.KeyboardButton{
		Text:        text,
		RequestUser: req,
	})
	return m
}

type InlineMenu struct {
	illuminate.InlineKeyboardMarkup
	rowIndex int
}

func NewInlineMenu() *InlineMenu {
	menu := &InlineMenu{
		InlineKeyboardMarkup: illuminate.InlineKeyboardMarkup{
			InlineKeyboard: make([][]illuminate.InlineKeyboardButton, 1),
		},
	}

	return menu
}

func (m *InlineMenu) Row(buttons ...illuminate.InlineKeyboardButton) *InlineMenu {
	if len(m.InlineKeyboard[m.rowIndex]) == 0 {
		m.InlineKeyboard[m.rowIndex] = buttons
	} else {
		m.InlineKeyboard = append(m.InlineKeyboard, buttons)
		m.rowIndex++
	}

	return m
}

func (m *InlineMenu) Fill(perLine int, buttons ...illuminate.InlineKeyboardButton) *InlineMenu {
	for i := 0; i < len(buttons); i += perLine {
		end := i + perLine
		if end > len(buttons) {
			end = len(buttons)
		}
		m.Row(buttons[i:end]...)
	}
	return m
}

func (m *InlineMenu) Btn(btn illuminate.InlineKeyboardButton) *InlineMenu {
	m.InlineKeyboard[m.rowIndex] = append(m.InlineKeyboard[m.rowIndex], btn)
	return m
}

func (m *InlineMenu) CallbackBtn(text, data string) *InlineMenu {
	m.InlineKeyboard[m.rowIndex] = append(m.InlineKeyboard[m.rowIndex], illuminate.InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	})
	return m
}

func (m *InlineMenu) URLBtn(text, url string) *InlineMenu {
	m.InlineKeyboard[m.rowIndex] = append(m.InlineKeyboard[m.rowIndex], illuminate.InlineKeyboardButton{
		Text: text,
		Url:  url,
	})
	return m
}

func (m *InlineMenu) LoginBtn(text, loginURL string) *InlineMenu {
	m.InlineKeyboard[m.rowIndex] = append(m.InlineKeyboard[m.rowIndex], illuminate.InlineKeyboardButton{
		Text: text,
		LoginUrl: &illuminate.LoginUrl{
			Url: loginURL,
		},
	})
	return m
}

func (m *InlineMenu) SwitchInlineQueryBtn(text, query string) *InlineMenu {
	m.InlineKeyboard[m.rowIndex] = append(m.InlineKeyboard[m.rowIndex], illuminate.InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: &query,
	})
	return m
}

func (m *InlineMenu) SwitchInlineCurrentChatBtn(text, query string) *InlineMenu {
	m.InlineKeyboard[m.rowIndex] = append(m.InlineKeyboard[m.rowIndex], illuminate.InlineKeyboardButton{
		Text:                         text,
		SwitchInlineQueryCurrentChat: &query,
	})
	return m
}

func (m *InlineMenu) SwitchInlineChosenChatBtn(
	text string, query *illuminate.SwitchInlineQueryChosenChat,
) *InlineMenu {
	m.InlineKeyboard[m.rowIndex] = append(m.InlineKeyboard[m.rowIndex], illuminate.InlineKeyboardButton{
		Text:                        text,
		SwitchInlineQueryChosenChat: query,
	})
	return m
}

func (m *InlineMenu) GameBtn(text string) *InlineMenu {
	m.InlineKeyboard[m.rowIndex] = append(m.InlineKeyboard[m.rowIndex], illuminate.InlineKeyboardButton{
		Text:         text,
		CallbackGame: &illuminate.CallbackGame{},
	})
	return m
}

func (m *InlineMenu) PayBtn(text string) *InlineMenu {
	m.InlineKeyboard[m.rowIndex] = append(m.InlineKeyboard[m.rowIndex], illuminate.InlineKeyboardButton{
		Text: text,
		Pay:  true,
	})
	return m
}

func CallbackBtn(text, data string) illuminate.InlineKeyboardButton {
	return illuminate.InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	}
}
