package model

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

// MessageTemplate 消息模板
type MessageTemplate struct {
	app.TenantPk
	Code          string              `orm:"column(code);unique"`      // owl中的唯一编码
	Name          string              `orm:"column(name)"`             // 模板名字
	Category      msg.MessageCategory `orm:"column(message_category)"` // 类型sms/mail/apns
	Body          string              `orm:"column(body);null"`        // 发送的内容主体，可空
	Params        string              `orm:"column(params);null"`      // 参数列表映射: "对外参数":["对内参数列表"], "a":["m", "n"]
	Provider      msg.MessageProvider `orm:"column(message_provider)"`
	CloudTemplate string              `orm:"column(cloud_template);null"` // vendor为cloud时, 存储平台中的模板ID
	Description   string              `orm:"column(description);null"`
	app.TableChangeInfo
}

const (
	TemplateEnable  = 1 // enable
	TemplateDisable = 0 // disable
)
