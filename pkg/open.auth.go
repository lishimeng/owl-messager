package pkg

import "github.com/lishimeng/app-starter"

type CredentialReq struct {
	AppId  string `json:"appId,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type CredentialResp struct {
	app.Response
	Token string `json:"token,omitempty"`
}
