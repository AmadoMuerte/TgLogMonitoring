package tgLogMonitoring

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	botToken string
	chatsId  []string
}

type LogMonitoring struct {
	cfg Config
}

func (l *LogMonitoring) NewLog(message string) {
	tgUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", l.cfg.botToken)

	for _, chatID := range l.cfg.chatsId {
		message = time.Now().Format("2006-01-02 15:04:05") + "\n" + url.QueryEscape(message)
		messageLength := len(message)
		if messageLength > 4096 {
			message = message[:4096]
		}
		resp, err := http.PostForm(tgUrl, url.Values{"chat_id": {chatID}, "text": {}})
		if err != nil {
			fmt.Println("Error sending request:", err)
			continue
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Error closing response body:", err)
			}
		}(resp.Body)
	}
}
