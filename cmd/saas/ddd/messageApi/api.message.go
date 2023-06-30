package messageApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

type Req struct {
}

type RespMessageInfo struct {
	Id           int    `json:"id,omitempty"`
	Status       int    `json:"status,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	UpdateTime   string `json:"updateTime,omitempty"`
	Category     int    `json:"category,omitempty"`
	Subject      string `json:"subject,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	NextSendTime string `json:"nextSendTime,omitempty"`
}

type RespMessageInfoWrapper struct {
	app.Response
	RespMessageInfo
}

type RespMessageInfoListWrapper struct {
	app.PagerResponse
}

func GetMessageList(ctx iris.Context) {
	var resp app.PagerResponse
	var status = ctx.URLParamIntDefault("status", repo.ConditionIgnore)
	var category = ctx.URLParamIntDefault("category", repo.ConditionIgnore)
	var pageSize = ctx.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	var pageNo = ctx.URLParamIntDefault("pageNo", repo.DefaultPageNo)
	page := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNo,
	}
	page, messages, err := repo.GetMessages(status, category, page)
	if err != nil {
		log.Debug("get messages failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "get messages failed"
		tool.ResponseJSON(ctx, resp)
		return
	}
	if len(messages) > 0 {
		for _, ms := range messages {
			var tmpInfo = RespMessageInfo{
				Id:           ms.Id,
				Category:     ms.Category,
				Subject:      ms.Subject,
				Priority:     ms.Priority,
				NextSendTime: tool.FormatTime(ms.NextSendTime),
				Status:       ms.Status,
				CreateTime:   tool.FormatTime(ms.CreateTime),
				UpdateTime:   tool.FormatTime(ms.UpdateTime),
			}
			page.Data = append(page.Data, tmpInfo)
		}
	}

	resp.Pager = page
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)

}

func GetMessageInfo(ctx iris.Context) {
	log.Debug("get message")
	var resp RespMessageInfoWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgIdNum
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetMessageById(id)
	if err != nil {
		log.Debug("get message failed")
		log.Debug(err)
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgNotFount
		tool.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = RespMessageInfo{
		Id:           ms.Id,
		Category:     ms.Category,
		Subject:      ms.Subject,
		Priority:     ms.Priority,
		NextSendTime: tool.FormatTime(ms.NextSendTime),
		Status:       ms.Status,
		CreateTime:   tool.FormatTime(ms.CreateTime),
		UpdateTime:   tool.FormatTime(ms.UpdateTime),
	}
	resp.RespMessageInfo = tmpInfo
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

// Send
/**
@Summary send a message now, this will change message to a high priority

@Router /api/message/send/{id} [post]
*/
func Send(ctx iris.Context) {
	log.Debug("send message[manual]")
	var resp app.Response
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = tool.RespMsgNotFount
		resp.Message = tool.RespMsgNotFount
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Info("set message high priority:%d", id)
	_, err = repo.UpdateMessagePriority(id, model.MessagePriorityHigh)
	if err != nil {
		log.Info("set message priority:%d failed", id)
		log.Info(err)
		resp.Code = -1
		resp.Message = "failed"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
