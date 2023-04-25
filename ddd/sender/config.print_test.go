package sender

import (
	"github.com/lishimeng/owl/internal/db/model"
	"testing"
)

func TestVendorConfig(t *testing.T) {
	var c, ok = vendorSupport[model.MailVendorMicrosoft.String()+model.SenderCategorySms]
	if !ok {
		t.Log("unknown config")
		return
	}
	var m = getJsonConstructor(c)
	for key, value := range m {
		t.Logf("%s:\t%s\n", key, value)
	}
}
