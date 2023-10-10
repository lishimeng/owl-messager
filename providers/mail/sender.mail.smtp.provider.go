package mail

import (
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/providers/mail/smtp"
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

func (s *smtpSender) Send(req messager.MailRequest) (err error) {

	content, err := buildMailBody(req.Template, req.Params)
	if err != nil {
		return
	}
	err = s.proxy.Send(req.Subject, content, req.Receivers...)
	return
}
