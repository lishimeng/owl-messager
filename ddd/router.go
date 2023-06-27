package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/owl-messager/ddd/sender"
	"github.com/lishimeng/owl-messager/ddd/um"
)

func Router(root iris.Party) {

	um.Route(root.Party("/messages"))
	sender.Route(root.Party("/sender"))
}
