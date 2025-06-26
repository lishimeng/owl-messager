package um

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/pkg/errors"
)

func serviceAddIm(org int, templateCode string, params, receiver string) (m model.MessageInfo, err error) {
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
	if tpl.Category != msg.ImMessage {
		log.Debug("tpl category: %s, expect: %s", tpl.Category, msg.ImMessage)
		err = errors.New("template not found")
		return
	}

	m, err = service.CreateImMessage(
		org,
		tpl,
		params,
		receiver)

	return
}
