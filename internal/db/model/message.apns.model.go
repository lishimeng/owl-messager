package model

// ApnsMessageInfo APNS
type ApnsMessageInfo struct {
	MessageHeader
	ApnsMode  ApnsMode `gorm:"column:apns_mode"`
	BundleId  string   `gorm:"column:bundle_id"`
	Params    string   `gorm:"column:params"`
	Sender    int      `gorm:"column:sender_id"`
	Receivers string   `gorm:"column:receiver"`
	Subject   string   `gorm:"column:subject"`
}

type ApnsMode int

const (
	DevelopMode    ApnsMode = 0
	ProductionMode ApnsMode = 1
)
