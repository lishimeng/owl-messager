package sms

import (
	"encoding/json"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/messager"
	"strings"
)

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

// 腾讯云SMS

type TencentSdk struct {
	config model.TencentSmsConfig
	client *sms.Client
}

func NewTencent(conf model.TencentSmsConfig) (sdk messager.SmsProvider, err error) {

	credential := common.NewCredential(conf.AppId, conf.AppKey)

	client, err := sms.NewClient(credential, conf.Region, profile.NewClientProfile())

	if err != nil {
		return
	}
	//client.WithDebug(true)
	sdk = &TencentSdk{
		config: conf,
		client: client,
	}
	return
}

func (sdk *TencentSdk) Send(message messager.Request) (resp messager.Response, err error) {
	to := strings.Split(message.Receivers, ",")
	var m = make(map[string]interface{})
	err = json.Unmarshal([]byte(message.Params), &m)
	if err != nil {
		// TODO
		return
	}
	params := map2array(m)

	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = common.StringPtr(sdk.config.SmsAppId)
	req.SignName = common.StringPtr(sdk.config.SignName)

	/* 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	 * 示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	req.PhoneNumberSet = common.StringPtrs(to)

	req.TemplateId = common.StringPtr(message.Template)
	if len(params) > 0 {
		req.TemplateParamSet = common.StringPtrs(params)
	}

	bs, _ := json.Marshal(req)
	log.Debug(string(bs))

	result, err := sdk.client.SendSms(req)
	if err != nil {
		return
	}
	resp.RequestId = *result.Response.RequestId
	resp.Payload = result.ToJsonString()
	return
}
