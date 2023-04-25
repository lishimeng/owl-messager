package sender

import (
	"github.com/lishimeng/owl/internal/db/model"
	"testing"
)

func TestVendorConfig(t *testing.T) {
	var c = vendorSupport[model.MailVendorMicrosoft.String()+model.SenderCategoryMail]
	var m = getJsonConstructor(c)
	for key, value := range m {
		t.Logf("%s:\t%s\n", key, value)
	}
}
