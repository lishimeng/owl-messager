package repo

import (
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"time"
)

func GetMailByMessageId(msgId int) (m model.MailMessageInfo, err error) {
	err = orm().Model(&model.MailMessageInfo{}).Equal("message_id", msgId).First(&m)
	return
}

func CreateMailMessage(ctx persistence.TxContext, message model.MessageInfo,
	template model.MessageTemplate,
	templateParams string,
	subject, receiver string) (m model.MailMessageInfo, err error) {

	m.Org = message.Org
	m.MessageId = message.Id
	m.Template = template.Id
	m.Params = templateParams
	m.Subject = subject
	m.Receivers = receiver
	m.Status = model.MessageInit
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	err = ctx.Create(&m)
	return
}
