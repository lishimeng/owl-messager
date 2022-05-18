package senderApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/repo"
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
		common.ResponseJSON(ctx, resp)
		return
	}

	if len(senders) > 0 {
		for _, ms := range senders {
			var tmpInfo = SmsSenderInfo{
				Id:         ms.Id,
				SenderCode: ms.Code,
				Vendor:     string(ms.Vendor),
				Status:     ms.Status,
				CreateTime: common.FormatTime(ms.CreateTime),
				UpdateTime: common.FormatTime(ms.UpdateTime),
			}

			page.Data = append(page.Data, tmpInfo)
		}
	}

	resp.Pager = page
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
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
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgIdNum
		common.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetSmsSenderById(id)
	if err != nil {
		log.Debug("get sms sender account failed")
		log.Debug(err)
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgNotFount
		common.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = SmsSenderInfo{
		Id:         ms.Id,
		SenderCode: ms.Code,
		Vendor:     string(ms.Vendor),
		Status:     ms.Status,
		CreateTime: common.FormatTime(ms.CreateTime),
		UpdateTime: common.FormatTime(ms.UpdateTime),
	}
	resp.SmsSenderInfo = tmpInfo
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func AddSmsSender(_ iris.Context) {
	code := common.GetRandomString(common.DefaultCodeLen)
	code = "sms_sender_" + code

	// TODO
}

func UpdateSmsSender(ctx iris.Context) {
	log.Debug("update sms sender")
	var resp app.Response

	// TODO
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func DeleteSmsSender(ctx iris.Context) {
	log.Debug("delete sms sender")
	var resp app.Response
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgIdNum
		common.ResponseJSON(ctx, resp)
		return
	}
	err = repo.DeleteSmsSender(id)
	if err != nil {
		log.Info("delete sms sender failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "delete sms sender failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
