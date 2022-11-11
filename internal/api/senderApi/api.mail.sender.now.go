package senderApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/repo"
)

type InfoNow struct {
	Id         int    `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Config     string `json:"config,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int `json:"port,omitempty"`
	Email     string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	Status     int    `json:"status,omitempty,omitempty"`
	CreateTime string `json:"createTime,omitempty,omitempty"`
	UpdateTime string `json:"updateTime,omitempty,omitempty"`
}

type InfoWrapperNow struct {
	app.Response
	Data interface{} `json:"data"`
}

func GetMailSenderListNow(ctx iris.Context) {
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
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(senders) > 0 {
		for _, ms := range senders {
			var tmpInfo = InfoNow{
				Id:         ms.Id,
				Code:       ms.Code,
                Host:       ms.Host,
                Port:       ms.Port,
                Email:      ms.Email,
                Password:   ms.Password,
				Config:     ms.Config,
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

// GetMailSenderInfo
/**
@Router /api/mail_sender/{id} [get]
*/
func GetMailSenderInfoNow(ctx iris.Context) {
	log.Debug("get mail sender")
	var resp InfoWrapperNow
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgIdNum
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetMailSenderById(id)
	if err != nil {
		log.Debug("get mail sender account failed")
		log.Debug(err)
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgNotFount
		tool.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = InfoNow{
		Id:         ms.Id,
		Code:       ms.Code,
		Host:       ms.Host,
		Port:       ms.Port,
		Email:      ms.Email,
		Password:   ms.Password,
		Config:     ms.Config,
		Status:     ms.Status,
		CreateTime: tool.FormatTime(ms.CreateTime),
		UpdateTime: tool.FormatTime(ms.UpdateTime),
	}
	resp.Data = tmpInfo
	resp.Response.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
	return
}

func AddMailSenderNow(_ iris.Context) {
	code := tool.GetRandomString(common.DefaultCodeLen)
	code = "sender_" + code
}

func UpdateMailSenderNow(ctx iris.Context) {
	log.Debug("update mail sender")
	var resp app.Response
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func DeleteMailSenderNow(ctx iris.Context) {
	log.Debug("delete mail sender")
	var resp app.Response
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgIdNum
		tool.ResponseJSON(ctx, resp)
		return
	}
	err = repo.DeleteMailSender(id)
	if err != nil {
		log.Info("delete mail sender failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "delete sender failed"
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
