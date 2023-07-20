package open

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/midware"
)

func Route(root iris.Party) {
	credentials(root.Party("/oauth2"))
}

func credentials(root iris.Party) {
	root.Post("/token", genCredential)
	root.Get("/token", midware.WithAuth(verify)...)
}
