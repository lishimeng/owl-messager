package apnsApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/x/util"
)

type ApnsInfoResp struct {
	Id         int    `json:"id,omitempty"`
	Mode       int    `json:"mode"`
	MessageId  int    `json:"messageId,omitempty"` // message
	BundleId   string `json:"bundleId,omitempty"`
	Params     string `json:"params,omitempty"`
	Status     int    `json:"status,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}

type RespWrapper struct {
	app.Response
	ApnsInfoResp
}

func GetByMessage(ctx server.Context) {
	log.Debug("get apns")
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
	ms, err := repo.GetApnsByMessageId(id)
	if err != nil {
		log.Debug("get apns failed")
		log.Debug(err)
		resp.Response.Code = tool.RespCodeNotFound
		//resp.Message = tool.RespMsgNotFount
		ctx.Json(resp)
		return
	}

	var tmpInfo = ApnsInfoResp{
		Id:         ms.Id,
		MessageId:  ms.MessageId,
		BundleId:   ms.BundleId,
		Params:     ms.Params,
		Status:     ms.Status,
		CreateTime: util.FormatTime(ms.CreateTime),
		UpdateTime: util.FormatTime(ms.UpdateTime),
	}
	resp.ApnsInfoResp = tmpInfo
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
