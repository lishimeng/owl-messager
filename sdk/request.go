package sdk

// MailRequest 邮件服务请求参数
type MailRequest struct {
	Template string `json:"template"` // 模板
	//CloudTemplate bool        `json:"cloudTemplate,omitempty"` // 云端模板(不需要)
	TemplateParam interface{} `json:"params"`            // 参数
	Title         string      `json:"subject,omitempty"` // 标题
	Receiver      string      `json:"receiver"`          // 接收者，多个时用逗号分隔
}

// SmsRequest 短信服务请求参数
type SmsRequest struct {
	Template      string      `json:"template"` // 模板
	TemplateParam interface{} `json:"params"`   // 参数
	Receiver      string      `json:"receiver"` // 接收者，多个时用逗号分隔
}

// ApnsRequest Apns服务请求参数
type ApnsRequest struct {
	BundleId      string      `json:"bundleId,omitempty"` // bundle id
	Template      string      `json:"template"`           // 模板
	TemplateParam interface{} `json:"params"`             // 参数
	Title         string      `json:"subject,omitempty"`  // 标题
	Receiver      string      `json:"receiver"`           // 接收者，多个时用逗号分隔
}
