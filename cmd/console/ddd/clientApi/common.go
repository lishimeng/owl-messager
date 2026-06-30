package clientApi

type reqClient struct {
	TenantCode string `json:"tenantCode,omitempty"`
	AppId      string `json:"appId,omitempty"`
	Name       string `json:"name,omitempty"`
}

type respClient struct {
	Id         int    `json:"id,omitempty"`
	TenantCode string `json:"tenantCode,omitempty"`
	AppId      string `json:"appId,omitempty"`
	Name       string `json:"name,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
}
