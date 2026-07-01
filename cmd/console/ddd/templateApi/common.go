package templateApi

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lishimeng/app-starter"
)

type TemplateResp struct {
	Id            int    `json:"id,omitempty"`
	TenantCode    string `json:"tenantCode,omitempty"`
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
	TenantCode    string `json:"tenantCode,omitempty"`
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
	Attr        []string `json:"attr,omitempty"`
	Description string   `json:"description,omitempty"`
}

type Params map[string]Param

type respTemplate struct {
	app.Response
	Item TemplateResp `json:"item"`
}

// normalizeParams 接受 JSON 映射或逗号简写（userName, code），返回紧凑 JSON 存库。
func normalizeParams(params string) (string, error) {
	params = strings.TrimSpace(params)
	if params == "" {
		return "{}", nil
	}
	if strings.HasPrefix(params, "{") {
		return normalizeParamsJSON(params)
	}
	return paramsToMap(params)
}

func normalizeParamsJSON(params string) (string, error) {
	var m Params
	if err := json.Unmarshal([]byte(params), &m); err != nil {
		return "", fmt.Errorf("invalid params json: %w", err)
	}
	for key, p := range m {
		if strings.TrimSpace(key) == "" {
			return "", fmt.Errorf("param key cannot be empty")
		}
		if len(p.Attr) == 0 {
			m[key] = Param{
				Attr:        []string{key},
				Description: p.Description,
			}
		}
	}
	bs, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func formatParamsForDisplay(stored string) (string, error) {
	stored = strings.TrimSpace(stored)
	if stored == "" {
		return "{}", nil
	}
	var m Params
	if err := json.Unmarshal([]byte(stored), &m); err != nil {
		return stored, nil
	}
	bs, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return stored, nil
	}
	return string(bs), nil
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
