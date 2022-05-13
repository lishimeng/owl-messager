package sms

// 阿里云SMS

import (
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
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

	to := req.Receivers
	bs, _ := json.Marshal(req.Params)
	params := string(bs)
	ret, err := p.SendSms(to, req.Template, params)
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
		// 您的AccessKey ID
		AccessKeyId: &accessKey,
		// 您的AccessKey Secret
		AccessKeySecret: &accessSecret,
		RegionId:        &region,
	}
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")

	p.client, err = dysmsapi.NewClient(config)
	//p.client, err = dysmsapi.NewClientWithAccessKey(p.region, p.accessKey, p.accessSecret)
	return
}

func (p AliProvider) SendSms(receiver string, tplId string, tplParams string) (resp *dysmsapi.SendSmsResponse, err error) {
	var req *dysmsapi.SendSmsRequest
	//req.Scheme = "https"
	req.PhoneNumbers = &receiver
	req.SignName = &p.signName
	req.TemplateCode = &tplId
	req.TemplateParam = &tplParams

	req = req.SetPhoneNumbers(receiver).
		SetSignName(p.signName).
		SetTemplateCode(tplId).
		SetTemplateParam(tplParams)

	resp, err = p.client.SendSms(req)
	if err != nil {
		log.Info("send sms failed(ali sdk)")
		log.Info(err)
		return
	}
	log.Info("response is %#v\n", resp)
	return
}
