package dict

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
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

// providerList 返回 pkg/msg 目录中的 provider 列表。
func providerList(ctx server.Context) {
	log.Info("get providers from catalog")
	var resp Providers
	category := msg.MessageCategory(ctx.C.Params().Get("category"))
	for _, e := range msg.ListCatalog(category) {
		resp.Items = append(resp.Items, ProviderInfo{
			Name:     e.Provider.String(),
			Category: string(e.Category),
		})
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
