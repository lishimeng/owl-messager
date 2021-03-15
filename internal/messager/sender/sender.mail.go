package sender

import (
	"context"
	"fmt"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/provider/mail"
)

type Mail interface {

}

type mailSender struct {
	ctx context.Context

	payloadCh chan interface{}

	maxWorker int
}

func NewMailSender(ctx context.Context, maxWorker int) {

}

func (m *mailSender) work() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case p := <-m.payloadCh:
			e := m.send(p)
			if e != nil {
				fmt.Println("send mail failed")
				fmt.Println(e)
			}
		}
	}
}

func (m *mailSender) send(p interface{}) (err error) {
	// sender info
	si, err := repo.GetMailSenderByCode("")
	if err != nil {
		return
	}
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
			To: []string{""},
		},
	}
	s := mail.New()
	err = s.Send(metas, "", "")
	return
}