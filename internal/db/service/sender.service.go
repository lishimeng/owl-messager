package service

import (
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

// GetDefaultSmsSender 获取Sms默认发送账号
func GetDefaultSmsSender(org int) (sender model.SmsSenderInfo, err error) {

	sender, err = repo.GetDefaultSmsSender(org)
	if err != nil {
		return
	}
	err = sender.Config.Decode()
	return
}

// GetDefaultEmailSender 获取Email默认发送账号
func GetDefaultEmailSender(org int) (sender model.MailSenderInfo, err error) {

	sender, err = repo.GetDefaultMailSender(org)
	if err != nil {
		return
	}
	err = sender.Config.Decode()
	return
}
