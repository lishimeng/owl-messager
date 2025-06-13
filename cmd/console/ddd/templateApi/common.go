package templateApi

type TemplateResp struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Body          string `json:"body,omitempty"`
	CloudTemplate string `json:"cloudTemplate,omitempty"`
	Description   string `json:"description,omitempty"`
	Category      string `json:"category,omitempty"`
	Params        string `json:"params,omitempty"`
	Provider      string `json:"provider,omitempty"`
	Status        int    `json:"status,omitempty"`
	Code          string `json:"code,omitempty"`
	CreateTime    string `json:"createTime,omitempty"`
	UpdateTime    string `json:"updateTime,omitempty"`
}

type TemplateReq struct {
	Name          string `json:"name,omitempty"`
	Body          string `json:"body,omitempty"`
	CloudTemplate string `json:"cloudTemplate,omitempty"`
	Description   string `json:"description,omitempty"`
	Category      string `json:"category,omitempty"`
	Params        string `json:"params,omitempty"`
	Provider      string `json:"provider,omitempty"`
	Status        int    `json:"status,omitempty"`
	Code          string `json:"code,omitempty"`
}
