package sender

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/consoleorg"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/lishimeng/x/util"
	"time"
)

type reqSender struct {
	DefaultSender int                 `json:"defaultSender,omitempty"`
	TenantCode    string              `json:"tenantCode,omitempty"`
	Vendor        msg.MessageProvider `json:"vendor,omitempty"`
	Config        msg.SenderConfig    `json:"config,omitempty"`
	Code          string              `json:"code,omitempty"`
	Category      msg.MessageCategory `json:"category,omitempty"`
}

func SetMailSenderInfo(ctx server.Context) {
	var resp app.Response
	var req reqSender
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "json参数解析失败"
		ctx.Json(resp)
		return
	}
	code := util.UUIDString()
	switch req.Category {
	case msg.MailMessage:
		code = "sender_mail_" + code
	case msg.SmsMessage:
		code = "sender_sms_" + code
	case msg.ImMessage:
		code = "sender_im_" + code
	default:
		resp.Code = tool.RespCodeNotFound
		resp.Message = "失败：未知通讯方式"
		ctx.Json(resp)
		return
	}

	if req.TenantCode == "" {
		resp.Code = tool.RespCodeError
		resp.Message = "tenantCode required"
		ctx.Json(resp)
		return
	}

	_, err = repo.CreateMessageSender(req.TenantCode, req.Category, req.Vendor, 0, code, req.Config)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "创建失败"
		ctx.Json(resp)
		return
	}
	if req.DefaultSender == 1 {
		err = _setDefault(code, req.Category.String(), req.TenantCode, req.Vendor.String())
		if err != nil {
			resp.Code = tool.RespCodeError
			ctx.Json(resp)
			return
		}
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func UpMailSenderInfo(ctx server.Context) {
	var resp app.Response
	var req reqSender
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	_, err = repo.UpdateMessageSender(req.Code, req.Config)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	if req.DefaultSender == 1 {
		sender, e := repo.GetMessageSenderByCode(req.Code)
		if e != nil {
			resp.Code = tool.RespCodeNotFound
			ctx.Json(resp)
			return
		}
		err = _setDefault(req.Code, req.Category.String(), sender.TenantCode, req.Vendor.String())
		if err != nil {
			resp.Code = tool.RespCodeError
			ctx.Json(resp)
			return
		}
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

type respSender struct {
	Id            int    `json:"id,omitempty"`
	DefaultSender int    `json:"defaultSender,omitempty"`
	Vendor        string `json:"vendor,omitempty"`
	Code          string `json:"code,omitempty"`
	Config        string `json:"config,omitempty"`
	CreateTime    string `json:"createTime,omitempty"`
	UpdateTime    string `json:"updateTime,omitempty"`
}

type respPager struct {
	app.PagerResponse
	app.BasePager
	Items []respSender `json:"items"`
}

func GetMailSenderInfo(ctx server.Context) {
	//var resp RespInfo
	//var vendor = ctx.URLParamDefault("vendor", "")
	//log.Debug("vendor:%s", vendor)
	//m := model.MailSenderInfo{}
	//err := app.GetOrm().Context.QueryTable(new(model.MailSenderInfo)).Filter("Vendor", vendor).One(&m)
	//if err != nil {
	//	resp.Code = tool.RespCodeNotFound
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//resp.Scode = m.Code
	//resp.DefaultSender = m.Default
	//resp.Vendor = string(m.Vendor)
	//resp.Config = string(m.Config)
	//resp.Code = tool.RespCodeSuccess
	//log.Debug("resp:%s", resp)
	//tool.ResponseJSON(ctx, resp) TODO
}

func ListByPage(ctx server.Context) {
	var resp respPager
	var category = ctx.C.URLParamDefault("category", "")
	var pageNum = ctx.C.URLParamIntDefault("pageNum", 1)
	var pageSize = ctx.C.URLParamIntDefault("pageSize", 10)
	var pager app.SimplePager[model.MessageSenderInfo, respSender]
	pager.PageSize = pageSize
	pager.PageNum = pageNum
	pager.Transform = func(src model.MessageSenderInfo, dst *respSender) {
		dst.Id = src.Id
		dst.Vendor = string(src.Provider)
		dst.Code = src.Code
		dst.DefaultSender = src.Default
		dst.CreateTime = src.CreateTime.UTC().Format(time.RFC3339)
		dst.UpdateTime = src.UpdateTime.UTC().Format(time.RFC3339)
	}
	pager.QueryBuilder = func(tx persistence.TxContext) persistence.Query {
		q := tx.Model(&model.MessageSenderInfo{})
		if tenantCode := consoleorg.TenantFilter(ctx); tenantCode != "" {
			q = q.Equal("tenant_code", tenantCode)
		}
		if len(category) > 0 {
			q = q.Equal("message_category", category)
		}
		return q
	}
	pager.OrderExp = append(pager.OrderExp, "message_provider", "ctime")
	err := app.QueryPage(&pager)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "not found"
		ctx.Json(resp)
		return
	}
	resp.Items = pager.Data
	resp.BasePager = pager.BasePager
	resp.BasePager.More = pager.TotalPage * pager.PageSize

	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

type respSenderOne struct {
	app.Response
	Item respSender `json:"item"`
}

func GetSenderInfoByCategory(ctx server.Context) {
	var resp respSenderOne
	var code = ctx.C.URLParamDefault("code", "")
	//var category = ctx.C.URLParamDefault("category", "")
	if code == "" {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "请求中code为空"
		ctx.Json(resp)
		return
	}

	tpl, err := repo.GetMessageSenderByCode(code)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "未查到记录"
		ctx.Json(resp)
		return
	}
	err = tpl.Config.Decode()
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "记录异常"
		ctx.Json(resp)
		return
	}
	resp.Item = respSender{
		Code:          tpl.Code,
		DefaultSender: tpl.Default,
		Vendor:        string(tpl.Provider),
		Config:        string(tpl.Config),
	}
	resp.Code = tool.RespCodeSuccess
	resp.Message = "成功"
	ctx.Json(resp)
}

func DeleteSender(ctx server.Context) {
	var req reqSender
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

	var sender model.MessageSenderInfo
	err = app.GetOrm().Model(&model.MessageSenderInfo{}).Equal("code", code).First(&sender)
	if err != nil {
		respJson.Message = "删除失败"
		respJson.Code = tool.RespCodeNotFound
		ctx.Json(respJson)
		return
	}
	err = app.Transaction(func(tx persistence.TxContext) error {
		return tx.Delete(&sender)
	})
	if err != nil {
		respJson.Message = "删除失败"
		respJson.Code = tool.RespCodeNotFound
		ctx.Json(respJson)
		return
	}
	respJson.Code = tool.RespCodeSuccess
	respJson.Message = "成功！"
	ctx.Json(respJson)
}
