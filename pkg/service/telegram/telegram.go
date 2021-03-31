package telegram

import (
	"errors"
	"fmt"
	"github.com/absaleb/notifyer/internal/config"
	"net/http"
	"time"
)

func SendTelegram(text string) (*time.Time, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", config.TelegramKey, config.TelegramBotChatID, text)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		now := time.Now()
		return &now, nil
	}
	errString := fmt.Sprintf("error code %s", resp.Status)
	return nil, errors.New(errString)
}
