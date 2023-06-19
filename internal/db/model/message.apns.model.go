package model

// ApnsMessageInfo APNS
type ApnsMessageInfo struct {
	MessageHeader
	ApnsMode  ApnsMode `orm:"column(apns_mode)"`   // develop mode / production mode
	BundleId  string   `orm:"column(bundle_id)"`   // bundle id or application
	Params    string   `orm:"column(params);null"` // json params
	Sender    int      `orm:"column(sender_id)"`   // sender's Id
	Receivers string   `orm:"column(receiver)"`    // receiver list. comma split
	Subject   string   `orm:"column(subject)"`     // 主题
}

type ApnsMode int

const (
	DevelopMode    ApnsMode = 0
	ProductionMode ApnsMode = 1
)
