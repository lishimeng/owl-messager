package sender

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/internal/provider"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"strings"
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
	// sender info
	log.Info("send mail:%d", p.Id)

	// 获取模板
	tpl, err := repo.GetMessageTemplateById(p.Template)
	if err != nil {
		log.Info("tpl not exist")
		return
	}

	// 获取sender
	si, err := repo.GetDefMessageSender(tpl.Org, tpl.Category, tpl.Provider) // 使用默认sender

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

	receivers := strings.Split(p.Receivers, ",")

	req := messager.MailRequest{
		Subject:   p.Subject,
		Receivers: receivers,
		Template:  tpl,
		Params:    params,
	}

	err = s.Send(req)
	return
}
