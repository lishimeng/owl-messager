package model

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type MessageSenderInfo struct {
	TenantScope `gorm:"uniqueIndex:idx_sender_tenant_code"`
	Code        string              `gorm:"column:code;uniqueIndex:idx_sender_tenant_code"`
	Category msg.MessageCategory `gorm:"column:category;type:varchar(16)"`
	Provider msg.MessageProvider `gorm:"column:message_provider"`
	Default  int                 `gorm:"column:default_sender"`
	Config   msg.SenderConfig    `gorm:"column:config"`
	SenderAppInfo
	app.TableChangeInfo
}

func (MessageSenderInfo) TableName() string {
	return "message_sender_info"
}

type SenderAppInfo struct {
	AppIdentify string `gorm:"column:app_identify"`
}
