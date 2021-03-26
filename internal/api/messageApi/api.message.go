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

func GetMessageList(ctx iris.Context) {
	// TODO
}

func GetMessageInfo(ctx iris.Context) {
	log.Debug("get message")
	var resp RespMessageInfoWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = common.RespCodeNotFound
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
		Id:         ms.Id,
		Category: ms.Category,
		Subject:       ms.Subject,
		Priority:       ms.Priority,
		NextSendTime:      common.FormatTime(ms.NextSendTime),
		Status:     ms.Status,
		CreateTime: common.FormatTime(ms.CreateTime),
		UpdateTime: common.FormatTime(ms.UpdateTime),
	}
	resp.RespMessageInfo = tmpInfo
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

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
