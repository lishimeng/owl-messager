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

	err = createTemplateChannelMessage(ctx, &m, &m.TemplateChannelDetail, message, template, templateParams, receiver)
	return
}
