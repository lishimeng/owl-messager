package model

// ApnsMessageInfo APNS 渠道详情；主题见 MessageInfo.Subject
type ApnsMessageInfo struct {
	PushChannelDetail
}

func (ApnsMessageInfo) TableName() string {
	return "apns_message_info"
}

type ApnsMode int

const (
	DevelopMode    ApnsMode = 0
	ProductionMode ApnsMode = 1
)
