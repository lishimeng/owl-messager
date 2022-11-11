package model

// MailSenderInfo 邮件发送账号
type MailSenderInfo struct {
	SenderInfo
	Vendor MailVendor
	Host     string `orm:"column(host)"`
	Port     int `orm:"column(port)"`
	Email     string `orm:"column(email)"`
	Password     string `orm:"column(password)"`
	Config string `orm:"column(config)"` // json 配置(map: key-value)
}

type MailVendor string

const (
	MailVendorSmtp      MailVendor = "smtp"
	MailVendorMicrosoft MailVendor = "microsoft"
)

// MailVendors vendor:enable?1:0
var MailVendors map[MailVendor]byte

func init() {
	MailVendors = make(map[MailVendor]byte)
	MailVendors[MailVendorSmtp] = 0x01
	MailVendors[MailVendorMicrosoft] = 0x01
}

// Support 数据库记录的vendor是否被支持
func (s MailSenderInfo) Support() bool {
	val, ok := MailVendors[s.Vendor]
	return ok && (val > 0)
}

type SmtpConfig struct {
	Host        string `json:"host,omitempty"`
	Port        int    `json:"port,omitempty"`
	SenderEmail string `json:"senderEmail,omitempty"` // 发件邮箱
	SenderAlias string `json:"senderAlias,omitempty"` // 发件人名字
	AuthUser    string `json:"authUser,omitempty"`    // 发件账号(有些邮件服务器为邮箱地址)
	AuthPass    string `json:"authPass,omitempty"`    // 密码
}

type GraphConfig struct {
	ClientId       string `json:"clientId,omitempty"`
	Tenant         string `json:"tenant,omitempty"`
	Scope          string `json:"scope,omitempty"`
	Sender         string `json:"sender,omitempty"`
	Certificate    string `json:"certificate,omitempty"`
	CertificateKey string `json:"certificateKey,omitempty"`
}
