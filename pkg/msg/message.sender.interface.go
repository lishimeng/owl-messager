package msg

type Sender interface {
	Send()
	// 进入空闲
	OnIdle()
}

const (
	VendorDisable = iota
	VendorEnable
)

var (
	Providers     map[MessageCategory]map[MessageProvider]byte
	MailProviders map[MessageProvider]byte
	SmsProviders  map[MessageProvider]byte
	ImProviders   map[MessageProvider]byte
	ApnsProviders map[MessageProvider]byte
)
