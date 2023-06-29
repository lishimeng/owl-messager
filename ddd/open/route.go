package open

import (
	"github.com/kataras/iris/v12"
)

func Route(root iris.Party) {
	credentials(root.Party("/oauth2"))
}

func credentials(root iris.Party) {
	root.Post("/token", genCredential)
}
