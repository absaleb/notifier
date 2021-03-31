// Code generated by "stringer -type=MailingChannelType"; DO NOT EDIT.

package model

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EMAIL-1]
	_ = x[SMS-2]
	_ = x[TELEGRAM-3]
}

const _MailingChannelType_name = "EMAILSMSTELEGRAM"

var _MailingChannelType_index = [...]uint8{0, 5, 8, 16}

func (i MailingChannelType) String() string {
	i -= 1
	if i < 0 || i >= MailingChannelType(len(_MailingChannelType_index)-1) {
		return "MailingChannelType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _MailingChannelType_name[_MailingChannelType_index[i]:_MailingChannelType_index[i+1]]
}
