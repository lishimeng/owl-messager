package sender

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"reflect"
	"strings"
)

type VendorConfigResp struct {
	app.Response
	Config map[string]string `json:"config"`
}

type VendorConfigReq struct {
	Vendor string `json:"vendor,omitempty"` // 平台 枚举
	Method string `json:"method,omitempty"` // 功能 枚举 model.SenderCategory
}

// getConfigStruct
// 显示sender的配置字段,以map格式展示
// VendorConfigResp
func getConfigStruct(ctx iris.Context) {

	var resp VendorConfigResp
	var req VendorConfigReq
	req.Vendor = ctx.Params().Get("vendor")
	req.Method = ctx.Params().Get("category")
	config, ok := vendorSupport[req.Vendor+req.Method]
	if ok {
		m := getJsonConstructor(config)
		resp.Config = m
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func getJsonConstructor(v interface{}) (m map[string]string) {
	m = make(map[string]string)
	var a = reflect.TypeOf(v)
	for i := 0; i < a.NumField(); i++ {
		var f = a.Field(i)
		var tag = f.Tag
		jsonContext, ok := tag.Lookup("json")
		if ok {
			var name = strings.Split(jsonContext, ",")[0]
			m[name] = f.Type.Name()
		} else {
			continue
		}
	}
	return
}

// 配置平台支持的sender类型
var vendorSupport = map[string]interface{}{
	model.SmsVendorAli.String() + model.SenderCategorySms:     model.AliSmsConfig{},
	model.SmsVendorTencent.String() + model.SenderCategorySms: model.TencentSmsConfig{},
	model.SmsVendorHuawei.String() + model.SenderCategorySms:  model.HuaweiSmsConfig{},

	model.MailVendorSmtp.String() + model.SenderCategoryMail:      model.SmtpConfig{},
	model.MailVendorMicrosoft.String() + model.SenderCategoryMail: model.GraphConfig{},
	model.MailVendorTencent.String() + model.SenderCategoryMail:   model.TencentConfig{},
}
