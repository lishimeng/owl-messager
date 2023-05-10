package sender

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {

	root.Get("/", list) // 列表
	root.Get("/mail/{id}", mailSenderInfo)
	root.Get("/sms/{id}", smsSenderInfo)
	root.Get("/apns/{id}", apnsSenderInfo)
	root.Post("/{org}/{category}/{id:int}/set_default", setDefaultSender) // 设置组织下默认发信账号
	root.Get("/config/{category}/{vendor}", getConfigStruct)              // 列出配置字段
}
