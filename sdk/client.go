package sdk

type Client interface {
	SendMail(req MailRequest) (resp Response, err error)
	SendSms(req SmsRequest) (resp Response, err error)
	SendApns(req ApnsRequest) (resp Response, err error)
}

type Option func(*messageClient)

var (
	debugEnable = false
)

// Debug 输出sdk中的log
func Debug(enable bool) {
	debugEnable = enable
}

func WithHost(host string) Option {
	return func(client *messageClient) {
		client.host = host
	}
}

// WithAuth Credentials配置
func WithAuth(appKey, secret string) Option {
	return func(client *messageClient) {
		client.appId = appKey
		client.secret = secret
	}
}

// New 初始化.
//
// 内部存储了Credentials,应该确保复用,而不是每次新建
//
// Client 内置了刷新credentials功能,不需要考虑credentials的获取问题.
func New(options ...Option) (m Client) {
	c := &messageClient{}
	for _, opt := range options {
		opt(c)
	}
	m = c
	return
}
