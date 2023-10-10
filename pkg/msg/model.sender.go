package msg

import (
	"encoding/base64"
)

type MessageSender struct {
	Code     string          // 编号, 不可用ID, 防止数据库编号绑定
	Category MessageCategory // 消息类型
	Provider MessageProvider // 消息平台
	Default  int             // 默认账号 unique:org+vendor
	Config   SenderConfig    // json 配置(map: key-value)
}

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
