package mail

import (
	"errors"
	"github.com/lishimeng/owl/internal/provider/mail/ms"
)

type msSender struct {
	proxy *ms.AzureGraphProvider
}

func (h *msSender) Init(config map[string]string) {

}

func (h *msSender) Send(metas MetaInfo, subject string, body string) (err error) {

	if h.proxy == nil {
		err = errors.New("proxy nil")
	}

	err = h.proxy.Send(subject, body, metas.Receiver.To...)
	return
}
