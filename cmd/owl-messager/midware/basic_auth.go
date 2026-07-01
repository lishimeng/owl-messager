package midware

import (
	"encoding/base64"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/midware/auth/bearer"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

func parseBasicAuth(header string) (user, pass string, ok bool) {
	if !strings.HasPrefix(header, "Basic ") {
		return
	}
	raw, err := base64.StdEncoding.DecodeString(header[6:])
	if err != nil {
		return
	}
	parts := strings.SplitN(string(raw), ":", 2)
	if len(parts) != 2 || parts[0] == "" {
		return
	}
	return parts[0], parts[1], true
}

func openBasicAuth(ctx server.Context) {
	appId, secret, ok := parseBasicAuth(ctx.C.GetHeader(bearer.AuthHeader))
	if !ok {
		denyOpen(ctx)
		return
	}

	client, err := repo.GetClientByAppId(appId)
	if err != nil || client.Secret != secret {
		log.Debug("open basic auth rejected: %s", appId)
		denyOpen(ctx)
		return
	}
	if client.TenantCode == "" {
		log.Debug("open client missing tenant_code: %s", appId)
		denyOpen(ctx)
		return
	}
	if _, err = repo.GetTenant(client.TenantCode); err != nil {
		log.Debug("open basic auth tenant disabled or missing: %s", client.TenantCode)
		denyOpen(ctx)
		return
	}

	setOpenClientContext(ctx, client.AppId, client.TenantCode)
	ctx.C.Next()
}

func denyOpen(ctx server.Context) {
	var resp auth.Response
	resp.Code = iris.StatusUnauthorized
	resp.Message = auth.ErrNotAllowed.Error()
	ctx.Json(resp)
}

// WithOpenAuth validates Authorization: Basic base64(appId:secret) on each request.
func WithOpenAuth(handler func(server.Context)) []server.Handler {
	return []server.Handler{openBasicAuth, handler}
}
