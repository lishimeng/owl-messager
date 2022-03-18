package model

type SmsTemplateInfo struct {
	Pk
	Code             string `orm:"column(code);unique"` // 发送平台的编号
	Name             string `orm:"column(name)"`
	Body             string `orm:"column(body)"`               // 发送的内容主体，可空
	SenderTemplateId string `orm:"column(sender_template_id)"` // 发送平台的模板ID
	Signature        string `orm:"column(signature)"`          // 在发送平台上预留的签名，根据不同平台保存签名文本或签名ID
	Description      string `orm:"column(description);null"`
	// TODO 参数列表，用来预判断参数是否符合条件
	TableChangeInfo
}

const (
	SmsTemplateEnable  = 1 // enable
	SmsTemplateDisable = 0 // disable
)
