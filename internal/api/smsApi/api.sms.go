package smsApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/repo"
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

func GetByMessage(ctx iris.Context) {
	log.Debug("get sms")
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
	ms, err := repo.GetSmsByMessageId(id)
	if err != nil {
		log.Debug("get sms failed")
		log.Debug(err)
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgNotFount
		common.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = SmsInfoResp{
		Id:         ms.Id,
		MessageId:  ms.MessageId,
		TemplateId: ms.Template,
		SenderId:   ms.Sender,
		Params:     ms.Params,
		Status:     ms.Status,
		CreateTime: common.FormatTime(ms.CreateTime),
		UpdateTime: common.FormatTime(ms.UpdateTime),
	}
	resp.SmsInfoResp = tmpInfo
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
