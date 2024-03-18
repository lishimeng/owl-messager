package um

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/midware"
)

func Route(root server.Router) {
	root.Post("/{category:string}", midware.WithAuth(sendMessage)...)
}
