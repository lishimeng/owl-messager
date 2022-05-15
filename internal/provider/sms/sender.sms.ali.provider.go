package sms

// 阿里云SMS

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/lishimeng/go-log"
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

func (p *AliProvider) Send(req Request) (resp Response, err error) {

	to := req.Receivers
	ret, err := p.SendSms(to, req.Template, req.Params)
	if err != nil {
		return
	}
	resp.RequestId = *ret.Body.RequestId
	resp.Payload = ret.String()
	return
}

func (p *AliProvider) Init(accessKey string, accessSecret string, region string, signName string) (err error) {
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

func (p AliProvider) SendSms(receiver string, tplId string, tplParams string) (resp *dysmsapi.SendSmsResponse, err error) {
	var req dysmsapi.SendSmsRequest
	var request *dysmsapi.SendSmsRequest
	req.PhoneNumbers = &receiver
	req.SignName = &p.signName
	req.TemplateCode = &tplId
	req.TemplateParam = &tplParams

	request = req.SetPhoneNumbers(receiver).
		SetSignName(p.signName).
		SetTemplateCode(tplId).
		SetTemplateParam(tplParams)

	resp, err = p.client.SendSms(request)
	if err != nil {
		log.Info("send sms failed(ali sdk)")
		log.Info(err)
		return
	}
	log.Info("response is %#v\n", resp)
	return
}
