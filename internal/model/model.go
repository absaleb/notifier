//go:generate stringer -type=MailingChannelType
package model

import "time"

type MailingChannelType int

const (
	_ MailingChannelType = iota
	EMAIL
	SMS
	TELEGRAM
)

type MailingChannel struct {
	Type   MailingChannelType
	SentAt *time.Time
}

