package sender

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/service"
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
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func mailSenderInfo(ctx iris.Context) {

}

func smsSenderInfo(ctx iris.Context) {

}

func apnsSenderInfo(ctx iris.Context) {

}

func setDefaultSender(ctx iris.Context) {

	var err error
	var resp app.Response

	var org = ctx.Params().GetIntDefault("org", -1)
	var category = ctx.Params().Get("category")
	var id = ctx.Params().GetIntDefault("id", -1)
	if org <= 0 {
		log.Info("org: %d", org)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	if id <= 0 {
		log.Info("id: %d", id)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	switch category {
	case model.SenderCategoryMail:
		err = service.SetDefaultMailSender(id, org)
	case model.SenderCategorySms:

	default:
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	if err != nil {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
		// TODO
	}

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
	return
}
