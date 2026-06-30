package midware

import (
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/midware/auth/bearer"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

func skipApiAuth(path string) bool {
	return strings.HasPrefix(path, "/api/auth")
}

// MountApiAuth protects /api routes with Bearer basic_auth (k8s dashboard style).
func MountApiAuth(r server.Router) {
	r.Party().Use(irisConsoleAuth)
}

func irisConsoleAuth(ctx iris.Context) {
	if skipApiAuth(ctx.Path()) {
		ctx.Next()
		return
	}
	consoleBearerAuth(server.Context{C: ctx})
}

func consoleBearerAuth(ctx server.Context) {
	raw, ok := bearer.GetAuth(ctx)
	if !ok {
		denyConsole(ctx)
		return
	}

	client, err := repo.GetClientByBasicAuth(raw)
	if err != nil {
		log.Debug("console auth rejected")
		denyConsole(ctx)
		return
	}

	p := token.JwtPayload{
		Uid: client.AppId,
		Org: client.TenantCode,
	}
	ctx.C.Values().Set(auth.UserInfoKey, p)
	r := ctx.C.Request()
	r.Header.Set(auth.OrgKey, client.TenantCode)
	r.Header.Set(auth.UidKey, client.AppId)
	ctx.C.Next()
}

func denyConsole(ctx server.Context) {
	var resp auth.Response
	resp.Code = iris.StatusUnauthorized
	resp.Message = auth.ErrNotAllowed.Error()
	ctx.Json(resp)
}
