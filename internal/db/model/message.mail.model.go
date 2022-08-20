package model

// MailMessageInfo 邮件
type MailMessageInfo struct {
	MessageHeader
	Template  int    `orm:"column(template_id)"` // 模板
	Params    string `orm:"column(params);null"` // 参数
	Receivers string `orm:"column(receiver)"`    // 收件人，多个情况下用逗号分隔
	Subject   string `orm:"column(subject)"`     // 标题
}
