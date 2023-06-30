package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd/open"
	"github.com/lishimeng/owl-messager/internal/messager/msg"
	"github.com/pkg/errors"
	"net/url"
)

const ApiSendMessage = "/messages/"

const ApiCredential = "/open/oauth2/token"

const (
	CodeNotAllow int = 401
	CodeNotFound int = 404
	CodeSuccess  int = 200
)

// messageClient 消息服务
type messageClient struct {
	host       string // 消息服务主机地址. 如, "http://127.0.0.1/api"
	credential string
	appId      string
	secret     string
}

func (m *messageClient) SendMail(request MailRequest) (response Response, err error) {
	if debugEnable {
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
	if debugEnable {
		log.Debug("sendSms to: %s", request.Receiver)
	}
	response, err = m.send(msg.SmsCategory, request)
	if err != nil {
		log.Debug(err)
		return
	}
	return
}
func (m *messageClient) SendApns(request ApnsRequest) (response Response, err error) {
	if debugEnable {
		log.Debug("sendApns to: %s", request.Receiver)
	}
	response, err = m.send(msg.ApnsCategory, request)
	if err != nil {
		log.Debug(err)
		return
	}
	return
}

func (m *messageClient) refreshCredential() (response open.CredentialResp, err error) {
	host, err := url.JoinPath(m.host, ApiCredential)
	if err != nil {
		return
	}
	if debugEnable {
		log.Debug("credential url:%s", host)
	}
	response, err = getCredential(host, m.appId, m.secret)
	if err != nil {
		return
	}
	m.credential = response.Token
	return
}

func (m *messageClient) send(category string, request any) (response Response, err error) {
	jsonStr, _ := json.Marshal(request)
	u, err := url.JoinPath(m.host, ApiSendMessage, category)
	if err != nil {
		err = errors.Wrap(err, "get url fail")
		if debugEnable {
			log.Debug(err)
		}
		return
	}
	if debugEnable {
		log.Debug("send message url: %s", u)
	}
	code, response, err := _send(m.credential, u, jsonStr)
	if err != nil {
		err = errors.Wrap(err, "send fail")
		if debugEnable {
			log.Debug(err)
		}
		return
	}
	// http 无异常, 检查response code, 如果CodeNotAllow说明token不正常
	switch code {
	case CodeNotAllow:
		code, response, err = m.resend(u, jsonStr)
	case CodeSuccess:
		if debugEnable {
			log.Debug("credential valid, send success")
		}
	default: // not found, redirect ...
		err = errors.New(fmt.Sprintf("%d", code))
	}

	if err != nil {
		err = errors.Wrap(err, "send fail")
		if debugEnable {
			log.Debug(err)
		}
		return
	}
	if code == CodeNotAllow {
		// 如果还是 CodeNotAllow, 说明token系统出问题了
		err = errors.New(fmt.Sprintf("%d", CodeNotAllow))
	}
	return
}

func (m *messageClient) resend(u string, data []byte) (code int, response Response, err error) {
	if debugEnable {
		log.Debug("credential expired, refresh credential")
	}
	credentialResp, err := m.refreshCredential()
	if err != nil {
		err = errors.Wrap(err, "can't refresh credential")
		if debugEnable {
			log.Debug(err)
		}
		return
	}
	if credentialResp.Code != float64(tool.RespCodeSuccess) {
		response.Code = credentialResp.Code
		response.Message = credentialResp.Message
		return
	}
	// 如果获取了新token,重新发一遍
	if debugEnable {
		log.Debug("re_send the message")
	}
	code, response, err = _send(m.credential, u, data)
	return
}
