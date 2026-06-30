package um

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/mailattachment"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type uploadAttachmentResp struct {
	app.Response
	Item msg.MailAttachmentRef `json:"item"`
}

func uploadMailAttachment(ctx server.Context) {
	var resp uploadAttachmentResp
	orgHeader := ctx.C.GetHeader(auth.OrgKey)
	tenant, err := repo.GetTenant(orgHeader)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "unknown tenant"
		ctx.Json(resp)
		return
	}

	file, info, err := ctx.C.FormFile("file")
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "file required"
		ctx.Json(resp)
		return
	}
	defer file.Close()

	ref, err := mailattachment.Default().Save(tenant.Code, info.Filename, info.Header.Get("Content-Type"), file)
	if err != nil {
		log.Info("upload attachment failed: %v", err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Item = ref
	resp.Code = tool.RespCodeSuccess
	resp.Message = "OK"
	ctx.Json(resp)
}
