package um

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
	"github.com/pkg/errors"
)

func serviceAddMail(org int, templateCode string, params, subject, receiver string) (m model.MessageInfo, err error) {
	var tpl model.MailTemplateInfo
	tpl, err = repo.GetMailTemplateByCode(templateCode) // TODO org
	if err != nil {
		log.Debug(errors.Wrapf(err, "template not found:%s", templateCode))
		return
	}
	cloudTemplate := tpl.Cloud == model.MailCloudTemplate
	if cloudTemplate { // 本地模板, 获取模板信息
		cloudTemplateCode := tpl.CloudId
		m, err = service.CreateCloudMailMessage(
			org,
			cloudTemplateCode,
			params,
			subject, receiver, "")

	} else {
		m, err = service.CreateMailMessage(
			org,
			tpl,
			params,
			subject, receiver, "")
	}

	return
}
