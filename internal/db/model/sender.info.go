package model

// SenderInfo 发消息账号
type SenderInfo struct {
	Pk
	Code    string `orm:"column(code);unique"`    // 编号
	Default int    `orm:"column(default_sender)"` // 默认账号
	TableChangeInfo
}

const (
	DefaultSenderDisable = 0
	DefaultSenderEnable  = 1
)

const (
	SenderCategoryMail = "mail"
	SenderCategorySms  = "sms"
)
