package model

type ImMessageInfo struct {
	TemplateChannelDetail
}

func (ImMessageInfo) TableName() string {
	return "im_message_info"
}
