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
)

func init() {

	MailProviders = make(map[MessageProvider]byte)
	MailProviders[Smtp] = VendorEnable
	MailProviders[Microsoft] = VendorEnable
	MailProviders[Tencent] = VendorEnable

	SmsProviders = make(map[MessageProvider]byte)
	SmsProviders[Baidu] = VendorEnable
	SmsProviders[Ali] = VendorEnable
	SmsProviders[Tencent] = VendorEnable
	SmsProviders[Huawei] = VendorEnable
	SmsProviders[UpYun] = VendorEnable
	SmsProviders[QiNiu] = VendorEnable

	Providers = make(map[MessageCategory]map[MessageProvider]byte)
	Providers[MailMessage] = MailProviders
	Providers[SmsMessage] = SmsProviders
}
