package sender

import (
	"context"

	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/mailattachment"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/internal/provider"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type Mail interface {
	Send(model.MailMessageInfo) (err error)
}

type mailSender struct {
	ctx       context.Context
	maxWorker int
}

func NewMailSender(ctx context.Context) (m Mail, err error) {
	m = &mailSender{
		ctx:       ctx,
		maxWorker: 1,
	}
	return
}

func (m *mailSender) Send(p model.MailMessageInfo) (err error) {
	log.Info("send mail:%d", p.Id)

	tpl, err := repo.GetMessageTemplateById(p.Template)
	if err != nil {
		log.Info("tpl not exist")
		return
	}

	si, err := repo.GetDefMessageSender(tpl.Org, tpl.Category, tpl.Provider)
	if err != nil {
		log.Info("mail sender not exist")
		return
	}

	params, err := msg.HandleMessageParams(p.Params, tpl.Params)
	if err != nil {
		return
	}

	s, err := provider.DefaultMailFactory.Create(si.Provider, string(si.Config))
	if err != nil {
		log.Info("create mail sender failure:%d", si.Id)
		return
	}

	msgInfo, err := repo.GetMessageById(p.MessageId)
	if err != nil {
		log.Info("message not exist:%d", p.MessageId)
		return
	}

	attachments, err := mailattachment.Default().LoadForSend(p.Org, p.Attachments)
	if err != nil {
		log.Info("load attachments failed: %v", err)
		return
	}

	receivers := msg.SplitReceivers(p.Receivers)

	req := messager.MailRequest{
		Subject:     msgInfo.Subject,
		Receivers:   receivers,
		Template:    tpl,
		Params:      params,
		Attachments: attachments,
	}

	err = s.Send(req)
	if err != nil {
		return
	}
	mailattachment.Default().ScheduleRemove(p.Org, p.Attachments)
	return
}
