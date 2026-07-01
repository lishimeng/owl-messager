package um

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/midware"
)

func Route(root server.Router) {
	root.Get("/ping", ping)
	root.Post("/{category:string}", midware.WithOpenAuth(sendMessage)...)
	// 勿使用 /mail/attachments：会与 POST /messages/mail 路径冲突导致 404
	root.Post("/attachments", midware.WithOpenAuth(uploadMailAttachment)...)
}

func ping(ctx server.Context) {
	var resp app.Response
	resp.Code = tool.RespCodeSuccess
	resp.Message = "owl-messager"
	ctx.Json(resp)
}
