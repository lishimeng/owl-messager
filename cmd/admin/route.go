package admin

import (
	"github.com/lishimeng/app-starter/server"
	tenant "github.com/lishimeng/owl-messager/cmd/admin/ddd/app"
)

func Route(p server.Router) {
	// admin下的资源
	tenant.Route(p.Path("/tenant"))
}
