package sender

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type VendorConfigResp struct {
	app.Response
	Config map[string]string `json:"config"`
}

type VendorConfigReq struct {
	Vendor string `json:"vendor,omitempty"`
	Method string `json:"method,omitempty"`
}

// getConfigStruct 返回 sender 配置字段说明，数据来自 pkg/msg 目录。
func getConfigStruct(ctx server.Context) {
	var resp VendorConfigResp
	var req VendorConfigReq
	req.Vendor = ctx.C.Params().Get("vendor")
	req.Method = ctx.C.Params().Get("category")
	fields := msg.ProviderConfigFields(
		msg.MessageCategory(req.Method),
		msg.NormalizeProvider(msg.MessageProvider(req.Vendor)),
	)
	if len(fields) > 0 {
		resp.Config = fields
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
