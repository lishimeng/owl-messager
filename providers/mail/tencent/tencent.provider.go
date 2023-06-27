package tencent

import (
	"encoding/json"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ses "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses/v20201002"
	"strconv"
)

type Template struct {
	TemplateID   string
	TemplateData string
}

func newTencentClient(appId, secret, region string) (client *ses.Client, err error) {
	credential := common.NewCredential(appId, secret)
	prof := profile.NewClientProfile()
	prof.Debug = true
	//prof.DisableRegionBreaker = false
	client, err = ses.NewClient(credential, region, profile.NewClientProfile())

	return
}

type MailTencentProvider struct {
	Config model.TencentConfig
	Client *ses.Client
}

func New(config model.TencentConfig) (s *MailTencentProvider, err error) {
	s = &MailTencentProvider{}
	client, err := newTencentClient(config.AppId, config.Secret, config.Region)
	if err != nil {
		log.Info("create tencent client err")
		log.Info(err)

		return
	}
	s.Config = config
	s.Client = client
	return
}

// Send
//
// template为腾讯云平台上创建的邮件模板ID不是名称
func (s *MailTencentProvider) Send(param messager.MailRequest) (err error) {
	req := ses.NewSendEmailRequest()

	req.FromEmailAddress = &s.Config.Sender
	for _, receiver := range param.Receivers {
		var r = receiver
		req.Destination = append(req.Destination, &r)
	}

	req.Subject = &param.Subject
	tid, err := strconv.ParseUint(param.Template, 10, 64)
	if err != nil {
		return
	}

	bs, err := json.Marshal(param.Params)
	if err != nil {
		return
	}

	var params = string(bs)

	req.Template = &ses.Template{
		TemplateID:   &tid,
		TemplateData: &params,
	}
	err = s.send(req)
	return
}

func (s *MailTencentProvider) send(request *ses.SendEmailRequest) (err error) {

	resp, err := s.Client.SendEmail(request)
	if err != nil {
		log.Info(err)
		return
	}
	log.Info(resp.ToJsonString())
	return
}
