package smsApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/x/util"
)

type SmsInfoResp struct {
	Id         int    `json:"id,omitempty"`
	MessageId  int    `json:"messageId,omitempty"` // message
	TemplateId int    `json:"templateId,omitempty"`
	SenderId   int    `json:"senderId,omitempty"`
	Params     string `json:"params,omitempty"`
	Status     int    `json:"status,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}

type RespWrapper struct {
	app.Response
	SmsInfoResp
}

func GetByMessage(ctx server.Context) {
	log.Debug("get sms")
	var resp RespWrapper
	id, err := ctx.C.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = tool.RespCodeNotFound
		//resp.Message = tool.RespMsgIdNum
		ctx.Json(resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetSmsByMessageId(id)
	if err != nil {
		log.Debug("get sms failed")
		log.Debug(err)
		resp.Response.Code = tool.RespCodeNotFound
		//resp.Message = tool.RespMsgNotFount
		ctx.Json(resp)
		return
	}

	var tmpInfo = SmsInfoResp{
		Id:         ms.Id,
		MessageId:  ms.MessageId,
		TemplateId: ms.Template,
		Params:     ms.Params,
		Status:     ms.Status,
		CreateTime: util.FormatTime(ms.CreateTime),
		UpdateTime: util.FormatTime(ms.UpdateTime),
	}
	resp.SmsInfoResp = tmpInfo
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
