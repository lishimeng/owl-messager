package um

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/pkg/errors"
)

func serviceAddApns(org int, templateCode, params, subject, bundleId, receiver string) (m model.MessageInfo, err error) {
	if len(bundleId) == 0 {
		err = errors.New("bundleId required")
		return
	}

	sender, err := repo.GetDefMessageSender(org, msg.ApnsMessage, msg.Apns)
	if err != nil {
		log.Debug(errors.Wrapf(err, "apns sender not found for org:%d", org))
		return
	}

	// template 可选：若传入则校验存在且为 apns 类型
	if len(templateCode) > 0 {
		var tpl model.MessageTemplate
		tpl, err = repo.GetMessageTemplateByCode(templateCode, org)
		if err != nil || tpl.Category != msg.ApnsMessage {
			err = errors.New("template not found")
			return
		}
	}

	m, err = service.CreateApnsMessage(
		sender,
		int(model.ProductionMode),
		bundleId,
		params,
		subject,
		receiver,
	)
	return
}
