package templateApi

import (
	"encoding/json"
	"github.com/lishimeng/app-starter"
	"strings"
)

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

type Param struct {
	Attr []string `json:"attr,omitempty"`
}

type Params map[string]Param

type respTemplate struct {
	app.Response
	Item TemplateResp `json:"item"`
}

func paramsToMap(params string) (string, error) {
	m := make(Params)
	param := strings.Split(params, ",")
	for _, part := range param {
		key := strings.TrimSpace(part)
		if key == "" {
			continue
		}
		m[key] = Param{Attr: []string{key}}
	}
	bs, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func mapToParams(paramsMap string) (p string, err error) {
	var m Params
	err = json.Unmarshal([]byte(paramsMap), &m)
	if err != nil {
		return
	}
	for key, _ := range m {
		p = p + key + ", "
	}
	return
}
