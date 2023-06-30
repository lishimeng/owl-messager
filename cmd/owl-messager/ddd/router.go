package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd/open"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd/um"
)

func Route(root *iris.Application) {

	p := root.Party("/")
	um.Route(p.Party("/messages"))
	open.Route(p.Party("/open"))
}
