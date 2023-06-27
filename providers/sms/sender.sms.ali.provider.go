package sms

// 阿里云SMS

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/messager"
)

type AliProvider struct {
	accessKey    string
	accessSecret string
	region       string
	client       *dysmsapi.Client
	signName     string
}

var (
	aliyunHost = "dysmsapi.aliyuncs.com"
)

func (p *AliProvider) Send(req messager.Request) (resp messager.Response, err error) {

	to := req.Receivers
	var signature = p.signName // sender 中的signature优先级最低
	if len(req.Sign) > 0 {
		signature = req.Sign
	}
	ret, err := p.sendSms(to, signature, req.Template, req.Params)
	if err != nil {
		return
	}
	resp.RequestId = *ret.Body.RequestId
	resp.Payload = ret.String()
	return
}

func (p *AliProvider) Init(conf model.AliSmsConfig) (err error) {
	var accessKey = conf.AppKey
	var accessSecret = conf.AppSecret
	var region = conf.Region
	var signName = conf.SignName
	p.accessKey = accessKey
	p.accessSecret = accessSecret
	p.region = region
	p.signName = signName

	config := &openapi.Config{
		AccessKeyId:     &accessKey,
		AccessKeySecret: &accessSecret,
		RegionId:        &region,
	}
	config.Endpoint = &aliyunHost

	p.client, err = dysmsapi.NewClient(config)
	return
}

func (p *AliProvider) sendSms(receiver string, signName string, tplId string, tplParams string) (resp *dysmsapi.SendSmsResponse, err error) {
	var req dysmsapi.SendSmsRequest
	var request *dysmsapi.SendSmsRequest
	req.PhoneNumbers = &receiver
	req.SignName = &signName
	req.TemplateCode = &tplId
	req.TemplateParam = &tplParams

	request = req.SetPhoneNumbers(receiver).
		SetSignName(p.signName).
		SetTemplateCode(tplId).
		SetTemplateParam(tplParams)

	var opts util.RuntimeOptions
	opts.SetIgnoreSSL(true) // 忽略ssl验证
	resp, err = p.client.SendSmsWithOptions(request, &opts)
	if err != nil {
		log.Info("send sms failed(ali sdk)")
		log.Info(err)
		return
	}
	log.Info("response is %#v\n", resp)
	return
}
