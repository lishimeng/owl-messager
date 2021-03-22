package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
)

func GetMailByMessageId(msgId int) (m model.MailMessageInfo,err error) {
	return
}

func CreateMailMessage(message model.MessageInfo, sender model.MailSenderInfo, template model.MailTemplateInfo,
	templateParams string,
	subject, body, receiver, cc string) (m model.MailMessageInfo, err error) {

	m.MessageId = message.Id
	m.Template = template.Id
	m.Params = templateParams
	m.Sender = sender.Id
	m.Subject = subject
	m.Body = body
	m.Receivers = receiver
	if len(cc) > 0 {
		m.Cc = cc
	}

	_, err = app.GetOrm().Context.Insert(&m)
	return
}
