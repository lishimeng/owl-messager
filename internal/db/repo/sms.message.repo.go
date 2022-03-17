package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
	"time"
)

func GetSmsByMessageId(msgId int) (m model.SmsMessageInfo, err error) {
	return
}

func CreateSmsMessage(message model.MessageInfo, sender model.SmsSenderInfo, template model.SmsTemplateInfo,
	templateParams string,
	subject, receiver, cc string) (m model.SmsMessageInfo, err error) {

	m.MessageId = message.Id
	m.Template = template.Id
	m.Params = templateParams
	m.Sender = sender.Id
	m.Subject = subject
	m.Receivers = receiver
	if len(cc) > 0 {
		m.Cc = cc
	}

	m.Status = model.SmsTemplateEnable
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()

	_, err = app.GetOrm().Context.Insert(&m)
	return
}
