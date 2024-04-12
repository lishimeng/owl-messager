package dict

import "github.com/lishimeng/app-starter/server"

func Router(r server.Router) {
	r.Get("/providers", providerList)
	r.Get("/providers/{category}", providerList)
	r.Get("/provider/{id}", providerList)
}
