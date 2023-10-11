package sdk

import (
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/lishimeng/owl-messager/utils"
	"github.com/pkg/errors"
)

const ApiSendMessage = "/messages/"

const ApiCredential = "/open/oauth2/token"

const ApiTemplates = "/template"

const (
	CodeNotAllow int = 401
	CodeNotFound int = 404
	CodeSuccess  int = 200
)

// messageClient 消息服务
type messageClient struct {
	host       string // 消息服务主机地址. 如, "http://127.0.0.1"
	credential string
	appId      string
	secret     string
}

func (m *messageClient) SendMail(request MailRequest) (response Response, err error) {
	if debugEnable {
		log.Debug("sendMail to: %s", request.Receiver)
	}
	response, err = m.send(string(msg.MailMessage), request)
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
	response, err = m.send(string(msg.SmsMessage), request)
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
	response, err = m.send(string(msg.ApnsMessage), request)
	if err != nil {
		log.Debug(err)
		return
	}
	return
}

func (m *messageClient) Templates(request TemplateRequest) (resp TemplateResponse, err error) {
	if debugEnable {
		log.Debug("get templates: %s", request.Category)
	}
	req := make(map[string]string)
	if request.PageNo > 0 {
		req["pageNo"] = fmt.Sprintf("%d", request.PageNo)
	}
	if request.PageSize > 0 {
		req["pageSize"] = fmt.Sprintf("%d", request.PageSize)
	}

	err = NewRpc(m.host).Auth(m.appId, m.secret).BuildReq(func(rest *utils.RestClient) (int, error) {
		code, e := rest.Path(ApiTemplates, request.Category.String()).
			ResponseJson(&resp).Get(req)
		if resp.Code == CodeNotAllow { // 拦截业务异常
			code = CodeNotAllow
		}
		return code, e
	}).Exec()

	if err != nil {
		err = errors.Wrap(err, "get template fail")
		if debugEnable {
			log.Debug(err)
		}
		return
	}
	return
}

func (m *messageClient) send(category string, request any) (response Response, err error) {

	err = NewRpc(m.host).Auth(m.appId, m.secret).BuildReq(func(rest *utils.RestClient) (int, error) {
		code, e := rest.Path(ApiSendMessage, category).ResponseJson(&response).Post(request)
		if response.Code == CodeNotAllow { // 拦截业务异常
			code = CodeNotAllow
		}
		return code, e
	}).Exec()

	if err != nil {
		err = errors.Wrap(err, "send fail")
		if debugEnable {
			log.Debug(err)
		}
		return
	}

	return
}
