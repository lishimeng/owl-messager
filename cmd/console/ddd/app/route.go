package tenant

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {
	root.Post("/", add)
}

const (
	CodeErr = 100505
)
