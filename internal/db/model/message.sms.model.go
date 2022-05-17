package model

// SmsMessageInfo 短信
type SmsMessageInfo struct {
	MessageHeader

	Template int `orm:"column(template_id)"` // sms template

	Params string `orm:"column(params);null"` // json params(map)

	Sender int `orm:"column(sender_id);null"` // 指定发送者

	Receivers string `orm:"column(receiver)"` // receiver list. comma split

	Signature string `orm:"column(signature);null"` // signature

}
