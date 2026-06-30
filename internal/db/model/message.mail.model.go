package model

type MailMessageInfo struct {
	MessageHeader
	Template  int    `gorm:"column:template_id"`
	Params    string `gorm:"column:params"`
	Receivers string `gorm:"column:receiver"`
	Subject   string `gorm:"column:subject"`
}
