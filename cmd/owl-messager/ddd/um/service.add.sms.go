package um

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/pkg/errors"
)

func serviceAddSms(org int, templateCode, tplParams, receiver string) (m model.MessageInfo, err error) {
	var tpl model.MessageTemplate
	tpl, err = repo.GetMessageTemplateByCode(templateCode) // TODO org
	if err != nil {
		log.Debug(errors.Wrapf(err, "template not found:%s", templateCode))
		return
	}
	if tpl.Org != org {
		err = errors.New("template not found")
		return
	}
	if tpl.Category != msg.SmsMessage {
		log.Debug("tpl category: %s, expect: %s", tpl.Category, msg.SmsMessage)
		err = errors.New("template not found")
		return
	}

	if err != nil {
		return
	}
	m, err = service.CreateSmsMessage(
		org,
		tpl,
		tplParams,
		receiver,
	)
	return
}
