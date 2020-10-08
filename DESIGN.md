# Desgin

- [Desgin](#desgin)
  - [Requirements](#requirements)
  - [Inspiration](#inspiration)
  - [Packages](#packages)
  - [Client](#client)
    - [Do](#do)
    - [GetMe](#getme)
    - [SendText](#sendtext)
  - [PromoteChatMember](#promotechatmember)
  - [SetMyCommands](#setmycommands)

## Requirements
 - Use [context](https://golang.org/pkg/context/) for **cancelation** and **deadlines** in IO.
 - Go way

## Inspiration

 - https://github.com/aiogram/aiogram
 - https://github.com/bwmarrin/discordgo
 - https://github.com/google/go-github
 - https://github.com/slack-go/slack
 - https://github.com/centrifugal/centrifuge

## Packages

| Package | Path                          | Description                                                                       |
| :------ | :---------------------------- | :-------------------------------------------------------------------------------- |
| `tg`    | github.com/bots-house/go-tg     | Contains `Client` with methods and types used for sending and receiving messages. |
| `tgr`   | github.com/bots-house/go-tg/tgr | Contains `Bot` and routing related types                                          |

## Client

Construct a new Telegram Bot API client using token from [@BotFather](https://t.me/BotFather).

```go
client := tg.NewClient("12345:secret")
```

Custumize client using options, e.g. use proxy for transport.

```go
// create http.Client with proxy
doer := &http.Client{
  Transport: &http.Transport{
    Proxy: http.ProxyURL("socks5://127.0.0.1:9150")
  }
}

// create Telegram Bot API client with custom doer.
client := tg.NewClient("12345:secret", tg.WithDoer(doer))
```

### Do

`Do` is method to execute any request to Telegram Bot API. It's low level.

```go
// open local file for send
photo, err := tg.NewInputFileLocal("testdata/logo.png")
if err != nil {
  // ...
}
defer photo.Close()

// create request
req := tg.NewRequest("sendPhoto")
req.AddPeer("chat_id", tg.Username("go_tg"))
req.AddFile("photo", photo)
req.AddString("caption", "this is our logo")
req.AddParseMode(tg.Markdown)

if err := client.Do(ctx, req, nil); err != nil {
  // ...
}
```

### [GetMe](https://core.telegram.org/bots/api#getMe)

```go
me, err := client.GetMe(ctx)
if err != nil {
  // ...
}
```

### [SendText](https://core.telegram.org/bots/api#sendMessage)

```go
opts := &tg.SendTextOpts{
  ParseMode: tg.Markdown,
  DisableNotification: true,
}

msg, err := client.SendText(
  ctx,
  tg.Username("MrLinch"),
  "how are u?",
  opts,
)
```

## [PromoteChatMember](https://core.telegram.org/bots/api#promoteChatMember)

```go
opts := &tg.PromoteChatMemberOpts{
  CanChangeInfo: true,
  CanPostMessages: true,
}

if err := client.PromoteChatMember(ctx,
  tg.Username("oh_my_channely"),
  tg.UserID(12341920),
  opts,
); err != nil {
  // ...
}
```

## [SetMyCommands](https://core.telegram.org/bots/api#setMyCommands)

```go
if err := client.SetMyCommands(ctx, []tg.BotCommand{
  {"start", "get current stats"},
}); err != nil {
  // ...
}
```
