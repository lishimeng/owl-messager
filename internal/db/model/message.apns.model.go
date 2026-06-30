package model

// ApnsMessageInfo APNS 渠道详情；主题见 MessageInfo.Subject
type ApnsMessageInfo struct {
	PushChannelDetail
}

type ApnsMode int

const (
	DevelopMode    ApnsMode = 0
	ProductionMode ApnsMode = 1
)
