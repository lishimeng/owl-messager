package um

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/midware"
)

func Route(root iris.Party) {
	root.Post("/{category:string}", midware.WithAuth(sendMessage)...)
}
