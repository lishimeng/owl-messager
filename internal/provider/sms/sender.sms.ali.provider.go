package sms
// 阿里云SMS

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/lishimeng/go-log"
)

type Provider struct {
	accessKey    string
	accessSecret string
	region       string
	client *dysmsapi.Client
	signName string
}

func (p *Provider) Init(accessKey string, accessSecret string, region string, signName string) (err error) {
	p.accessKey = accessKey
	p.accessSecret = accessSecret
	p.region = region
	p.signName = signName
	p.client, err = dysmsapi.NewClientWithAccessKey(p.region, p.accessKey, p.accessSecret)
	return
}

func (p Provider) Send(toer string, tplId string, tplParams string) (err error) {
	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = toer
	req.SignName = ""
	req.TemplateCode = tplId
	req.TemplateParam = tplParams

	resp, err := p.client.SendSms(req)
	if err != nil {
		log.Info("send sms failed(ali sdk)")
		log.Info(err)
		return
	}
	log.Info("response is %#v\n", resp)
	return
}