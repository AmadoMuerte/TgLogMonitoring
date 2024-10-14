package tgLogMonitoring

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Config struct {
	// telegram bot token
	BotToken string
	// telegram chats id
	ChatsId []string
	// show date now
	DateShow bool
}

type New struct {
	Cfg Config
}

func (l *New) Send(message interface{}, title string) {
	tgUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", l.Cfg.BotToken)

	dateNow := time.Now().Format("2006-01-02 15:04:05")
	messageStr := fmt.Sprintf("%v", message)
	var formattedMessage string

	switch l.Cfg.DateShow {
	case true:
		formattedMessage = fmt.Sprintf("<b>%s - %s</b>\n<pre language='bash'>%s</pre>", title, dateNow, messageStr)

	case false:
		formattedMessage = fmt.Sprintf("<b>%s</b>\n<pre language='bash'>%s</pre>", title, messageStr)
	}

	var wg sync.WaitGroup

	for _, chatID := range l.Cfg.ChatsId {
		wg.Add(1)
		go func(chatID string) {
			defer wg.Done()

			values := url.Values{
				"chat_id":    {chatID},
				"text":       {formattedMessage},
				"parse_mode": {"html"},
			}
			resp, err := http.PostForm(tgUrl, values)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()
		}(chatID)
	}

	wg.Wait()
}
