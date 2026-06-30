package model

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type MessageSenderInfo struct {
	app.Pk
	Org      int                 `gorm:"column:org;uniqueIndex:idx_sender_org_code"`
	Code     string              `gorm:"column:code;uniqueIndex:idx_sender_org_code"`
	Category msg.MessageCategory `gorm:"column:message_category"`
	Provider msg.MessageProvider `gorm:"column:message_provider"`
	Default  int                 `gorm:"column:default_sender"`
	Config   msg.SenderConfig    `gorm:"column:config"`
	SenderAppInfo
	app.TableChangeInfo
}

type SenderAppInfo struct {
	AppIdentify string `gorm:"column:app_identify"`
}
