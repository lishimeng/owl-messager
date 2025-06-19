package sender

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
)

type Item struct {
}

func list(ctx server.Context) {

	var resp app.ResponseWrapper

	var items []Item
	// TODO

	resp.Data = items
	// TODO pager
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func mailSenderInfo(ctx server.Context) {

}

func smsSenderInfo(ctx server.Context) {

}

func apnsSenderInfo(ctx server.Context) {

}
