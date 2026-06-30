package sender

import (
	"errors"

	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/sdk"
)

type messagerProxy struct {
	Host   string `json:"host"`
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
}

func (p messagerProxy) validate() error {
	if len(p.Host) == 0 || len(p.AppId) == 0 || len(p.Secret) == 0 {
		return errors.New("host, appId and secret required")
	}
	return nil
}

func (p messagerProxy) client() sdk.Client {
	return sdk.New(
		sdk.WithAuth(p.AppId, p.Secret),
		sdk.WithHost(p.Host),
	)
}

type testMailReq struct {
	messagerProxy
	sdk.MailRequest
}

type testImReq struct {
	messagerProxy
	sdk.ImRequest
}

type testSmsReq struct {
	messagerProxy
	sdk.SmsRequest
}

func testMailSender(ctx server.Context) {
	var req testMailReq

	if err := ctx.C.ReadJSON(&req); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	if err := req.validate(); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp, err := req.client().SendMail(req.MailRequest)
	writeTestResp(ctx, resp, err)
}

func testImSender(ctx server.Context) {
	var req testImReq

	if err := ctx.C.ReadJSON(&req); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	if err := req.validate(); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp, err := req.client().SendIm(req.ImRequest)
	writeTestResp(ctx, resp, err)
}

func testSmsSender(ctx server.Context) {
	var req testSmsReq

	if err := ctx.C.ReadJSON(&req); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	if err := req.validate(); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	if req.TemplateParam == nil {
		req.TemplateParam = map[string]interface{}{"code": "123456"}
	}

	resp, err := req.client().SendSms(req.SmsRequest)
	writeTestResp(ctx, resp, err)
}

func writeTestResp(ctx server.Context, resp sdk.Response, err error) {
	if err != nil {
		resp.Code = tool.RespCodeError
	}
	if resp.Code == -1 {
		resp.Code = tool.RespCodeError
	}
	ctx.Json(resp)
}
