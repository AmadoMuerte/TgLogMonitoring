## package for send messages in telegram chats

### init
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

```
`send message`
```go

tgLogMonitoring.NewLog(cfg, "Hello world!")

```
