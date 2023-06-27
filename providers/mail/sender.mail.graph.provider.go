package mail

import (
	"errors"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/providers/mail/ms"
)

type msSender struct {
	proxy *ms.AzureGraphProvider
}

func NewMicrosoft(config string) (s messager.MailProvider, err error) {
	var p *ms.AzureGraphProvider
	var h = msSender{}
	p, err = ms.New(config)
	h.proxy = p
	s = &h
	return
}

func (h *msSender) Send(req messager.MailRequest) (err error) {

	if h.proxy == nil {
		err = errors.New("proxy nil")
	}

	err = h.proxy.Send(req.Subject, req.TextContent, req.Receivers...)
	return
}
