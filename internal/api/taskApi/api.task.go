package taskApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
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
	page, tasks, err := repo.GetTaskList(status, page)
	if err != nil {
		log.Debug("get templates failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "get templates failed"
		tool.ResponseJSON(ctx, resp)
		return
	}
	if len(tasks) > 0 {
		for _, ms := range tasks {
			var tmpInfo = TaskInfoResp{
				Id:                ms.Id,
				MessageId:         ms.MessageId,
				MessageInstanceId: ms.MessageInstanceId,
				Status:            ms.Status,
				CreateTime:        tool.FormatTime(ms.CreateTime),
				UpdateTime:        tool.FormatTime(ms.UpdateTime),
			}
			page.Data = append(page.Data, tmpInfo)
		}
	}

	resp.Code = tool.RespCodeSuccess
	resp.Pager = page
	tool.ResponseJSON(ctx, resp)
}

func GetTaskInfo(ctx iris.Context) {

	log.Debug("get task")
	var resp RespWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgIdNum
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetMessageTask(id)
	if err != nil {
		log.Debug("get task failed")
		log.Debug(err)
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgNotFount
		tool.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = TaskInfoResp{
		Id:                ms.Id,
		MessageId:         ms.MessageId,
		MessageInstanceId: ms.MessageInstanceId,
		Status:            ms.Status,
		CreateTime:        tool.FormatTime(ms.CreateTime),
		UpdateTime:        tool.FormatTime(ms.UpdateTime),
	}
	resp.TaskInfoResp = tmpInfo
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func GetByMessage(ctx iris.Context) {
	log.Debug("get task by message")
	var resp RespWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgIdNum
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetTaskByMessage(id)
	if err != nil {
		log.Debug("get task failed")
		log.Debug(err)
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgNotFount
		tool.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = TaskInfoResp{
		Id:                ms.Id,
		MessageId:         ms.MessageId,
		MessageInstanceId: ms.MessageInstanceId,
		Status:            ms.Status,
		CreateTime:        tool.FormatTime(ms.CreateTime),
		UpdateTime:        tool.FormatTime(ms.UpdateTime),
	}
	resp.TaskInfoResp = tmpInfo
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
