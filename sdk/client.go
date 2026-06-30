package sdk

import (
	"errors"
	"github.com/lishimeng/owl-messager/utils"
)

type Client interface {
	SendMail(req MailRequest) (resp Response, err error)
	SendSms(req SmsRequest) (resp Response, err error)
	SendApns(req ApnsRequest) (resp Response, err error)
	SendIm(req ImRequest) (resp Response, err error)
	Templates(req TemplateRequest) (resp TemplateResponse, err error)
}

type Option func(*messageClient)

var (
	InnerServerErr = errors.New("500")
	NotFoundErr    = errors.New("404")
)

var (
	debugEnable = false
)

// Debug 输出sdk中的log
func Debug(enable bool) {
	debugEnable = enable
	utils.DebugEnable = enable
}

func WithHost(host string) Option {
	return func(client *messageClient) {
		client.host = host
	}
}

// WithAuth sets appId and secret for HTTP Basic authentication.
func WithAuth(appKey, secret string) Option {
	return func(client *messageClient) {
		client.appId = appKey
		client.secret = secret
	}
}

// New creates a reusable message client.
//
// Pass WithAuth(appId, secret) so each request uses Authorization: Basic.
func New(options ...Option) (m Client) {
	c := &messageClient{}
	for _, opt := range options {
		opt(c)
	}
	m = c
	return
}
