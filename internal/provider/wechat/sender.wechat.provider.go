package wechat

type Provider struct {
	Host string// wechat公众号provider地址
}

type PayloadItem struct {
	Content string `json:"content"`
	Color string `json:"color,omitempty"`
}
type To struct {
	OpenId string // 公众号中的用户唯一ID
	Uid string // 平台用户ID
	Payload map[string]PayloadItem // 数据
	TemplateId string // 公众平台上的模板
}
