package notifier

import (
	"errors"
	"github.com/absaleb/notifyer/internal/model"
	"github.com/absaleb/notifyer/pkg/service/telegram"
	log "github.com/sirupsen/logrus"
	"time"
)

type Message struct {
	Text      string
	CreatedAt time.Time
	Channels  []*model.MailingChannel
}

func (message *Message) SendMessage() {
	if message == nil {
		log.Warningln("nil message")
		return
	}

	for _, ch := range message.Channels {
		switch ch.Type {
		case model.TELEGRAM:
			t, err := telegram.SendTelegram(message.Text)
			if err != nil {
				log.Errorln(err)
			}
			ch.SentAt = t
		default:
			log.Errorln(errors.New("Unsupported channel type"))
		}
	}
}


