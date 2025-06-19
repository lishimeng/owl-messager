package clientApi

type reqClient struct {
	Org   int    `json:"org,omitempty"`
	AppId string `json:"appId,omitempty"`
	Name  string `json:"name,omitempty"`
}

type respClient struct {
	Id         int    `json:"id,omitempty"`
	Org        int    `json:"org,omitempty"`
	AppId      string `json:"appId,omitempty"`
	Name       string `json:"name,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
}
