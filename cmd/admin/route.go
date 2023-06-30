package admin

import (
	"github.com/kataras/iris/v12"
	tenant "github.com/lishimeng/owl-messager/cmd/admin/ddd/app"
)

func Route(p iris.Party) {
	// admin下的资源
	tenant.Route(p.Party("/tenant"))
}
