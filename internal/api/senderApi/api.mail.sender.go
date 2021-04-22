package senderApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/repo"
)

type Info struct {
	Id         int    `json:"id,omitempty"`
	SenderCode string `json:"senderCode,omitempty"`
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
	Email      string `json:"email,omitempty"`
	Alias      string `json:"alias,omitempty"`
	Passwd     string `json:"password,omitempty"`
	Status     int    `json:"status,omitempty,omitempty"`
	CreateTime string `json:"createTime,omitempty,omitempty"`
	UpdateTime string `json:"updateTime,omitempty,omitempty"`
}

type InfoWrapper struct {
	app.Response
	Info
}

func GetMailSenderList(ctx iris.Context) {
	var resp app.PagerResponse
	var status = ctx.URLParamIntDefault("status", repo.ConditionIgnore)
	var pageSize = ctx.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	var pageNo = ctx.URLParamIntDefault("pageNo", repo.DefaultPageNo)
	page := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNo,
	}
	page, senders, err := repo.GetMailSenderList(status, page)
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
			var tmpInfo = Info{
				Id:         ms.Id,
				SenderCode: ms.Code,
				Host:       ms.Host,
				Port:       ms.Port,
				Email:      ms.Email,
				Alias:      ms.Alias,
				Passwd:     ms.Passwd,
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

// GetMailSenderInfo
/**
@Router /api/mail_sender/{id} [get]
*/
func GetMailSenderInfo(ctx iris.Context) {
	log.Debug("get mail sender")
	var resp InfoWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgIdNum
		common.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetMailSenderById(id)
	if err != nil {
		log.Debug("get mail sender account failed")
		log.Debug(err)
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgNotFount
		common.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = Info{
		Id:         ms.Id,
		SenderCode: ms.Code,
		Host:       ms.Host,
		Port:       ms.Port,
		Email:      ms.Email,
		Alias:      ms.Alias,
		Passwd:     ms.Passwd,
		Status:     ms.Status,
		CreateTime: common.FormatTime(ms.CreateTime),
		UpdateTime: common.FormatTime(ms.UpdateTime),
	}
	resp.Info = tmpInfo
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func AddMailSender(_ iris.Context) {
	code := common.GetRandomString(common.DefaultCodeLen)
	code = "sender_" + code
}

func UpdateMailSender(ctx iris.Context) {
	log.Debug("update mail sender")
	var resp app.Response
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func DeleteMailSender(ctx iris.Context) {
	log.Debug("delete mail sender")
	var resp app.Response
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgIdNum
		common.ResponseJSON(ctx, resp)
		return
	}
	err = repo.DeleteMailSender(id)
	if err != nil {
		log.Info("delete mail sender failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "delete sender failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
