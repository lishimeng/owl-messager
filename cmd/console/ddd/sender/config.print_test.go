package sender

import (
	"github.com/lishimeng/owl-messager/pkg/msg"
	"testing"
)

func TestVendorConfig(t *testing.T) {
	var c, ok = vendorSupport[msg.Microsoft.String()+msg.MailMessage.String()]
	if !ok {
		t.Log("unknown config")
		return
	}
	var m = getJsonConstructor(c)
	for key, value := range m {
		t.Logf("%s:\t%s\n", key, value)
	}
}
