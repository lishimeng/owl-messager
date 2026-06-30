package model

type SmsMessageInfo struct {
	MessageHeader
	Template  int    `gorm:"column:template_id"`
	Params    string `gorm:"column:params"`
	Receivers string `gorm:"column:receiver"`
}
