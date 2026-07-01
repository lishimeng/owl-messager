package sender

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/sdk"
)

type messagerProxy struct {
	Host   string `json:"host"`
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
}

func (p *messagerProxy) normalize() {
	p.Host = resolveMessagerHost(p.Host)
}

func (p messagerProxy) validate() error {
	if len(p.AppId) == 0 || len(p.Secret) == 0 {
		return errors.New("appId and secret required")
	}
	return nil
}

func parseTemplateParams(raw any) (any, error) {
	if raw == nil {
		return nil, nil
	}
	switch v := raw.(type) {
	case string:
		s := strings.TrimSpace(v)
		if s == "" {
			return nil, nil
		}
		var m map[string]any
		if err := json.Unmarshal([]byte(s), &m); err != nil {
			return nil, fmt.Errorf("invalid params json: %w", err)
		}
		return m, nil
	default:
		return raw, nil
	}
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

type testConfigResp struct {
	app.Response
	DefaultHost string `json:"defaultHost,omitempty"`
	SendPath    string `json:"sendPath,omitempty"`
}

func getTestConfig(ctx server.Context) {
	var resp testConfigResp
	resp.Code = tool.RespCodeSuccess
	resp.DefaultHost = defaultMessagerHost()
	resp.SendPath = "POST {host}/messages/{mail|sms|im}"
	resp.Message = "留空 host 则本地直发（不经过 HTTP）；远程测试填 owl-messager 根地址"
	ctx.Json(resp)
}

func testMailSender(ctx server.Context) {
	var req testMailReq

	if err := ctx.C.ReadJSON(&req); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeNotFound
		resp.Message = "json参数解析失败"
		ctx.Json(resp)
		return
	}
	req.normalize()
	if err := req.validate(); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	params, err := parseTemplateParams(req.TemplateParam)
	if err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	req.TemplateParam = params

	if shouldUseInternalSend(req.Host) {
		resp, sendErr := internalSendMail(req)
		writeTestResp(ctx, resp, sendErr)
		return
	}
	if err := probeMessager(req.Host); err != nil {
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
		resp.Message = "json参数解析失败"
		ctx.Json(resp)
		return
	}
	req.normalize()
	if err := req.validate(); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	params, err := parseTemplateParams(req.TemplateParam)
	if err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	req.TemplateParam = params

	if shouldUseInternalSend(req.Host) {
		resp, sendErr := internalSendIm(req)
		writeTestResp(ctx, resp, sendErr)
		return
	}
	if err := probeMessager(req.Host); err != nil {
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
		resp.Message = "json参数解析失败"
		ctx.Json(resp)
		return
	}
	req.normalize()
	if err := req.validate(); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	params, err := parseTemplateParams(req.TemplateParam)
	if err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	if params == nil {
		params = map[string]any{"code": "123456"}
	}
	req.TemplateParam = params

	if shouldUseInternalSend(req.Host) {
		resp, sendErr := internalSendSms(req)
		writeTestResp(ctx, resp, sendErr)
		return
	}
	if err := probeMessager(req.Host); err != nil {
		var resp sdk.Response
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp, err := req.client().SendSms(req.SmsRequest)
	writeTestResp(ctx, resp, err)
}

func writeTestResp(ctx server.Context, resp sdk.Response, err error) {
	if err != nil {
		resp.Code = tool.RespCodeError
		if resp.Message == "" {
			switch err.Error() {
			case "404":
				resp.Message = "Messager 返回 404：该地址没有 POST /messages/{category}。请确认端口上是 owl-messager（可用 GET /messages/ping 验证），不是 console"
			case "401":
				resp.Message = "Messager 认证失败：请检查 AppId 与 Secret"
			default:
				resp.Message = err.Error()
			}
		}
	}
	if resp.Code == -1 || resp.Code == 500 {
		resp.Code = tool.RespCodeError
	}
	if resp.Code == 401 {
		resp.Code = tool.RespCodeError
	}
	if resp.Code != tool.RespCodeSuccess {
		resp.Code = tool.RespCodeError
	}
	ctx.Json(resp)
}
