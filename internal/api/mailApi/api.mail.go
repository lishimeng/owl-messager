package mailApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/repo"
)

type MailInfoResp struct {
	Id         int    `json:"id,omitempty"`
	MessageId  int    `json:"messageId,omitempty"` // message
	TemplateId int    `json:"templateId,omitempty"`
	Params     string `json:"params,omitempty"`
	Status     int    `json:"status,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}

type RespWrapper struct {
	app.Response
	MailInfoResp
}

func GetByMessage(ctx iris.Context) {
	log.Debug("get mail")
	var resp RespWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgIdNum
		common.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetMailByMessageId(id)
	if err != nil {
		log.Debug("get mail failed")
		log.Debug(err)
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgNotFount
		common.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = MailInfoResp{
		Id:         ms.Id,
		MessageId:  ms.MessageId,
		TemplateId: ms.Template,
		Params:     ms.Params,
		Status:     ms.Status,
		CreateTime: common.FormatTime(ms.CreateTime),
		UpdateTime: common.FormatTime(ms.UpdateTime),
	}
	resp.MailInfoResp = tmpInfo
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
