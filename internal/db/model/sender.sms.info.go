package model

// SmsSenderInfo 短信发送账号
type SmsSenderInfo struct {
	Pk
	Code      string    `orm:"column(code);unique"`
	Name      string    `orm:"column(name)"`
	Host      string    `orm:"column(host)"`
	Vendor    SmsVendor `orm:"column(vendor)"` // 消息平台
	AppKey    string    `orm:"column(app_key);null"`
	AppSecret string    `orm:"column(app_secret)"`

	TableChangeInfo
}

type SmsVendor string

const (
	SmsVendorBaidu   SmsVendor = "baidu_yun"
	SmsVendorAli     SmsVendor = "ali_yun"
	SmsVendorHuawei  SmsVendor = "huawei_yun"
	SmsVendorQiNiu   SmsVendor = "qi_niu_yun"
	SmsVendorTencent SmsVendor = "tencent_yun"
	SmsVendorUpYun   SmsVendor = "up_yun"
)

var SmsVendors map[SmsVendor]byte

func init() {
	SmsVendors = make(map[SmsVendor]byte)
	SmsVendors[SmsVendorBaidu] = 0x01
	SmsVendors[SmsVendorAli] = 0x01
	SmsVendors[SmsVendorHuawei] = 0x01
	SmsVendors[SmsVendorQiNiu] = 0x01
	SmsVendors[SmsVendorTencent] = 0x01
	SmsVendors[SmsVendorUpYun] = 0x01
}

// Support 数据库记录的vendor是否被支持
func (s SmsSenderInfo) Support() bool {
	val, ok := SmsVendors[s.Vendor]
	return ok && (val > 0)
}
