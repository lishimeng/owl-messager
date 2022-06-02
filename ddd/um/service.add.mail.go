package um

import (
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/db/service"
)

func serviceAddMail(templateCode, params, subject, receiver string) (m model.MessageInfo, err error) {
	tpl, err := repo.GetMailTemplateByCode(templateCode)
	if err != nil {
		return
	}
	m, err = service.CreateMailMessage(
		nil,
		tpl,
		params,
		subject, receiver, "")
	return
}
