package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd/templates"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd/um"
)

func Route(root server.Router) {
	mountOpenAPI(root)
	mountOpenAPI(root.Path("/api"))
}

func mountOpenAPI(root server.Router) {
	um.Route(root.Path("/messages"))
	templates.Route(root.Path("/template"))
}
