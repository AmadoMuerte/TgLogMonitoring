## package for send messages in telegram chats

`Import package`
```go

import tgLogMonitoring "github.com/AmadoMuerte/TgLogMonitoring"

```

`set your config`
```go
cfg := tgLogMonitoring.Config{
	BotToken: "token",
	ChatsId: []string{
		"chat1",
		"chat2",
	},
}

tgLog := tgLogMonitoring.New{Cfg: cfg}

```
`send message`
```go

tgLog.Send("hello world", "My test log")

```
