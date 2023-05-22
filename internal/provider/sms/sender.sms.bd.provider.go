package sms

// 百度云SMS
import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/sms"
	"github.com/baidubce/bce-sdk-go/services/sms/api"
	"github.com/lishimeng/go-log"
)

/**
配置项名称	类型	含义
Endpoint	string	请求服务的域名
ProxyUrl	string	客户端请求的代理地址
Region	string	请求资源的区域
UserAgent	string	用户名称，HTTP请求的User-Agent头
Credentials	*auth.BceCredentials	请求的鉴权对象，分为普通AK/SK与STS两种
SignOption	*auth.SignOptions	认证字符串签名选项
Retry	RetryPolicy	连接重试策略
ConnectionTimeoutInMillis	int	连接超时时间，单位毫秒，默认20分钟
*/

type BDSmsProvider struct {
	client    *sms.Client
	signature string
}

func (p *BDSmsProvider) Init(accessKey, accessSecret string, signature string) (err error) {
	endpoint := "https://smsv3.bj.baidubce.com"
	p.client, err = sms.NewClient(accessKey, accessSecret, endpoint)
	p.signature = signature
	p.client.Config.Retry = bce.NewNoRetryPolicy()
	p.client.Config.ConnectionTimeoutInMillis = 30 * 1000
	return
}

func (p *BDSmsProvider) Send(toer string, tplId string, params map[string]interface{}) (err error) {
	resp, err := p.client.SendSms(&api.SendSmsArgs{
		Mobile:      toer,
		Template:    tplId,
		SignatureId: p.signature,
		ContentVar:  params,
	})
	if err != nil {
		log.Info("send sms failed(bd sms)")
		log.Info(err)
		return
	}
	log.Info("response is %#v\n", resp)
	return
}
