package wechat

// 小程序-模板消息

const (
	uniformHostTpl = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=%s"
)

type MaProvider struct {
	token string
}

func (m *MaProvider) ChangeToken() {

}

func (m *MaProvider) Send() {

}
