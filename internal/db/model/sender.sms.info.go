package model

import "encoding/json"

// SmsSenderInfo 短信发送账号
type SmsSenderInfo struct {
	SenderInfo
	Vendor    SmsVendor         `orm:"column(vendor)"` // 消息平台
	ConfigMap map[string]string `orm:"-"`
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

func (sms SmsVendor) String() string {
	return string(sms)
}

const (
	SmsVendorEnable = 0x01
)

var SmsVendors map[SmsVendor]byte

func init() {
	SmsVendors = make(map[SmsVendor]byte)
	SmsVendors[SmsVendorBaidu] = SmsVendorEnable
	SmsVendors[SmsVendorAli] = SmsVendorEnable
	SmsVendors[SmsVendorHuawei] = SmsVendorEnable
	SmsVendors[SmsVendorQiNiu] = SmsVendorEnable
	SmsVendors[SmsVendorTencent] = SmsVendorEnable
	SmsVendors[SmsVendorUpYun] = SmsVendorEnable
}

// Support 数据库记录的vendor是否被支持
func (s *SmsSenderInfo) Support() bool {
	val, ok := SmsVendors[s.Vendor]
	return ok && (val > 0)
}

const (
	SmsSenderEnable  = 1 // enable
	SmsSenderDisable = 0 // disable
)

type AliSmsConfig struct {
	AppKey    string `json:"appKey,omitempty"`
	AppSecret string `json:"appSecret,omitempty"`
	Region    string `json:"region,omitempty"`
	SignName  string `json:"signName,omitempty"`
}

type TencentSmsConfig struct {
	AppId    string `json:"appId,omitempty"`
	AppKey   string `json:"appKey,omitempty"`
	SmsAppId string `json:"smsAppId,omitempty"`
	Region   string `json:"region,omitempty"`
	SignName string `json:"signName,omitempty"`
}
type HuaweiSmsConfig struct {
	Host     string `json:"host,omitempty"`
	AppId    string `json:"appId,omitempty"`
	AppKey   string `json:"appKey,omitempty"`
	Sender   string `json:"sender,omitempty"`
	SignName string `json:"signName,omitempty"`
}
