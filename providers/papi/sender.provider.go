package papi

type Sender interface {
	Id() string
}

// MailSender
// send mail
type MailSender interface {
	Sender
}

// SmsSender

type SmsSender interface {
	Sender
}

// ApnsSender

type ApnsSender interface {
	Sender
}

// WxSender

type WxSender interface {
	Sender
}
