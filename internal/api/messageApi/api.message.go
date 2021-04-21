package messageApi

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
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
		common.ResponseJSON(ctx, resp)
		return
	}
	if len(messages) > 0 {
		for _, ms := range messages {
			var tmpInfo = RespMessageInfo{
				Id:           ms.Id,
				Category:     ms.Category,
				Subject:      ms.Subject,
				Priority:     ms.Priority,
				NextSendTime: common.FormatTime(ms.NextSendTime),
				Status:       ms.Status,
				CreateTime:   common.FormatTime(ms.CreateTime),
				UpdateTime:   common.FormatTime(ms.UpdateTime),
			}
			page.Data = append(page.Data, tmpInfo)
		}
	}

	resp.Pager = page
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)

}

func GetMessageInfo(ctx iris.Context) {
	log.Debug("get message")
	var resp RespMessageInfoWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgIdNum
		common.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetMessageById(id)
	if err != nil {
		log.Debug("get message failed")
		log.Debug(err)
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgNotFount
		common.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = RespMessageInfo{
		Id:           ms.Id,
		Category:     ms.Category,
		Subject:      ms.Subject,
		Priority:     ms.Priority,
		NextSendTime: common.FormatTime(ms.NextSendTime),
		Status:       ms.Status,
		CreateTime:   common.FormatTime(ms.CreateTime),
		UpdateTime:   common.FormatTime(ms.UpdateTime),
	}
	resp.RespMessageInfo = tmpInfo
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
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
		resp.Code = common.RespMsgNotFount
		resp.Message = common.RespMsgNotFount
		common.ResponseJSON(ctx, resp)
		return
	}
	log.Info("set message high priority:%d", id)
	_, err = repo.UpdateMessagePriority(id, model.MessagePriorityHigh)
	if err != nil {
		log.Info("set message priority:%d failed", id)
		log.Info(err)
		resp.Code = -1
		resp.Message = "failed"
		common.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
