package historyApi

import (
	"encoding/json"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/provider/template"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type historyDetail struct {
	TemplateName string `json:"template_name"`
	Receivers    string `json:"receivers"`
	Category     string `json:"category"`
	Provider     string `json:"provider"`
	Params       string `json:"params"`
	Content      string `json:"content"`
}

type respHistoryOne struct {
	app.Response
	Item historyDetail `json:"item"`
}

func GetHistoryOne(ctx server.Context) {
	var resp respHistoryOne
	var templateId int
	id, err := ctx.C.URLParamInt("message_id")
	category := ctx.C.URLParamDefault("category", "")
	resp.Item.Category = category
	if err != nil {
		resp.Message = err.Error()
		resp.Status = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	switch msg.MessageCategory(category) {
	case msg.MailMessage:
		var info model.MailMessageInfo
		err = app.GetOrm().Context.QueryTable(new(model.MailMessageInfo)).
			Filter("message_id", id).
			One(&info)
		if err != nil {
			resp.Message = "not found"
			resp.Status = tool.RespCodeNotFound
			ctx.Json(resp)
			return
		}
		resp.Item.Params = info.Params
		resp.Item.Receivers = info.Receivers
		templateId = info.Template
	case msg.SmsMessage:
		var info model.SmsMessageInfo
		err = app.GetOrm().Context.QueryTable(new(model.SmsMessageInfo)).
			Filter("message_id", id).
			One(&info)
		if err != nil {
			resp.Message = "not found"
			resp.Status = tool.RespCodeNotFound
			ctx.Json(resp)
			return
		}
		resp.Item.Params = info.Params
		resp.Item.Receivers = info.Receivers
		templateId = info.Template
	case msg.ImMessage:
		var info model.ImMessageInfo
		err = app.GetOrm().Context.QueryTable(new(model.ImMessageInfo)).
			Filter("message_id", id).
			One(&info)
		if err != nil {
			resp.Message = "not found"
			resp.Status = tool.RespCodeNotFound
			ctx.Json(resp)
			return
		}
		resp.Item.Params = info.Params
		resp.Item.Receivers = info.Receivers
		templateId = info.Template
	default:
		resp.Message = "unknown category"
		resp.Status = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	// 手动二次查询template，不使用RelatedSel或JOIN
	var tpl model.MessageTemplate
	err = app.GetOrm().Context.QueryTable(new(model.MessageTemplate)).
		Filter("id", templateId).
		One(&tpl)
	if err != nil {
		resp.Message = "template not found"
		resp.Status = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	resp.Item.Provider = tpl.Provider.String()
	resp.Item.TemplateName = tpl.Name
	body := tpl.Body
	if body != "" {
		var params map[string]interface{}
		err = json.Unmarshal([]byte(resp.Item.Params), &params)
		if err != nil {
			resp.Message = err.Error()
			resp.Status = tool.RespCodeError
			ctx.Json(resp)
			return
		}
		resp.Item.Content, err = template.Rend(params, body)
		if err != nil {
			resp.Message = err.Error()
			resp.Status = tool.RespCodeNotFound
		}
	}
	ctx.Json(resp)
}
