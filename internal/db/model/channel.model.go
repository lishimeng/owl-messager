package model

// TemplateChannelDetail 模板类渠道（mail / sms / im）公共字段
type TemplateChannelDetail struct {
	MessageHeader
	Template  int    `gorm:"column:template_id"`
	Params    string `gorm:"column:params"`
	Receivers string `gorm:"column:receiver"`
}

// PushChannelDetail 推送类渠道（apns 等）公共字段
type PushChannelDetail struct {
	MessageHeader
	Params    string   `gorm:"column:params"`
	Receivers string   `gorm:"column:receiver"`
	SenderId  int      `gorm:"column:sender_id"`
	BundleId  string   `gorm:"column:bundle_id"`
	ApnsMode  ApnsMode `gorm:"column:apns_mode"`
}
