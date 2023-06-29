package um

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/owl-messager/internal/api/midware"
)

func Route(root iris.Party) {
	root.Post("/{category:string}", midware.WithAuth(sendMessage)...)
}
