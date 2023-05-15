package model

// SmsMessageInfo 短信
type SmsMessageInfo struct {
	MessageHeader
	Template  int    `orm:"column(template_id)"` // sms template
	Params    string `orm:"column(params);null"` // json params(map)
	Receivers string `orm:"column(receiver)"`    // receiver list. comma split
}
