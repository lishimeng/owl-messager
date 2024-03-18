package sender

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/owl-messager/cmd/console/midware"
)

func Route(root server.Router) {

	root.Get("/", midware.WithAuth(list)...) // 列表
	root.Get("/mail/{id}", mailSenderInfo)
	root.Get("/sms/{id}", smsSenderInfo)
	root.Get("/apns/{id}", apnsSenderInfo)
	root.Post("/{org}/{category}/{id:int}/set_default", setDefaultSender) // 设置组织下默认发信账号
	root.Get("/config/{category}/{vendor}", getConfigStruct)              // 列出配置字段
	root.Post("/mail/set_default", SetMailSenderInfo)                     //新增mail配置
	root.Post("/mail/up_default", UpMailSenderInfo)                       //编辑mail配置
	root.Get("/mail/vendor", GetMailSenderInfo)                           //获取mail配置
	root.Get("/mail/list/page", midware.WithAuth(ListByPage)...)          //获取列表
	root.Get("/mail/info/category", GetSenderInfoByCategory)              //获取列表
}
