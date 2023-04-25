package sender

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {

	root.Get("/", list)
	root.Get("/mail/{id}", mailSenderInfo)
	root.Get("/sms/{id}", smsSenderInfo)
	root.Get("/apns/{id}", apnsSenderInfo)
	root.Post("/{org}/{category}/{id:int}/set_default", setDefaultSender)
	root.Get("/config/{category}/{vendor}", vendorConfig)
}
