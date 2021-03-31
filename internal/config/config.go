package config

import "github.com/absaleb/notifyer/internal/common"

var(
	TelegramKey = common.GetEnv("NOTIFIER-TELEGRAM-API-KEY","")
	TelegramBotChatID = common.GetEnv("NOTIFIER-TELEGRAM-CHAT-ID","")
)

