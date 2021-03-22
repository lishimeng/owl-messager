package templateApi

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/repo"
)

type Info struct {
	Id int `json:"id,omitempty"`
	TemplateCode string `json:"templateCode,omitempty"`
	TemplateBody string `json:"templateBody,omitempty"`
	Status int `json:"status,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}

type InfoWrapper struct {
	app.Response
	Info
}

func GetMailTemplateList(ctx iris.Context) {

}

func GetMailTemplateInfo(ctx iris.Context) {
	var resp InfoWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = -1
		resp.Message = "id must be a int value"
		common.ResponseJSON(ctx, resp)
		return
	}
	tpl, err := repo.GetMailTemplateById(id)
	if err != nil {
		log.Info("delete mail template failed")
		resp.Response.Code = -1
		resp.Message = "delete template failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = Info{
		Id:           tpl.Id,
		TemplateCode: tpl.Code,
		TemplateBody: tpl.Body,
		Status:       tpl.Status,
		CreateTime:   common.FormatTime(tpl.CreateTime),
		UpdateTime:   common.FormatTime(tpl.UpdateTime),
	}
	resp.Info = tmpInfo

	common.ResponseJSON(ctx, resp)
}

type MailTemplateReq struct {
	Id int `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Body string `json:"body"`
}

func AddMailTemplate(ctx iris.Context) {
	log.Debug("send mail api")
	var req MailTemplateReq
	var resp InfoWrapper
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = -1
		common.ResponseJSON(ctx, resp)
		return
	}

	// check params
	if len(req.Body) == 0 {
		log.Debug("param body nil")
		resp.Code = -1
		resp.Message = "body nil"
		common.ResponseJSON(ctx, resp)
		return
	}

	code := common.GetRandomString(common.DefaultCodeLen)

	m, err := repo.CreateMailTemplate(code, req.Body)
	if err != nil {
		log.Info("can't create template")
		log.Info(err)
		resp.Code = -1
		resp.Message = "create template failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("create template success, id:%d", m.Id)
	resp.Id = m.Id

	var tmpInfo = Info{
		Id:           m.Id,
		TemplateCode: m.Code,
		TemplateBody: m.Body,
		Status:       m.Status,
		CreateTime:   common.FormatTime(m.CreateTime),
		UpdateTime:   common.FormatTime(m.UpdateTime),
	}
	resp.Info = tmpInfo
	common.ResponseJSON(ctx, resp)
}

func UpdateMailTemplate(ctx iris.Context) {

}

func DeleteMailTemplate(ctx iris.Context) {

	var resp app.Response
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = -1
		resp.Message = "id must be a int value"
		common.ResponseJSON(ctx, resp)
		return
	}
	err = repo.DeleteMailTemplate(id)
	if err != nil {
		log.Info("delete mail template failed")
		resp.Code = -1
		resp.Message = "delete template failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	common.ResponseJSON(ctx, resp)
}