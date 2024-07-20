## package for send messages in telegram chats

### init

`set your config`
```go

cfg := tgLogMonitoring.config {
    BotToken: "yourBotToken",
	ChatsId: {
		"id1",
		"id2"
    }
}

```
`send message`
```go

tgLogMonitoring.NewLog(cfg, "hello world!")

```
