package repo

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl/internal/db/model"
)

func GetSmsByMessageId(msgId int) (m model.SmsMessageInfo, err error) {
	err = app.GetOrm().Context.
		QueryTable(new(model.SmsMessageInfo)).
		Filter("MessageId", msgId).
		One(&m)
	return
}

func CreateSmsMessage(ctx persistence.TxContext, message model.MessageInfo, sender *model.SmsSenderInfo, template model.SmsTemplateInfo,
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

	_, err = ctx.Context.Insert(&m)
	return
}
