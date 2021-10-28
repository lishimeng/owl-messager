package openApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/openapi"
)

func Token(ctx iris.Context) {

	err := openapi.Srv.HandleTokenRequest(ctx.ResponseWriter(), ctx.Request())
	if err != nil {
		log.Info(err)
		// TODO

	}
}

func Authorize(ctx iris.Context) {

	err := openapi.Srv.HandleAuthorizeRequest(ctx.ResponseWriter(), ctx.Request())
	if err != nil {
		log.Info(err)
	}
}
