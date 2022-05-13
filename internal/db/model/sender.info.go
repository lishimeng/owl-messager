package model

// SenderInfo 发消息账号
type SenderInfo struct {
	Pk
	Code    string `orm:"column(code);unique"` // 编号
	Default int    `orm:"column(in_default)"`  // 默认账号
	TableChangeInfo
}
