package sdk

import (
	"bytes"
	"encoding/json"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/messager/msg"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

const ApiPath = "/v2/messages/"

// MessageClient 消息服务
type MessageClient struct {
	host        string // 消息服务主机地址. 如, "http://127.0.0.1/api"
	debugEnable bool
}

type Option func(*MessageClient)

func WithHost(host string) Option {
	return func(client *MessageClient) {
		client.host = host
	}
}

func WithDebug(enable bool) Option {
	return func(client *MessageClient) {
		client.debugEnable = enable
	}
}

// NewClient 初始化邮件服务信息。host: 邮件服务主机地址
func NewClient(options ...Option) (m *MessageClient) {
	m = &MessageClient{}
	for _, opt := range options {
		opt(m)
	}
	return
}

func (m *MessageClient) getURL(category string) (address string, err error) {
	address, err = url.JoinPath(m.host, ApiPath, category)
	return
}

func (m *MessageClient) SendMail(request MailRequest) (response Response, err error) {
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
func (m *MessageClient) SendSms(request SmsRequest) (response Response, err error) {
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
func (m *MessageClient) SendApns(request ApnsRequest) (response Response, err error) {
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

func (m *MessageClient) send(category string, request interface{}) (response Response, err error) {
	u, err := m.getURL(category)
	if err != nil {
		log.Debug(errors.Wrap(err, "get url fail"))
		return
	}
	if m.debugEnable {
		log.Debug("sendMail url: %s", u)
	}
	client := &http.Client{Timeout: 8 * time.Second}
	jsonStr, _ := json.Marshal(request)
	resp, err := client.Post(u, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Debug(errors.Wrap(err, "client Post err"))
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	result, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(result, &response)
	if err != nil {
		log.Debug(errors.Wrap(err, "response json unmarshal err"))
		return
	}
	if m.debugEnable {
		log.Debug("sendMail response: %v", response)
	}
	return
}
