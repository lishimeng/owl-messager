package service

import "github.com/lishimeng/owl/internal/db/model"

func CreateApnsMessage(sender model.SmsSenderInfo, template model.SmsTemplateInfo, templateParams string,
	receiver string) (m model.MessageInfo, err error) {

	// TODO add apns

	// TODO add Message
	return
}
