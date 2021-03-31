package main

import (
	"github.com/absaleb/notifyer/internal/model"
	"github.com/absaleb/notifyer/pkg/service/notifier"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.Info("starting...")
	message := &notifier.Message{
		Text:      "[ERROR] alert!!!",
		CreatedAt: time.Now(),
		Channels:  []*model.MailingChannel{
			&model.MailingChannel{
				Type: model.TELEGRAM,
			},
		},
	}

	message.SendMessage()

}
