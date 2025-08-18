package sender

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/etc"
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

func testImSender(ctx server.Context) {
	var resp sdk.Response
	var req sdk.ImRequest
	resp.Code = tool.RespCodeSuccess
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}

	// 请求代理
	resp, err = sdk.New(
		sdk.WithAuth(etc.Config.Console.AppKey, etc.Config.Console.Secret),
		sdk.WithHost(etc.Config.Console.Host)).
		SendIm(req)

	if err != nil {
		resp.Code = tool.RespCodeError
	}

	if resp.Code == -1 {
		resp.Code = tool.RespCodeError
	}

	ctx.Json(resp)
}

func testSmsSender(ctx server.Context) {
	var resp sdk.Response
	var req sdk.SmsRequest
	resp.Code = tool.RespCodeSuccess
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	req.TemplateParam = map[string]interface{}{
		"code": "123456",
	}

	// 请求代理
	resp, err = sdk.New(
		sdk.WithAuth("1b9f6d187d196025e5a57441a1175bea13b6ce74b47de0c5424d6d36ebfcbe13", "910136561f366c30192282fb5459b17bac594f2c1bd6ed1731b6a04d28bf9c6e"),
		sdk.WithHost("http://127.0.0.1:91")).
		SendSms(req)

	if err != nil {
		resp.Code = tool.RespCodeError
	}

	if resp.Code == -1 {
		resp.Code = tool.RespCodeError
	}

	ctx.Json(resp)
}
