package msg

type MessageProvider string

const (
	Smtp      MessageProvider = "smtp"
	Microsoft MessageProvider = "microsoft"
	Baidu     MessageProvider = "baidu_yun"
	Ali       MessageProvider = "ali_yun"
	Huawei    MessageProvider = "huawei_yun"
	QiNiu     MessageProvider = "qi_niu_yun"
	Tencent   MessageProvider = "tencent_yun"
	UpYun     MessageProvider = "up_yun"
	Apns      MessageProvider = "apns"
)

func (mp MessageProvider) String() string {
	return string(mp)
}

type AliSmsConfig struct {
	AppKey    string `json:"appKey,omitempty"`
	AppSecret string `json:"appSecret,omitempty"`
	Region    string `json:"region,omitempty"`
	SignName  string `json:"signName,omitempty"`
}

type TencentSmsConfig struct {
	AppId    string `json:"appId,omitempty"`
	AppKey   string `json:"appKey,omitempty"`
	SmsAppId string `json:"smsAppId,omitempty"`
	Region   string `json:"region,omitempty"`
	SignName string `json:"signName,omitempty"`
}
type HuaweiSmsConfig struct {
	Host     string `json:"host,omitempty"`
	AppId    string `json:"appId,omitempty"`
	AppKey   string `json:"appKey,omitempty"`
	Sender   string `json:"sender,omitempty"`
	SignName string `json:"signName,omitempty"`
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

type TencentConfig struct {
	AppId  string `json:"appId,omitempty"`
	Secret string `json:"secret,omitempty"`
	Region string `json:"region,omitempty"`
	Sender string `json:"sender,omitempty"`
}
