package um

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/midware"
)

func Route(root server.Router) {
	root.Post("/mail/attachments", midware.WithOpenAuth(uploadMailAttachment)...)
	root.Post("/{category:string}", midware.WithOpenAuth(sendMessage)...)
}
