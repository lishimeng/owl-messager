package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
)

func GetSmsByMessageId(msgId int) (m model.SmsMessageInfo, err error) {
	// TODO
	return
}

func CreateSmsMessage(message model.MessageInfo, sender *model.SmsSenderInfo, template model.SmsTemplateInfo,
	templateParams string, receiver string) (m model.SmsMessageInfo, err error) {

	m.MessageId = message.Id
	m.Template = template.Id
	m.Params = templateParams
	m.Signature = template.Signature
	m.Receivers = receiver
	if sender != nil {
		m.Sender = sender.Id
	}

	m.Status = model.SmsTemplateEnable

	_, err = app.GetOrm().Context.Insert(&m)
	return
}
