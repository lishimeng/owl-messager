package sender

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/internal/provider"
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
	var si model.SmsSenderInfo
	si, err = repo.GetDefaultSmsSender(0) // 默认sender
	if err != nil {
		log.Info("sms sender not exist")
		log.Info(err)
		return
	}

	tpl, err := repo.GetSmsTemplateById(mi.Template)
	if err != nil {
		log.Info("sms template not exist:%d", mi.Template)
		return
	}

	p, err := provider.GetFactory().Create(si.Vendor, si.Config)
	if err != nil {
		log.Info("create sms provider failure:%d", si.Id)
		return
	}

	var req = messager.Request{
		Template:  tpl.CloudTemplateId,
		Params:    mi.Params,
		Receivers: mi.Receivers,
		Sign:      tpl.Signature,
	}

	resp, err := p.Send(req)

	if err != nil {
		log.Info(err)
		return
	}
	log.Info(resp)
	return
}
