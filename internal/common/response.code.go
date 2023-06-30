package common

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
)

// 业务回复码

// OAuth
const (
	CodeAppNotFound    = 100404
	CodeSecretNotValid = 100401

	CodeParamNil = 1009001
)

const (
	MsgAppNotFound    = "App not found"
	MsgSecretNotValid = "App or Secret not valid"
	MsgParamNil       = "Param lost"
)

func RespLostParam(param string, ctx iris.Context) {
	var resp app.Response
	resp.Code = CodeParamNil
	resp.Message = fmt.Sprintf("%s:%s", MsgParamNil, param)
	tool.ResponseJSON(ctx, resp)
}
