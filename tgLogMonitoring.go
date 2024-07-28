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

func (l *New) Send(message string, title string) {
	tgUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", l.Cfg.BotToken)

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	formattedMessage := fmt.Sprintf("<strong>%s - %s</strong>\n<code>%s</code>", title, currentTime, message)

	for _, chatID := range l.Cfg.ChatsId {
		values := url.Values{
			"chat_id":    {chatID},
			"text":       {formattedMessage},
			"parse_mode": {"html"},
		}
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
