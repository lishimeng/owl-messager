package model

// 邮件
type MailMessageInfo struct {
	MessageHeader

	Template int `orm:"column(template_id)"` // mail template

	Params string `orm:"column(params)"` // json params

	Sender int `orm:"column(sender_id)"` // sender's Id

	Receivers string `orm:"column(receiver)"` // receiver list. comma split

	Cc string `orm:"column(cc)"` // CC

	// 主题
	Subject string `orm:"column(subject)"`
}
