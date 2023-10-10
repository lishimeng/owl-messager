package messager

import "github.com/lishimeng/owl-messager/internal/db/model"

type MailRequest struct {
	Subject   string
	Receivers []string
	Template  model.MessageTemplate
	Params    map[string]any
}

type MailProvider interface {
	Send(req MailRequest) error
}

// Request SMS发送请求
type Request struct {
	Template  model.MessageTemplate
	Params    map[string]any
	Receivers string
}

// Response 服务器回复
type Response struct {
	RequestId string      // 本次请求的唯一标识，由服务器分配。用来追溯历史
	Payload   interface{} // 个性化的服务器返回信息
}

// SmsProvider 发短信工具
type SmsProvider interface {
	Send(req Request) (resp Response, err error)
}
