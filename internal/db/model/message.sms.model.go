package model

type SmsMessageInfo struct {
	TemplateChannelDetail
}

func (SmsMessageInfo) TableName() string {
	return "sms_message_info"
}
