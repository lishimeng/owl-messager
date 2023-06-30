package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/messager/msg"
	"github.com/pkg/errors"
	"net/url"
)

const ApiSendMessage = "/v2/messages/"

const ApiCredential = "/v2/open/oauth2/token"

const (
	CodeNotAllow int = 401
	CodeSuccess  int = 200
)

// messageClient 消息服务
type messageClient struct {
	host        string // 消息服务主机地址. 如, "http://127.0.0.1/api"
	debugEnable bool
	credential  string
	appId       string
	secret      string
}

type Option func(*messageClient)

func WithHost(host string) Option {
	return func(client *messageClient) {
		client.host = host
	}
}

func WithAuth(appKey, secret string) Option {
	return func(client *messageClient) {
		client.appId = appKey
		client.secret = secret
	}
}

func WithDebug(enable bool) Option {
	return func(client *messageClient) {
		client.debugEnable = enable
	}
}

// New 初始化
func New(options ...Option) (m Client) {
	c := &messageClient{}
	for _, opt := range options {
		opt(c)
	}
	m = c
	return
}

func (m *messageClient) SendMail(request MailRequest) (response Response, err error) {
	if m.debugEnable {
		log.Debug("sendMail to: %s", request.Receiver)
	}
	response, err = m.send(msg.EmailCategory, request)
	if err != nil {
		log.Debug(err)
		return
	}
	return
}
func (m *messageClient) SendSms(request SmsRequest) (response Response, err error) {
	if m.debugEnable {
		log.Debug("sendMail to: %s", request.Receiver)
	}
	response, err = m.send(msg.SmsCategory, request)
	if err != nil {
		log.Debug(err)
		return
	}
	return
}
func (m *messageClient) SendApns(request ApnsRequest) (response Response, err error) {
	if m.debugEnable {
		log.Debug("sendMail to: %s", request.Receiver)
	}
	response, err = m.send(msg.ApnsCategory, request)
	if err != nil {
		log.Debug(err)
		return
	}
	return
}

func (m *messageClient) refreshCredential() (err error) {
	host, err := url.JoinPath(m.host, ApiCredential)
	if err != nil {
		return
	}
	response, err := getCredential(host, m.appId, m.secret)
	if err != nil {
		return
	}
	m.credential = response.Token
	return
}

func (m *messageClient) send(category string, request interface{}) (response Response, err error) {
	jsonStr, _ := json.Marshal(request)
	u, err := url.JoinPath(m.host, ApiSendMessage, category)
	if err != nil {
		log.Debug(errors.Wrap(err, "get url fail"))
		return
	}
	if m.debugEnable {
		log.Debug("sendMail url: %s", u)
	}
	code, response, err := _send(m.credential, u, jsonStr)
	if err != nil {
		log.Debug(errors.Wrap(err, "send fail"))
		return
	}
	// http 无异常, 检查response code, 如果CodeNotAllow说明token不正常
	if code == CodeNotAllow {
		log.Debug("credential expired, refresh credential")
		err = m.refreshCredential()
		if err != nil {
			log.Debug(errors.Wrap(err, "can't refresh credential"))
			return
		}
		// 如果获取了新token,重新发一遍
		log.Debug("re_send the message")
		code, response, err = _send(m.credential, u, jsonStr)
	}
	if err != nil {
		log.Debug(errors.Wrap(err, "send fail"))
		return
	}
	if code == CodeNotAllow {
		// 如果还是 CodeNotAllow, 说明token系统出问题了
		err = errors.New(fmt.Sprintf("%d", CodeNotAllow))
	}
	if m.debugEnable {
		log.Debug("sendMail response: %v", response)
	}
	return
}
