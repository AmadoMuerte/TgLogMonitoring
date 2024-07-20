package tgLogMonitoring

import (
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	BotToken string
	ChatsId  []string
}

func NewLog(cfg Config, message string) {
	tgUrl := "https://api.telegram.org/bot" + cfg.BotToken + "/sendMessage"

	for i := 0; i < len(cfg.ChatsId); i++ {
		req, _ := http.NewRequest("POST", tgUrl, nil)
		q := req.URL.Query()
		q.Add("chat_id", cfg.ChatsId[i])
		q.Add("text", message)
		req.URL.RawQuery = q.Encode()

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Error closing response body:", err)
			}
		}(resp.Body)
	}
}
