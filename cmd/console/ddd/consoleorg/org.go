package consoleorg

import (
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
)

const DefaultTenantCode = "default"

// Code returns tenant.code from request header (JWT / auth middleware).
func Code(ctx server.Context) string {
	code := ctx.C.GetHeader(auth.OrgKey)
	if code == "" {
		return DefaultTenantCode
	}
	return code
}
