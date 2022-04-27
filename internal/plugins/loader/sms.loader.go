package loader

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/plugins/container"
	"github.com/lishimeng/owl/internal/provider/sms"
)

func init() {
	for {
		if container.Ready() {
			initSmsProviderManager()
			return
		}
	}
}

func initSmsProviderManager() {
	// 注册provider manager
	var pm = sms.New()
	container.Add(&pm)
}

// Load 加载一个sms provider
func Load(id string, p sms.Provider) {
	var manager, err = container.Get(new(sms.ProviderManager))
	if err != nil {
		log.Info(err)
		return
	}
	manager.Add(model.SmsVendor(id), p)
	return
}
