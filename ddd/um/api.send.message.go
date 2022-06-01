package um

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
)

type Req struct {
	Template      string      `json:"template"`          // 模板
	TemplateParam interface{} `json:"params"`            // 参数
	Title         string      `json:"subject,omitempty"` // 标题
	Receiver      string      `json:"receiver"`          // 接收者，多个时用逗号分隔
	Category      int         `json:"category"`          // 消息类型
}

type Resp struct {
	app.Response
	MessageId int `json:"messageId,omitempty"`
}

func sendMessage(ctx iris.Context) {
	log.Info("Union message send function")
}
