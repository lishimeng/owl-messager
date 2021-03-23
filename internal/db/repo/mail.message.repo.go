package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
	"time"
)

func GetMailByMessageId(msgId int) (m model.MailMessageInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MailMessageInfo)).Filter("MessageId", msgId).One(&m)
	return
}

func CreateMailMessage(message model.MessageInfo, sender model.MailSenderInfo, template model.MailTemplateInfo,
	templateParams string,
	subject, receiver, cc string) (m model.MailMessageInfo, err error) {

	m.MessageId = message.Id
	m.Template = template.Id
	m.Params = templateParams
	m.Sender = sender.Id
	m.Subject = subject
	m.Receivers = receiver
	if len(cc) > 0 {
		m.Cc = cc
	}

	m.Status = model.MailTemplateEnable
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()

	_, err = app.GetOrm().Context.Insert(&m)
	return
}
