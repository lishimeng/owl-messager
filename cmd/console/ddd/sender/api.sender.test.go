package sender

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/sdk"
)

func testMailSender(ctx server.Context) {
	var resp sdk.Response
	var req sdk.MailRequest
	resp.Code = tool.RespCodeSuccess
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}

	// 请求代理
	resp, err = sdk.New(
		sdk.WithAuth("test", "secret_code"),
		sdk.WithHost("http://127.0.0.1:91")).
		SendMail(req)

	if err != nil {
		resp.Code = tool.RespCodeError
	}

	if resp.Code == -1 {
		resp.Code = tool.RespCodeError
	}

	ctx.Json(resp)
}
