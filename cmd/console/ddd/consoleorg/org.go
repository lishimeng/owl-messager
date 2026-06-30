package consoleorg

import "github.com/lishimeng/app-starter/server"

// TenantFilter returns optional tenant_code query filter; empty means all tenants (management console).
func TenantFilter(ctx server.Context) string {
	return ctx.C.URLParamDefault("tenantCode", "")
}
