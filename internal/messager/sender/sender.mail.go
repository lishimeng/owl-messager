package sender

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/provider"
	"github.com/lishimeng/owl/internal/provider/template"
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

	mailBody, err := m.buildMailBody(p)
	if err != nil {
		log.Info("build mail body failure")
		return
	}

	s, err := provider.DefaultMailFactory.Create(si.Vendor, si.Config)
	if err != nil {
		log.Info("create mail sender failure:%d", si.Id)
		return
	}

	receivers := strings.Split(p.Receivers, ",")

	err = s.Send(p.Subject, mailBody, receivers...)
	return
}

func (m *mailSender) buildMailBody(p model.MailMessageInfo) (body string, err error) {
	tpl, err := repo.GetMailTemplateById(p.Template)
	if err != nil {
		log.Info("mail template not exist:%d", p.Template)
		return
	}
	var params map[string]interface{}
	err = json.Unmarshal([]byte(p.Params), &params)
	if err != nil {
		log.Info("params of mail template is not json format:%s", p.Params)
		return
	}
	mailBody, err := template.Rend(params, tpl.Body, tpl.Category)
	if err != nil {
		log.Info("template render failed")
		log.Info(err)
		return
	}
	if len(mailBody) <= 0 {
		log.Info("mail body empty")
		err = errors.New("mail body empty")
		return
	}
	return
}
