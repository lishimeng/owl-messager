package sms

// 阿里云SMS

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/lishimeng/go-log"
)

type AliProvider struct {
	accessKey    string
	accessSecret string
	region       string
	client       *dysmsapi.Client
	signName     string
}

func (p *AliProvider) Send(req Request) (resp Response, err error) {

	to := ""     // TODO 转换参数
	params := "" // TODO 转换参数
	ret, err := p.SendSms(to, req.Template, params)
	if err != nil {
		return
	}
	resp.RequestId = ret.RequestId
	resp.Payload, err = json.Marshal(ret)
	return
}

func (p *AliProvider) Init(accessKey string, accessSecret string, region string, signName string) (err error) {
	p.accessKey = accessKey
	p.accessSecret = accessSecret
	p.region = region
	p.signName = signName
	p.client, err = dysmsapi.NewClientWithAccessKey(p.region, p.accessKey, p.accessSecret)
	return
}

func (p AliProvider) SendSms(toer string, tplId string, tplParams string) (resp *dysmsapi.SendSmsResponse, err error) {
	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = toer
	req.SignName = p.signName
	req.TemplateCode = tplId
	req.TemplateParam = tplParams

	resp, err = p.client.SendSms(req)
	if err != nil {
		log.Info("send sms failed(ali sdk)")
		log.Info(err)
		return
	}
	log.Info("response is %#v\n", resp)
	return
}
