package model

import "encoding/json"

// SmsSenderInfo 短信发送账号
type SmsSenderInfo struct {
	Pk
	Code      string    `orm:"column(code);unique"`
	Name      string    `orm:"column(name)"`
	Vendor    SmsVendor `orm:"column(vendor)"` // 消息平台
	Config    string    `orm:"column(config)"` // json 配置(map: key-value)
	ConfigMap map[string]string
	TableChangeInfo
}

func (s *SmsSenderInfo) UnmarshalConfig() (err error) {
	err = json.Unmarshal([]byte(s.Config), &(s.ConfigMap))
	return
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

const (
	SmsSenderEnable  = 1 // enable
	SmsSenderDisable = 0 // disable
)
