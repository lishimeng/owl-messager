package messageApi

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
)

func GetMessageList(ctx iris.Context) {
	// TODO
}

func GetMessageInfo(ctx iris.Context) {

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
		resp.Code = -1
		resp.Message = "id must be a int value"
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

	common.ResponseJSON(ctx, resp)
}
