package model

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type MessageSenderInfo struct {
	app.TenantPk
	Code     string              `orm:"column(code);unique"`      // 编号
	Category msg.MessageCategory `orm:"column(message_category)"` // 消息类型
	Provider msg.MessageProvider `orm:"column(message_provider)"` // 消息平台
	Default  int                 `orm:"column(default_sender)"`   // 默认账号 unique:org+vendor
	Config   msg.SenderConfig    `orm:"column(config)"`           // json 配置(map: key-value)
	SenderAppInfo
	app.TableChangeInfo
}

type SenderAppInfo struct { // wechat mini/hms/gms/apns等服务,标记应用程序通道
	AppIdentify string `orm:"column(app_identify);null"`
}

const (
	DefaultSenderDisable = 0
	DefaultSenderEnable  = 1
)
const (
	SenderDisable = 0
	SenderEnable  = 1
)
