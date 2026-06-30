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

func init() { // buildin

	MailProviders = make(map[MessageProvider]byte)
	MailProviders[Smtp] = VendorEnable
	MailProviders[Microsoft] = VendorEnable
	MailProviders[Tencent] = VendorEnable

	SmsProviders = make(map[MessageProvider]byte)
	SmsProviders[Baidu] = VendorEnable
	SmsProviders[Ali] = VendorEnable
	SmsProviders[Tencent] = VendorEnable
	SmsProviders[Huawei] = VendorEnable
	// UpYun: provider not registered yet; keep disabled until implemented.
	// SmsProviders[UpYun] = VendorEnable
	SmsProviders[QiNiu] = VendorEnable

	ImProviders = make(map[MessageProvider]byte)
	ImProviders[FastMsg] = VendorEnable

	ApnsProviders = make(map[MessageProvider]byte)
	ApnsProviders[Apns] = VendorEnable

	Providers = make(map[MessageCategory]map[MessageProvider]byte)
	Providers[MailMessage] = MailProviders
	Providers[SmsMessage] = SmsProviders
	Providers[ImMessage] = ImProviders
	Providers[ApnsMessage] = ApnsProviders
}
