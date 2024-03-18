package templates

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/midware"
)

func Route(root server.Router) {
	root.Get("/{category}", midware.WithAuth(templates)...)
}
