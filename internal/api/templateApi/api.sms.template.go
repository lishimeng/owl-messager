package templateApi

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/api/common"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/util"
)

type SmsVendors struct {
	app.PagerResponse
}

// GetSmsVendors 平台支持的sms类型
func GetSmsVendors(ctx iris.Context) {
	var resp SmsVendors

	resp.Code = tool.RespCodeSuccess
	resp.Message = "Sms Vendors"

	for key, v := range model.SmsVendors {
		if v == model.SmsVendorEnable {
			resp.Data = append(resp.Data, key)
		}
	}
	tool.ResponseJSON(ctx, resp)
}

// SmsTemplateReq
// 创建请求
type SmsTemplateReq struct {
	Id          int    `json:"id,omitempty"`          // update时可用
	Code        string `json:"code,omitempty"`        // update时可用
	Name        string `json:"name,omitempty"`        // 名称
	Params      string `json:"params,omitempty"`      // 参数列表,可选,无用
	Description string `json:"description,omitempty"` // 描述
	Vendor      string `json:"vendor,omitempty"`      // 平台
	TemplateId  string `json:"templateId,omitempty"`  // 平台里的模板编号
}

type SmsTemplateResp struct {
}

func AddSmsTemplate(ctx iris.Context) {

	log.Debug("add sms template")
	var req SmsTemplateReq
	var resp InfoWrapper
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = -1
		tool.ResponseJSON(ctx, resp)
		return
	}

	// check params
	if len(req.Name) == 0 {
		log.Debug("param name nil")
		resp.Code = -1
		resp.Message = "name nil"
		tool.ResponseJSON(ctx, resp)
		return
	}
	if len(req.TemplateId) == 0 {
		log.Debug("param template nil")
		resp.Code = -1
		resp.Message = "template nil"
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.Vendor) == 0 {
		log.Debug("param vendor nil")
		resp.Code = -1
		resp.Message = "vendor nil"
		tool.ResponseJSON(ctx, resp)
		return
	}

	code := tool.GetRandomString(common.DefaultCodeLen)
	code = "tpl_" + code

	m, err := repo.CreateSmsTemplate(code, req.Name, req.TemplateId, req.Params, req.Description, req.Vendor)
	if err != nil {
		log.Info("can't create template")
		log.Info(err)
		resp.Code = -1
		resp.Message = "create template failed"
		tool.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("create template success, id:%d", m.Id)
	resp.Id = m.Id

	var tmpInfo = Info{
		Id:           m.Id,
		TemplateCode: m.Code,
		TemplateBody: m.Body,
		Status:       m.Status,
		CreateTime:   tool.FormatTime(m.CreateTime),
		UpdateTime:   tool.FormatTime(m.UpdateTime),
	}
	resp.Info = tmpInfo
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

type SmsStatusReq struct {
	Status int `json:"status,omitempty"`
	Id     int `json:"id,omitempty"`
}

func ChangeSmsTemplateStatus(ctx iris.Context) {

	var req SmsStatusReq
	var resp app.Response
	var err error

	err = ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	if req.Id <= 0 {
		log.Debug("param id nil")
		resp.Code = tool.RespCodeError
		resp.Message = "id nil"
		tool.ResponseJSON(ctx, resp)
		return
	}

	if !util.StatusIn(req.Status, model.SmsTemplateStatus) {
		log.Debug("param unknown status: %d", req.Status)
		resp.Code = tool.RespCodeError
		resp.Message = fmt.Sprintf("unknown status:%d", req.Status)
		tool.ResponseJSON(ctx, resp)
		return
	}

	tpl, err := repo.GetSmsTemplateById(req.Id)
	if err != nil {
		log.Debug("template not found")
		resp.Code = tool.RespCodeNotFound
		resp.Message = "template not found"
		tool.ResponseJSON(ctx, resp)
		return
	}

	tpl.Status = req.Status

	_, err = repo.UpdateSmsTemplate(tpl, "status")
	if err != nil {
		log.Debug(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}