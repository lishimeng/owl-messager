package midware

import (
	"errors"

	"github.com/lishimeng/app-starter/server"
)

const (
	openAppIdKey      = "open_app_id"
	openTenantCodeKey = "open_tenant_code"
)

var ErrOpenAuthRequired = errors.New("open auth required")

func setOpenClientContext(ctx server.Context, appId, tenantCode string) {
	ctx.C.Values().Set(openAppIdKey, appId)
	ctx.C.Values().Set(openTenantCodeKey, tenantCode)
}

// TenantCodeFromContext returns tenant_code bound to the authenticated OpenClient.
func TenantCodeFromContext(ctx server.Context) (string, error) {
	v := ctx.C.Values().Get(openTenantCodeKey)
	code, ok := v.(string)
	if !ok || code == "" {
		return "", ErrOpenAuthRequired
	}
	return code, nil
}

// AppIdFromContext returns app_id of the authenticated OpenClient.
func AppIdFromContext(ctx server.Context) (string, error) {
	v := ctx.C.Values().Get(openAppIdKey)
	appId, ok := v.(string)
	if !ok || appId == "" {
		return "", ErrOpenAuthRequired
	}
	return appId, nil
}
