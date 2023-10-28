# Golang Telegram Bot Fremework
[![Test](https://github.com/kbgod/illuminate/actions/workflows/test.yml/badge.svg)](https://github.com/kbgod/illuminate/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/kbgod/illuminate/graph/badge.svg?token=VHJJZGTWUI)](https://codecov.io/gh/kbgod/illuminate)
Based on [paulsonoflars/gotgbot](https://github.com/paulsonoflars/gotgbot) types generation and inspired by [mr-linch/go-tg](https://github.com/mr-linch/go-tg)

All the telegram types and methods are generated from
[a bot api spec](https://github.com/PaulSonOfLars/telegram-bot-api-spec). These are generated in the `gen_*.go` files.

## Features:

- All telegram API types and methods are generated from the bot api docs, which makes this library:
    - Guaranteed to match the docs
    - Easy to update
    - Self-documenting (Re-uses pre-existing telegram docs)
- Type safe; no weird interface{} logic, all types match the bot API docs.
- No third party library bloat; only uses standard library.
- Updates are each processed in their own go routine, encouraging concurrent processing, and keeping your bot
  responsive.
- Code panics are automatically recovered from and logged, avoiding unexpected downtime.

## Getting started

Download the library with the standard `go get` command:

```bash
go get github.com/kbgod/illuminate
```

### Example bots

*in progress...*

## Docs

Docs can be found [here](https://github.com/kbgod/illuminate/illuminate/v2).

## Contributing

*in progress...*
