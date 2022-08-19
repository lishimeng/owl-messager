package sender

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/provider/mail"
	"github.com/lishimeng/owl/internal/provider/template"
	"strings"
)

type Mail interface {
	Send(model.MailMessageInfo) (err error)
}

type mailSender struct {
	ctx context.Context

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
	var si model.MailSenderInfo
	if p.Sender > 0 {
		si, err = repo.GetMailSenderById(p.Sender) // 使用指定的sender
	} else {
		si, err = repo.GetDefaultMailSender("") // 使用默认sender
	}
	if err != nil {
		log.Info("mail sender not exist:%d", p.Sender)
		return
	}

	toers := strings.Split(p.Receivers, ",")

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
	if len(mailBody) > 0 {

	}
	s, err := mail.F.Create(si.Vendor, si.Config)
	if err != nil {
		log.Info("create mail sender failure:%d", si.Id)
		return
	}
	err = s.Send(p.Subject, mailBody, toers...)
	return
}
