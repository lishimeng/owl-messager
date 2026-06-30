package authApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

type loginReq struct {
	Token string `json:"token"`
}

type tokenResp struct {
	app.Response
	Token string `json:"token,omitempty"`
}

// verifyToken checks basic_auth and echoes it back for the client to store.
func verifyToken(ctx server.Context) {
	var req loginReq
	var resp tokenResp

	if err := ctx.C.ReadJSON(&req); err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "invalid json"
		ctx.Json(resp)
		return
	}
	if len(req.Token) == 0 {
		resp.Code = tool.RespCodeError
		resp.Message = "token required"
		ctx.Json(resp)
		return
	}

	if _, err := repo.GetClientByBasicAuth(req.Token); err != nil {
		log.Debug("console login rejected: invalid basic_auth")
		resp.Code = 401
		resp.Message = "invalid token"
		ctx.Json(resp)
		return
	}

	resp.Token = req.Token
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
