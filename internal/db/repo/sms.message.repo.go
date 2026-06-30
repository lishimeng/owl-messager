package repo

import (
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func GetSmsByMessageId(msgId int) (m model.SmsMessageInfo, err error) {
	err = orm().Model(&model.SmsMessageInfo{}).Equal("message_id", msgId).First(&m)
	return
}

func CreateSmsMessage(ctx persistence.TxContext, message model.MessageInfo,
	template model.MessageTemplate,
	templateParams string,
	receiver string) (m model.SmsMessageInfo, err error) {

	m.Org = message.Org
	m.MessageId = message.Id
	m.Template = template.Id
	m.Params = templateParams
	m.Receivers = receiver
	m.Status = model.MessageInit
	err = ctx.Create(&m)
	return
}
