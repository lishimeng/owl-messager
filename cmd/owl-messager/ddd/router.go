package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd/open"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd/templates"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd/um"
)

func Route(root server.Router) {

	p := root.Path("/")
	um.Route(p.Path("/messages"))
	open.Route(p.Path("/open"))

	templates.Route(p.Path("/template"))
}
