package sender

import (
	"github.com/lishimeng/owl-messager/pkg/msg"
	"testing"
)

func TestVendorConfig(t *testing.T) {
	m := msg.ProviderConfigFields(msg.MailMessage, msg.Microsoft)
	if len(m) == 0 {
		t.Log("unknown config")
		return
	}
	for key, value := range m {
		t.Logf("%s:\t%s\n", key, value)
	}
}
