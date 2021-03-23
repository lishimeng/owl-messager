package senderApi

import (
	"github.com/kataras/iris"
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

}

func GetMailSenderInfo(ctx iris.Context) {
	log.Debug("get mail sender")
	var resp InfoWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = -1
		resp.Message = "id must be a int value"
		common.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetMailSenderById(id)
	if err != nil {
		log.Debug("get mail sender account failed")
		log.Debug(err)
		resp.Response.Code = -1
		resp.Message = "sender account not exist"
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

	common.ResponseJSON(ctx, resp)
}

func AddMailSender(ctx iris.Context) {

}

func UpdateMailSender(ctx iris.Context) {

}

func DeleteMailSender(ctx iris.Context) {

}
