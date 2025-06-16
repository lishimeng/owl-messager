package templateApi

import (
	"github.com/lishimeng/app-starter"
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

	_, err = app.GetOrm().Context.QueryTable(new(model.MessageTemplate)).
		Filter("Code", code).
		Delete()
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
