package repo

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"time"
)

func GetMailByMessageId(msgId int) (m model.MailMessageInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MailMessageInfo)).Filter("MessageId", msgId).One(&m)
	return
}

func CreateMailMessage(ctx persistence.TxContext, message model.MessageInfo,
	template model.MailTemplateInfo,
	templateParams string,
	subject, receiver string) (m model.MailMessageInfo, err error) {

	m.MessageId = message.Id
	m.Template = template.Id
	m.Params = templateParams

	m.Subject = subject
	m.Receivers = receiver

	m.Status = model.MailTemplateEnable
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()

	_, err = ctx.Context.Insert(&m)
	return
}

func CreateCloudMailMessage(ctx persistence.TxContext, message model.MessageInfo,
	cloudTemplateId string,
	templateParams string,
	subject, receiver string) (m model.MailMessageInfo, err error) {

	m.MessageId = message.Id
	m.CloudTemplate = 1
	m.CloudTemplateId = cloudTemplateId
	m.Params = templateParams

	m.Subject = subject
	m.Receivers = receiver

	m.Status = model.MailTemplateEnable
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()

	_, err = ctx.Context.Insert(&m)
	return
}
