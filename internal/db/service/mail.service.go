package service

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/mailattachment"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

// CreateMailMessage 创建普通邮件
func CreateMailMessage(org int, template model.MessageTemplate, templateParams string,
	subject, receiver string, attachmentIDs []string) (m model.MessageInfo, err error) {

	var attachmentsJSON string
	if len(attachmentIDs) > 0 {
		refs, e := mailattachment.Default().Resolve(org, attachmentIDs)
		if e != nil {
			err = e
			return
		}
		attachmentsJSON = msg.MarshalMailAttachmentRefs(refs)
	}

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		m, e = repo.CreateMessage(ctx, org, subject, msg.MailMessage)
		if e != nil {
			return
		}
		_, e = repo.CreateMailMessage(ctx, m, template, templateParams, receiver, attachmentsJSON)
		if e != nil {
			return
		}
		if len(attachmentIDs) > 0 {
			e = mailattachment.Default().MarkBound(org, attachmentIDs)
		}
		return
	})
	return
}
