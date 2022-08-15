package sender

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/api/common"
)

type Item struct {
}

func list(ctx iris.Context) {

	var resp app.PagerResponse

	var items []Item
	// TODO

	for _, item := range items {
		resp.Data = append(resp.Data, item)
	}
	// TODO pager
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func mailSenderInfo(ctx iris.Context) {

}

func smsSenderInfo(ctx iris.Context) {

}

func apnsSenderInfo(ctx iris.Context) {

}
