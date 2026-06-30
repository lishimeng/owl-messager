package model

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

// MessageTemplate 消息模板
type MessageTemplate struct {
	app.Pk
	Org           int                 `gorm:"column:org;uniqueIndex:idx_tpl_org_code"`
	Code          string              `gorm:"column:code;uniqueIndex:idx_tpl_org_code"`
	Name          string              `gorm:"column:name"`
	Category      msg.MessageCategory `gorm:"column:message_category"`
	Body          string              `gorm:"column:body"`
	Params        string              `gorm:"column:params"`
	Provider      msg.MessageProvider `gorm:"column:message_provider"`
	CloudTemplate string              `gorm:"column:cloud_template"`
	Description   string              `gorm:"column:description"`
	app.TableChangeInfo
}
