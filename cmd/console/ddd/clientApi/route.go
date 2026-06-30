package clientApi

import "github.com/lishimeng/app-starter/server"

func Route(root server.Router) {
	root.Get("/list", getClientByPage)     // 获取列表
	root.Post("/create", createClient)     // 创建客户端
	root.Post("/del", deleteClient)        // 删除客户端
	root.Get("/secret", queryClientSecret) // 查询 app secret
}
