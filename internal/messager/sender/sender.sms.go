package sender

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/internal/provider"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type Sms interface {
	Send(model.SmsMessageInfo) (err error)
}

type smsSender struct {
	ctx       context.Context
	maxWorker int
}

func NewSmsSender(ctx context.Context) (m Sms, err error) {
	m = &smsSender{
		ctx:       ctx,
		maxWorker: 1,
	}
	return
}

func (m *smsSender) Send(mi model.SmsMessageInfo) (err error) {
	// sender info
	log.Info("send sms:%d", mi.Id)
	// 获取模板
	tpl, err := repo.GetMessageTemplateById(mi.Template)
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

	params, err := msg.HandleMessageParams(mi.Params, tpl.Params)
	if err != nil {
		return
	}

	p, err := provider.GetFactory().Create(si.Provider, string(si.Config))
	if err != nil {
		log.Info("create sms provider failure:%d", si.Id)
		return
	}

	var req = messager.Request{
		Template:  tpl,
		Params:    params,
		Receivers: mi.Receivers,
	}

	resp, err := p.Send(req)

	if err != nil {
		log.Info(err)
		return
	}
	log.Info(resp)
	return
}
