package senderApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/common"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

type SmsSenderInfo struct {
	Id         int    `json:"id,omitempty"`
	SenderCode string `json:"senderCode,omitempty"`
	Vendor     string `json:"vendor,omitempty"`
	Status     int    `json:"status,omitempty,omitempty"`
	CreateTime string `json:"createTime,omitempty,omitempty"`
	UpdateTime string `json:"updateTime,omitempty,omitempty"`
}

type SmsInfoWrapper struct {
	app.Response
	SmsSenderInfo
}

func GetSmsSenderList(ctx iris.Context) {
	var resp app.PagerResponse
	var status = ctx.URLParamIntDefault("status", repo.ConditionIgnore)
	var pageSize = ctx.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	var pageNo = ctx.URLParamIntDefault("pageNo", repo.DefaultPageNo)
	page := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNo,
	}
	page, senders, err := repo.GetSmsSenderList(status, page)
	if err != nil {
		log.Debug("get senders failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "get senders failed"
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(senders) > 0 {
		for _, ms := range senders {
			var tmpInfo = SmsSenderInfo{
				Id:         ms.Id,
				SenderCode: ms.Code,
				Vendor:     string(ms.Vendor),
				Status:     ms.Status,
				CreateTime: tool.FormatTime(ms.CreateTime),
				UpdateTime: tool.FormatTime(ms.UpdateTime),
			}

			page.Data = append(page.Data, tmpInfo)
		}
	}

	resp.Pager = page
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

// GetSmsSenderInfo
/**
@Router /api/sms_sender/{id} [get]
*/
func GetSmsSenderInfo(ctx iris.Context) {
	log.Debug("get sms sender")
	var resp SmsInfoWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgIdNum
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetSmsSenderById(id)
	if err != nil {
		log.Debug("get sms sender account failed")
		log.Debug(err)
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgNotFount
		tool.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = SmsSenderInfo{
		Id:         ms.Id,
		SenderCode: ms.Code,
		Vendor:     string(ms.Vendor),
		Status:     ms.Status,
		CreateTime: tool.FormatTime(ms.CreateTime),
		UpdateTime: tool.FormatTime(ms.UpdateTime),
	}
	resp.SmsSenderInfo = tmpInfo
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func AddSmsSender(_ iris.Context) {
	code := tool.GetRandomString(common.DefaultCodeLen)
	code = "sms_sender_" + code

	// TODO
}

func UpdateSmsSender(ctx iris.Context) {
	log.Debug("update sms sender")
	var resp app.Response

	// TODO
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func DeleteSmsSender(ctx iris.Context) {
	log.Debug("delete sms sender")
	var resp app.Response
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgIdNum
		tool.ResponseJSON(ctx, resp)
		return
	}
	err = repo.DeleteSmsSender(id)
	if err != nil {
		log.Info("delete sms sender failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "delete sms sender failed"
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
