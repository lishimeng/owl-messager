package sms

import (
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/sms"
)

// 七牛云SMS

type QiniuSdk struct {
	appKey    string
	appSecret string

	token *auth.Credentials

	manager *sms.Manager
}

func NewQiniu(appkey, appSecret string) (sdk Provider) {

	qiniu := QiniuSdk{}
	sdk = &qiniu

	return
}

func (qiniu *QiniuSdk) Send(req Request) (resp Response, err error) {
	ret, err := qiniu.SendSms(req.Sign, req.Template, req.Receivers, req.Params)
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
