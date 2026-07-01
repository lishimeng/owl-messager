package sender

import (
	"github.com/lishimeng/app-starter/server"
)

func Route(root server.Router) {

	root.Get("/", list)
	root.Get("/mail/{id}", mailSenderInfo)
	root.Get("/sms/{id}", smsSenderInfo)
	root.Get("/apns/{id}", apnsSenderInfo)
	root.Post("/set_default", setDefaultSender)              // 设置组织下默认发信账号
	root.Get("/config/{category}/{vendor}", getConfigStruct) // 列出配置字段
	root.Post("/mail/set_default", SetMailSenderInfo)        //新增mail配置
	root.Post("/mail/up_default", UpMailSenderInfo)          //编辑mail配置
	root.Post("/mail/del", DeleteSender)                     //删除发送人
	root.Get("/mail/vendor", GetMailSenderInfo)              //获取mail配置
	root.Get("/mail/list/page", ListByPage)                  //获取列表
	root.Get("/mail/info/category", GetSenderInfoByCategory) //获取配置
	root.Get("/test/config", getTestConfig)                  //测试发送默认配置
	root.Post("/test/mail", testMailSender)                  //测试请求，需要转发到owl-messager
	root.Post("/test/im", testImSender)                      //测试请求，需要转发到owl-messager
	root.Post("/test/sms", testSmsSender)                    //测试请求，需要转发到owl-messager
}
