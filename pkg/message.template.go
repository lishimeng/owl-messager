package pkg

type TemplateInfo struct {
	Id            int    `json:"id,omitempty"`
	Code          string `json:"code,omitempty"`
	Name          string `json:"name,omitempty"`
	Category      string `json:"category,omitempty"`
	Body          string `json:"body,omitempty"`
	Params        string `json:"params,omitempty"`
	Provider      string `json:"provider,omitempty"`
	CloudTemplate string `json:"cloudTemplate,omitempty"`
	Description   string `json:"description,omitempty"`
}
