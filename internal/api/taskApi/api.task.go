package taskApi

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/repo"
)

type TaskInfoResp struct {
	Id                int    `json:"id,omitempty"`
	MessageId         int    `json:"messageId,omitempty"`         // message
	MessageInstanceId int    `json:"messageInstanceId,omitempty"` // sms id/mail id
	Status            int    `json:"status,omitempty,omitempty"`
	CreateTime        string `json:"createTime,omitempty,omitempty"`
	UpdateTime        string `json:"updateTime,omitempty,omitempty"`
}

type RespWrapper struct {
	app.Response
	TaskInfoResp
}

func GetTaskList(ctx iris.Context) {
	log.Debug("get task list")
	var resp app.PagerResponse

	var status = ctx.URLParamIntDefault("status", repo.ConditionIgnore)
	var pageSize = ctx.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	var pageNo = ctx.URLParamIntDefault("pageNo", repo.DefaultPageNo)
	page := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNo,
	}
	page, err := repo.GetTaskList(status, page)
	if err != nil {
		log.Debug("get templates failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "get templates failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func GetTaskInfo(ctx iris.Context) {

	log.Debug("get task")
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
	ms, err := repo.GetMessageTask(id)
	if err != nil {
		log.Debug("get task failed")
		log.Debug(err)
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgNotFount
		common.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = TaskInfoResp{
		Id:                ms.Id,
		MessageId:         ms.MessageId,
		MessageInstanceId: ms.MessageInstanceId,
		Status:            ms.Status,
		CreateTime:        common.FormatTime(ms.CreateTime),
		UpdateTime:        common.FormatTime(ms.UpdateTime),
	}
	resp.TaskInfoResp = tmpInfo
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
