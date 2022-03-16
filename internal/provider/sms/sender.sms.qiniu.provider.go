package sms

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/sms"
)

// 七牛云SMS

type QiniuSdk interface {
}

type qiniuSdkImpl struct {
	appKey    string
	appSecret string

	token *auth.Credentials

	manager *sms.Manager
}

func NewQiniu(appkey, appSecret string) (sdk QiniuSdk) {
	return
}

func (qiniu *qiniuSdkImpl) handleToken() {
	qiniu.token = auth.New(qiniu.appKey, qiniu.appSecret)
	qiniu.manager = sms.NewManager(qiniu.token)
}

func (qiniu *qiniuSdkImpl) SendSms(signature, template string, to []string, params map[string]interface{}) (err error) {
	resp, err := qiniu.manager.SendMessage(sms.MessagesRequest{
		SignatureID: signature,
		TemplateID:  template,
		Mobiles:     to,
		Parameters:  params,
	})
	if err != nil {
		return
	}
	fmt.Println(resp.JobID)
	return
}
