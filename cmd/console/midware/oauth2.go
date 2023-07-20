package midware

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/midware/auth"
)

// WithAuth token验证器,
// auth.JwtBasic header预处理
// auth.Forbidden401Handler 无权限时返回401, 返回格式按照参数 auth.ForbiddenOption
func WithAuth(handler func(iris.Context)) []iris.Handler {
	var handlers []iris.Handler
	handlers = append(handlers,
		auth.JwtBasic(),
		auth.Forbidden401Handler(
			auth.WithJsonResp(),
			//auth.WithScope(common.Scope),
		),
	)
	handlers = append(handlers, handler)
	return handlers
}
