package consoleorg

import (
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

const DefaultOrg = 1

// ID resolves tenant org id from request header; falls back to DefaultOrg.
func ID(ctx server.Context) int {
	code := ctx.C.GetHeader(auth.OrgKey)
	if code == "" {
		return DefaultOrg
	}
	tenant, err := repo.GetTenant(code)
	if err != nil {
		return DefaultOrg
	}
	return tenant.Id
}
