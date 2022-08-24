package mail

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/messager"
	"github.com/lishimeng/owl/internal/provider/mail/smtp"
)

type smtpSender struct {
	proxy *smtp.MailSmtpProvider
}

func NewSmtp(config string) (s messager.MailProvider, err error) {

	var p *smtp.MailSmtpProvider
	var h = smtpSender{}
	p, err = smtp.New(config)
	h.proxy = p
	s = &h
	return
}

func (s *smtpSender) Send(subject string, body string, to ...string) (err error) {

	log.Debug("mail body:%s", body)
	err = s.proxy.Send(subject, body, to...)
	return
}
