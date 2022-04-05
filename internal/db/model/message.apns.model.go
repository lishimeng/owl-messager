package model

// ApnsMessageInfo APNS
type ApnsMessageInfo struct {
	MessageHeader

	ApnsMode ApnsMode `orm:"column(apns_mode)"` // develop mode / production mode

	BundleId string `orm:"column(bundle_id)"` // bundle id or application

	Params string `orm:"column(params);null"` // json params

	Sender int `orm:"column(sender_id)"` // sender's Id

	Receivers string `orm:"column(receiver)"` // receiver list. comma split

	Cc string `orm:"column(cc);null"` // CC

	// 主题
	Subject string `orm:"column(subject)"`
}

type ApnsMode int

const (
	DevelopMode    ApnsMode = 0
	ProductionMode ApnsMode = 1
)
