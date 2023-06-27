package sdk

import (
	"bytes"
	"encoding/json"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/messager/msg"
	"io"
	"net/http"
	"time"
)

const ApiPath = "/v2/messages/"

// MessageClient 消息服务
type MessageClient struct {
	Host     string // 消息服务主机地址. 如, "http://127.0.0.1/api"
	category string // 消息类型. "mail","sms","apns"
}

// NewClient 初始化邮件服务信息。host: 邮件服务主机地址
func NewClient(host string) (m *MessageClient) {
	return &MessageClient{
		Host: host,
	}
}

func (m *MessageClient) setCategory(category string) {
	m.category = category
}

func (m *MessageClient) getURL() string {
	return m.Host + ApiPath + "/" + m.category
}

func (m *MessageClient) SendMail(request MailRequest) (response Response, err error) {
	log.Debug("sendMail to: %s", request.Receiver)
	m.setCategory(msg.EmailCategory)
	response, err = m.send(request)
	if err != nil {
		log.Debug(err)
		return
	}
	return
}
func (m *MessageClient) SendSms(request SmsRequest) (response Response, err error) {
	log.Debug("sendMail to: %s", request.Receiver)
	m.setCategory(msg.SmsCategory)
	response, err = m.send(request)
	if err != nil {
		log.Debug(err)
		return
	}
	return
}
func (m *MessageClient) SendApns(request ApnsRequest) (response Response, err error) {
	log.Debug("sendMail to: %s", request.Receiver)
	m.setCategory(msg.ApnsCategory)
	response, err = m.send(request)
	if err != nil {
		log.Debug(err)
		return
	}
	return
}

func (m *MessageClient) send(request interface{}) (response Response, err error) {
	url := m.getURL()
	log.Debug("sendMail url: %s", url)
	client := &http.Client{Timeout: 8 * time.Second}
	jsonStr, _ := json.Marshal(request)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Debug("client Post err")
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	result, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(result, &response)
	if err != nil {
		log.Debug("response Unmarshal err, %+v", response)
		return
	}
	log.Debug("sendMail response: %v", response)
	return
}
