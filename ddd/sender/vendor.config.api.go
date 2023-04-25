package sender

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
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

func vendorConfig(ctx iris.Context) {

	var err error
	var resp VendorConfigResp
	var req VendorConfigReq
	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info("read req failed")
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
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

var vendorSupport = map[string]interface{}{
	model.SmsVendorAli.String() + model.SenderCategorySms:     model.AliSmsConfig{},
	model.SmsVendorTencent.String() + model.SenderCategorySms: model.TencentSmsConfig{},
	model.SmsVendorHuawei.String() + model.SenderCategorySms:  model.HuaweiSmsConfig{},

	model.MailVendorSmtp.String() + model.SenderCategoryMail:      model.SmtpConfig{},
	model.MailVendorMicrosoft.String() + model.SenderCategoryMail: model.GraphConfig{},
	model.MailVendorTencent.String() + model.SenderCategoryMail:   model.TencentConfig{},
}
