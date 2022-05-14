package sender

import (
	"context"
	"encoding/json"
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
	//si, err := repo.GetMailSenderById(p.Sender)
	si, err := repo.GetDefaultMailSender("")
	if err != nil {
		log.Info("mail sender not exist:%d", p.Sender)
		return
	}

	toers := strings.Split(p.Receivers, ",")
	// TODO delete toer:""
	metas := mail.MetaInfo{
		Server: mail.MetaServer{
			Host: si.Host,
			Port: si.Port,
		},
		Sender: mail.MetaSender{
			Email:      si.Email,
			Name:       si.Alias,
			Passwd:     si.Passwd,
			EmailAlias: si.EmailAlias,
		},
		Receiver: mail.MetaReceiver{
			To: toers,
		},
	}

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
	c, err := template.Rend(params, tpl.Body, tpl.Category)
	if err != nil {
		log.Info("template render failed")
		log.Info(err)
		return
	}
	if len(c) > 0 {
		s := mail.New()
		err = s.Send(metas, p.Subject, c)
	}
	return
}
