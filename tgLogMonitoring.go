package tgLogMonitoring

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	BotToken string
	ChatsId  []string
}

type New struct {
	Cfg Config
}

func (l *New) Send(message string) {
	tgUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", l.Cfg.BotToken)

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	for _, chatID := range l.Cfg.ChatsId {
		message = currentTime + "\n" + url.QueryEscape(message)
		messageLength := len(message)
		if messageLength > 4096 {
			message = message[:4096]
		}
		values := url.Values{"chat_id": {chatID}, "text": {message}}
		resp, err := http.PostForm(tgUrl, values)
		if err != nil {
			fmt.Println("Error sending request:", err)
			continue
		}
		defer func(body io.ReadCloser) {
			err := body.Close()
			if err != nil {
				fmt.Println("Error closing response body:", err)
			}
		}(resp.Body)
	}
}
