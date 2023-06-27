package sender

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/internal/provider"
	"github.com/lishimeng/owl-messager/internal/provider/template"
	"strings"
)

type Mail interface {
	Send(model.MailMessageInfo) (err error)
}

type mailSender struct {
	ctx       context.Context
	maxWorker int
}

func NewMailSender(ctx context.Context) (m Mail, err error) {
	m = &mailSender{
		ctx:       ctx,
		maxWorker: 1,
	}
	return
}

func (m *mailSender) Send(p model.MailMessageInfo) (err error) {
	// sender info
	log.Info("send mail:%d", p.Id)

	si, err := repo.GetDefaultMailSender("") // 使用默认sender

	if err != nil {
		log.Info("mail sender not exist")
		return
	}

	var params map[string]interface{}
	err = json.Unmarshal([]byte(p.Params), &params)
	if err != nil {
		log.Info("params of mail template is not json format:%s", p.Params)
		return
	}

	var mailBody string

	if p.CloudTemplate == 0 { // 本地模板,解析邮件body
		mailBody, err = m.buildMailBody(p, params)
		if err != nil {
			log.Info("build mail body failure")
			return
		}
	}

	s, err := provider.DefaultMailFactory.Create(si.Vendor, si.Config)
	if err != nil {
		log.Info("create mail sender failure:%d", si.Id)
		return
	}

	receivers := strings.Split(p.Receivers, ",")

	req := messager.MailRequest{
		Subject:     p.Subject,
		TextContent: mailBody,
		Receivers:   receivers,
		Template:    p.CloudTemplateId,
		Params:      params,
	}

	err = s.Send(req)
	return
}

func (m *mailSender) buildMailBody(p model.MailMessageInfo, params map[string]interface{}) (body string, err error) {
	tpl, err := repo.GetMailTemplateById(p.Template)
	if err != nil {
		log.Info("mail template not exist:%d", p.Template)
		return
	}

	body, err = template.Rend(params, tpl.Body, tpl.Category)
	if err != nil {
		log.Info("template render failed")
		log.Info(err)
		return
	}
	if len(body) <= 0 {
		log.Info("mail body empty")
		err = errors.New("mail body empty")
		return
	}
	return
}
