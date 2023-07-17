package model

import (
	"encoding/base64"
	"github.com/lishimeng/app-starter"
)

type SenderConfig string

func (s *SenderConfig) Encode() error {
	tmp := base64.StdEncoding.EncodeToString([]byte((*s)))
	*s = SenderConfig(tmp)
	return nil
}
func (s *SenderConfig) Decode() error {
	tmp, err := base64.StdEncoding.DecodeString(string(*s))
	if err != nil {
		return err
	}
	*s = SenderConfig(tmp)
	return nil
}

// SenderInfo 发消息账号
type SenderInfo struct {
	app.TenantPk
	Code    string       `orm:"column(code);unique"`    // 编号
	Default int          `orm:"column(default_sender)"` // 默认账号
	Config  SenderConfig `orm:"column(config)"`         // json 配置(map: key-value)
	app.TableChangeInfo
}

const (
	DefaultSenderDisable = 0
	DefaultSenderEnable  = 1
)

const (
	SenderCategoryMail = "mail"
	SenderCategorySms  = "sms"
)
