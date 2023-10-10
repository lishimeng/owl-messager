package msg

type MessageTemplate struct {
	Code          string          // owl中的唯一编码
	Name          string          // 模板名字
	Category      MessageCategory // 类型sms/mail/apns
	Body          string          // 发送的内容主体，可空
	Params        string          // 参数列表映射: "对外参数":["对内参数列表"], "a":["m", "n"]
	Provider      MessageProvider
	CloudTemplate string // vendor为cloud时, 存储平台中的模板ID
	Description   string
}
