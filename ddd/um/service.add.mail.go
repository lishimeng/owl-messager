package um

import (
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
)

func serviceAddMail(templateCode string, cloudTemplate bool, cloudTemplateCode string, params, subject, receiver string) (m model.MessageInfo, err error) {
	var tpl model.MailTemplateInfo
	if !cloudTemplate { // 本地模板, 获取模板信息
		tpl, err = repo.GetMailTemplateByCode(templateCode)
		if err != nil {
			return
		}
		m, err = service.CreateMailMessage(
			nil,
			tpl,
			params,
			subject, receiver, "")
	} else {
		m, err = service.CreateCloudMailMessage(
			nil,
			cloudTemplateCode,
			params,
			subject, receiver, "")
	}

	return
}
