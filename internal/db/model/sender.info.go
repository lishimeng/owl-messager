package model

import "github.com/lishimeng/app-starter"

// SenderInfo 发消息账号
type SenderInfo struct {
	app.Pk
	Code    string `orm:"column(code);unique"`    // 编号
	Default int    `orm:"column(default_sender)"` // 默认账号
	app.TableChangeInfo
}

const (
	DefaultSenderDisable = 0
	DefaultSenderEnable  = 1
)

const (
	SenderCategoryMail = "mail"
	SenderCategorySms  = "sms"
)
