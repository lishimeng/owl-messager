package senderApi

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"time"
)

type Info struct {
	Id         int    `json:"id,omitempty"`
	SenderCode string `json:"senderCode,omitempty"`
	Config     string `json:"config,omitempty"`
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
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(senders) > 0 {
		for _, ms := range senders {
			var tmpInfo = Info{
				Id:         ms.Id,
				SenderCode: ms.Code,
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
func GetMailSenderInfo(ctx iris.Context) {
	log.Debug("get mail sender")
	var resp InfoWrapper
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

	var tmpInfo = Info{
		Id:         ms.Id,
		SenderCode: ms.Code,
		Config:     ms.Config,
		Status:     ms.Status,
		CreateTime: tool.FormatTime(ms.CreateTime),
		UpdateTime: tool.FormatTime(ms.UpdateTime),
	}
	resp.Info = tmpInfo
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func AddMailSender(ctx iris.Context) {
	//code := tool.GetRandomString(common.DefaultCodeLen)
	//code = "sender_" + code
	var reqJson model.MailSenderInfo
	var resp app.Response

	resp.Code = 200
	resp.Message = "成功！"

	errJson := ctx.ReadJSON(&reqJson)

	if errJson != nil {
		log.Info("addItems [ctx.ReadJSON] err", errJson)
		resp.Message = "items添加失败！"
		resp.Code = 500
		tool.ResponseJSON(ctx, resp) //转为json格式
		return
	}
	//reqJson.Id=
	_, err2 := app.GetOrm().Context.Insert(&reqJson)

	if err2 != nil {
		log.Info("addItems err", err2)
		resp.Message = "添加失败"
		resp.Code = 500
		tool.ResponseJSON(ctx, resp)
		return
	}
	tool.ResponseJSON(ctx, resp)
	return
}

func UpdateMailSender(ctx iris.Context) {
	log.Debug("update mail sender")
	var resp app.Response
	var respJson InfoNow

	resp.Code = tool.RespCodeSuccess
	errJson :=ctx.ReadJSON(&respJson)

	if errJson != nil {
		log.Info("updateItemsInfo [ctx.ReadJSON] err", errJson)
		resp.Message = "sku解析失败！"
		resp.Code = 500
		tool.ResponseJSON(ctx, resp) //转为json格式
		return
	}
	var up model.MailSenderInfo
	up.Id=respJson.Id
	up.Status=respJson.Status
	up.Code=respJson.Code
	up.Password=respJson.Password
	up.Email=respJson.Email
	up.Host=respJson.Host
	up.Port=respJson.Port
	up.UpdateTime=time.Now()

	_, err2 := orm.NewOrm().Update(&up)
	if err2 != nil {
		log.Info("updateItemsInfo", errJson)
		resp.Message = "sku修改失败！"
		resp.Code = 500
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Message = "修改成功!"
	resp.Code = 200
	tool.ResponseJSON(ctx, resp)
	return
}

func DeleteMailSender(ctx iris.Context) {
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
