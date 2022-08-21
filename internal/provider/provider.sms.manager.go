package provider

import (
	"encoding/json"
	"errors"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/provider/sms"
)

type SmsFactory struct {
}

var DefaultSmsFactory *SmsFactory

func init() {
	DefaultSmsFactory = &SmsFactory{}
}

func (f *SmsFactory) Create(vendor model.SmsVendor, config string) (p sms.Provider, err error) {

	switch vendor {
	case model.SmsVendorAli:
		var aliSmsConf model.AliSmsConfig
		h := sms.AliProvider{}
		err = json.Unmarshal([]byte(config), &aliSmsConf)
		if err != nil {
			return
		}
		err = h.Init(aliSmsConf)
		if err != nil {
			return
		}
		p = &h
	default:
		err = errors.New("unknown mail vendor")
	}

	log.Info("create sms provider")

	return
}
