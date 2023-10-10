package sms

import (
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/sms"
	"strings"
)

// 七牛云SMS

type QiniuSdk struct {
	appKey    string
	appSecret string
	sign      string
	token     *auth.Credentials
	manager   *sms.Manager
}

func NewQiniu(appkey, appSecret string) (sdk messager.SmsProvider) {

	qiniu := QiniuSdk{
		appKey:    appkey,
		appSecret: appSecret,
	}
	sdk = &qiniu

	return
}

func (qiniu *QiniuSdk) Send(req messager.Request) (resp messager.Response, err error) {
	to := strings.Split(req.Receivers, ",")
	ret, err := qiniu.SendSms(qiniu.sign, req.Template.CloudTemplate, to, req.Params)
	if err != nil {
		return
	}
	resp.RequestId = ret.JobID
	return
}

func (qiniu *QiniuSdk) handleToken() {
	qiniu.token = auth.New(qiniu.appKey, qiniu.appSecret)
	qiniu.manager = sms.NewManager(qiniu.token)
}

func (qiniu *QiniuSdk) SendSms(signature, template string, to []string, params map[string]interface{}) (resp sms.MessagesResponse, err error) {
	resp, err = qiniu.manager.SendMessage(sms.MessagesRequest{
		SignatureID: signature,
		TemplateID:  template,
		Mobiles:     to,
		Parameters:  params,
	})
	return
}
