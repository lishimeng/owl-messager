package msg

type MessageCategory string

func (mc MessageCategory) String() string {
	return string(mc)
}

const (
	MailMessage MessageCategory = "mail"
	SmsMessage  MessageCategory = "sms"
	ApnsMessage MessageCategory = "apns"
)

var MessageCategories map[MessageCategory]byte

func init() {
	MessageCategories = make(map[MessageCategory]byte)
	MessageCategories[MailMessage] = 1
	MessageCategories[SmsMessage] = 1
	MessageCategories[ApnsMessage] = 1
}

func IsValidCategory(category MessageCategory) (valid bool) {
	_, valid = MessageCategories[category]
	return
}
