package dict

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/provider"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type Providers struct {
	app.Response
	Items []ProviderInfo `json:"items,omitempty"`
}

type ProviderInfo struct {
	Name     string `json:"name,omitempty"`
	Category string `json:"category,omitempty"`
}

// 支持的provider列表
func providerList(ctx server.Context) {
	log.Info("get all providers...")
	var resp Providers

	mailList := provider.GetMailProviders()
	for _, mail := range mailList {
		resp.Items = append(resp.Items, ProviderInfo{
			Name:     mail.String(),
			Category: string(msg.MailMessage),
		})
	}

	smsList := provider.GetSmsProviders()
	for _, sms := range smsList {
		resp.Items = append(resp.Items, ProviderInfo{
			Name:     sms.String(),
			Category: string(msg.SmsMessage),
		})
	}

	imList := provider.GetImProviders()
	for _, im := range imList {
		resp.Items = append(resp.Items, ProviderInfo{
			Name:     im.String(),
			Category: string(msg.ImMessage),
		})
	}

	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func providerCategoryList(ctx server.Context) {
	category := ctx.C.Params().Get("category")
	log.Info(category)
}
