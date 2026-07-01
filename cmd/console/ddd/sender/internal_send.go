package sender

import (
	"errors"

	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd/um"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/sdk"
)

func verifyOpenClient(appId, secret string) (tenantCode string, err error) {
	if appId == "" || secret == "" {
		return "", errors.New("appId and secret required")
	}
	client, err := repo.GetClientByAppId(appId)
	if err != nil || client.Secret != secret {
		return "", errors.New("invalid appId or secret")
	}
	return client.TenantCode, nil
}

func internalSendMail(req testMailReq) (sdk.Response, error) {
	tenantCode, err := verifyOpenClient(req.AppId, req.Secret)
	if err != nil {
		return sdk.Response{Code: 401, Message: err.Error()}, nil
	}
	id, err := um.SendMailNow(tenantCode, req.Template, req.Title, req.Receiver, req.TemplateParam)
	if err != nil {
		return sdk.Response{Code: 500, Message: um.FormatSendError(err)}, nil
	}
	return sdk.Response{Code: 200, Message: "OK", MessageId: id}, nil
}

func internalSendSms(req testSmsReq) (sdk.Response, error) {
	tenantCode, err := verifyOpenClient(req.AppId, req.Secret)
	if err != nil {
		return sdk.Response{Code: 401, Message: err.Error()}, nil
	}
	id, err := um.SendSmsNow(tenantCode, req.Template, req.Receiver, req.TemplateParam)
	if err != nil {
		return sdk.Response{Code: 500, Message: um.FormatSendError(err)}, nil
	}
	return sdk.Response{Code: 200, Message: "OK", MessageId: id}, nil
}

func internalSendIm(req testImReq) (sdk.Response, error) {
	tenantCode, err := verifyOpenClient(req.AppId, req.Secret)
	if err != nil {
		return sdk.Response{Code: 401, Message: err.Error()}, nil
	}
	id, err := um.SendImNow(tenantCode, req.Template, req.Receiver, req.TemplateParam)
	if err != nil {
		return sdk.Response{Code: 500, Message: um.FormatSendError(err)}, nil
	}
	return sdk.Response{Code: 200, Message: "OK", MessageId: id}, nil
}
