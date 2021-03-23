package sender

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/provider/mail"
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
	si, err := repo.GetMailSenderByCode(p.SenderCode)
	if err != nil {
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
			Email:  si.Email,
			Name:   si.Alias,
			Passwd: si.Passwd,
		},
		Receiver: mail.MetaReceiver{
			To: toers,
		},
	}
	s := mail.New()
	err = s.Send(metas, p.Subject, p.Body)
	return
}
