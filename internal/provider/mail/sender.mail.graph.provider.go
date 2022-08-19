package mail

import (
	"errors"
	"github.com/lishimeng/owl/internal/provider/mail/email"
	"github.com/lishimeng/owl/internal/provider/mail/ms"
)

type msSender struct {
	proxy *ms.AzureGraphProvider
}

func NewMicrosoft(config string) (s email.Sender, err error) {
	var p *ms.AzureGraphProvider
	var h = msSender{}
	p, err = ms.New(config)
	h.proxy = p
	s = &h
	return
}

func (h *msSender) Send(subject string, body string, to ...string) (err error) {

	if h.proxy == nil {
		err = errors.New("proxy nil")
	}

	err = h.proxy.Send(subject, body, to...)
	return
}
