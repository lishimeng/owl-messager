package um

import (
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
)

func serviceAddSms(org int, smsTemplate, tplParams, receiver string) (m model.MessageInfo, err error) {

	tpl, err := repo.GetSmsTemplateByCode(smsTemplate)
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
