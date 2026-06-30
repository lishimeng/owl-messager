package repo

import (
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func GetMailByMessageId(msgId int) (m model.MailMessageInfo, err error) {
	err = orm().Model(&model.MailMessageInfo{}).Equal("message_id", msgId).First(&m)
	return
}

func CreateMailMessage(ctx persistence.TxContext, message model.MessageInfo,
	template model.MessageTemplate,
	templateParams string,
	receiver, attachmentsJSON string) (m model.MailMessageInfo, err error) {

	initTemplateChannel(&m.TemplateChannelDetail, message, template, templateParams, receiver)
	m.Attachments = attachmentsJSON
	err = ctx.Create(&m)
	return
}
