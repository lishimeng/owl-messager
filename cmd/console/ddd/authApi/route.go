package authApi

import "github.com/lishimeng/app-starter/server"

func Route(root server.Router) {
	root.Post("/token", verifyToken)
}
