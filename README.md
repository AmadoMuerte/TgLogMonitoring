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
	DateShow: false,
}
tgLog := tgLogMonitoring.New{Cfg: cfg}

```
`send your log or messsages`
```go

tgLog.Send("hello world", "My test log")

mySlice := []string{"1", "2", "3", "4"}
tgLog.Send(mySlice, "My Title")

```
