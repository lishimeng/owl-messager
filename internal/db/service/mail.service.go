package service

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

// CreateMailMessage 创建普通邮件
func CreateMailMessage(org int, template model.MessageTemplate, templateParams string,
	subject, receiver, cc string) (m model.MessageInfo, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		// create message
		m, e = repo.CreateMessage(ctx, org, subject, msg.MailMessage)
		if e != nil {
			return
		}
		// create mail
		_, _ = repo.CreateMailMessage(ctx, m, template, templateParams, subject, receiver)
		return
	})
	return
}

//func UpdateMailTemplate(id, status int, body, description string) (m model.MailTemplateInfo, err error) {
//	//m, err = repo.GetMailTemplateById(id)
//	//if err != nil {
//	//	return
//	//}
//	//err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
//	//	var cols []string
//	//	if status > repo.ConditionIgnore {
//	//		m.Status = status
//	//		cols = append(cols, "Status")
//	//	}
//	//	if len(body) > 0 {
//	//		m.Body = body
//	//		cols = append(cols, "Body")
//	//	}
//	//	if len(description) > 0 {
//	//		m.Description = description
//	//		cols = append(cols, "Description")
//	//	}
//	//
//	//	m, e = repo.UpdateMailTemplate(ctx, m, cols...)
//	//
//	//	return
//	//})
//	return TODO
//}
//
//// SetDefaultMailSender 设置默认发送账号
//func SetDefaultMailSender(id int, org int) (err error) {
//
//	senders, err := repo.GetMailSenders(org)
//	if err != nil {
//		return
//	}
//	for _, s := range senders {
//		if s.Id == id {
//			s.Default = model.DefaultSenderEnable
//		} else {
//			s.Default = model.DefaultSenderDisable
//		}
//	}
//	return
//}
