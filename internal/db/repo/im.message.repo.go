package repo

import (
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func GetImByMessageId(msgId int) (m model.ImMessageInfo, err error) {
	err = orm().Model(&model.ImMessageInfo{}).Equal("message_id", msgId).First(&m)
	return
}

func CreateImMessage(ctx persistence.TxContext, message model.MessageInfo,
	template model.MessageTemplate,
	templateParams string,
	receiver string) (m model.ImMessageInfo, err error) {

	m.Org = message.Org
	m.MessageId = message.Id
	m.Template = template.Id
	m.Params = templateParams
	m.Receivers = receiver
	m.Status = model.MessageInit
	err = ctx.Create(&m)
	return
}
