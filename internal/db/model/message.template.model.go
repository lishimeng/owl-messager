package model

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

// MessageTemplate 消息模板
type MessageTemplate struct {
	TenantScope   `gorm:"uniqueIndex:idx_tpl_tenant_code"`
	Code          string              `gorm:"column:code;uniqueIndex:idx_tpl_tenant_code"`
	Name          string              `gorm:"column:name"`
	Category      msg.MessageCategory `gorm:"column:category;type:varchar(16)"`
	Body          string              `gorm:"column:body"`
	Params        string              `gorm:"column:params"`
	Provider      msg.MessageProvider `gorm:"column:message_provider"`
	CloudTemplate string              `gorm:"column:cloud_template"`
	Description   string              `gorm:"column:description"`
	app.TableChangeInfo
}

func (MessageTemplate) TableName() string {
	return "message_template"
}
