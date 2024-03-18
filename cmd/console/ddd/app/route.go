package tenant

import (
	"github.com/lishimeng/app-starter/server"
)

func Route(root server.Router) {
	root.Post("/", add)
}

const (
	CodeErr = 100505
)
