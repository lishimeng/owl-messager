package templateApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func DeleteTemplate(ctx server.Context) {
	var req TemplateReq
	var respJson app.Response

	err := ctx.C.ReadJSON(&req)
	if err != nil {
		respJson.Code = tool.RespCodeNotFound
		respJson.Message = "json参数解析失败"
		ctx.Json(respJson)
		return
	}

	code := req.Code

	if code == "" {
		respJson.Message = "code为空"
		respJson.Code = tool.RespCodeNotFound
		ctx.Json(respJson)
		return
	}

	var tpl model.MessageTemplate
	err = app.GetOrm().Model(&model.MessageTemplate{}).Equal("code", code).First(&tpl)
	if err != nil {
		log.Info("delTemplate err", err)
		respJson.Message = "删除失败"
		respJson.Code = tool.RespCodeNotFound
		ctx.Json(respJson)
		return
	}
	err = app.Transaction(func(tx persistence.TxContext) error {
		return tx.Delete(&tpl)
	})
	if err != nil {
		log.Info("delTemplate err", err)
		respJson.Message = "删除失败"
		respJson.Code = tool.RespCodeNotFound
		ctx.Json(respJson)
		return
	}
	respJson.Code = tool.RespCodeSuccess
	respJson.Message = "成功！"
	ctx.Json(respJson)
}
