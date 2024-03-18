package messageApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/x/util"
)

type Req struct {
}

type RespMessageInfo struct {
	Id           int    `json:"id,omitempty"`
	Status       int    `json:"status,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	UpdateTime   string `json:"updateTime,omitempty"`
	Category     string `json:"category,omitempty"`
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

func GetMessageList(ctx server.Context) {
	var resp app.PagerResponse
	var status = ctx.C.URLParamIntDefault("status", repo.ConditionIgnore)
	var category = ctx.C.URLParamIntDefault("category", repo.ConditionIgnore)
	var pageSize = ctx.C.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	var pageNo = ctx.C.URLParamIntDefault("pageNo", repo.DefaultPageNo)
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
		ctx.Json(resp)
		return
	}
	if len(messages) > 0 {
		for _, ms := range messages {
			var tmpInfo = RespMessageInfo{
				Id:           ms.Id,
				Category:     ms.Category.String(),
				Subject:      ms.Subject,
				Priority:     ms.Priority,
				NextSendTime: util.FormatTime(ms.NextSendTime),
				Status:       ms.Status,
				CreateTime:   util.FormatTime(ms.CreateTime),
				UpdateTime:   util.FormatTime(ms.UpdateTime),
			}
			page.Data = append(page.Data, tmpInfo)
		}
	}

	resp.Pager = page
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)

}

func GetMessageInfo(ctx server.Context) {
	log.Debug("get message")
	var resp RespMessageInfoWrapper
	id, err := ctx.C.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = tool.RespCodeNotFound
		//resp.Message = tool.RespMsgIdNum
		ctx.Json(resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetMessageById(id)
	if err != nil {
		log.Debug("get message failed")
		log.Debug(err)
		resp.Response.Code = tool.RespCodeNotFound
		//resp.Message = tool.RespMsgNotFount
		ctx.Json(resp)
		return
	}

	var tmpInfo = RespMessageInfo{
		Id:           ms.Id,
		Category:     ms.Category.String(),
		Subject:      ms.Subject,
		Priority:     ms.Priority,
		NextSendTime: util.FormatTime(ms.NextSendTime),
		Status:       ms.Status,
		CreateTime:   util.FormatTime(ms.CreateTime),
		UpdateTime:   util.FormatTime(ms.UpdateTime),
	}
	resp.RespMessageInfo = tmpInfo
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

// Send
/**
@Summary send a message now, this will change message to a high priority

@Router /api/message/send/{id} [post]
*/
func Send(ctx server.Context) {
	log.Debug("send message[manual]")
	var resp app.Response
	id, err := ctx.C.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = tool.RespCodeNotFound
		//resp.Message = tool.RespMsgNotFount
		ctx.Json(resp)
		return
	}
	log.Info("set message high priority:%d", id)
	_, err = repo.UpdateMessagePriority(id, model.MessagePriorityHigh)
	if err != nil {
		log.Info("set message priority:%d failed", id)
		log.Info(err)
		resp.Code = -1
		resp.Message = "failed"
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
