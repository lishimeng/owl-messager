package midware

import (
	"encoding/base64"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/midware/auth/bearer"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/common"
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

	p := token.JwtPayload{
		Uid:   client.AppId,
		Org:   client.TenantCode,
		Scope: common.Scope,
	}
	ctx.C.Values().Set(auth.UserInfoKey, p)
	r := ctx.C.Request()
	r.Header.Set(auth.OrgKey, client.TenantCode)
	r.Header.Set(auth.UidKey, client.AppId)
	r.Header.Set(auth.ScopeKey, p.Scope)
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
