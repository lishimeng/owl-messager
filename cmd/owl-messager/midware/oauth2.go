package midware

import (
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/owl-messager/internal/common"
	"github.com/lishimeng/owl-messager/internal/etc"
)

// WithAuth token验证器,
// auth.JwtBasic header预处理
// auth.Forbidden401Handler 无权限时返回401, 返回格式按照参数 auth.ForbiddenOption
func WithAuth(handler func(server.Context)) []server.Handler {
	var handlers []server.Handler
	if etc.Config.Token.Enable {
		handlers = append(handlers,
			auth.JwtBasic(),
			auth.Forbidden401Handler(auth.WithJsonResp(), auth.WithScope(common.Scope)),
		)
	}
	handlers = append(handlers, handler)
	return handlers
}
