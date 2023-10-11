package ddd

import "github.com/lishimeng/owl-messager/internal/db/model"

func Tables() (t []interface{}) {
	t = append(t,
		new(model.MessageInfo),
		new(model.MessageTask),
		new(model.MessageRunningTask),
		new(model.OpenClient),
		new(model.SmsMessageInfo),
		new(model.ApnsMessageInfo),
	)
	return
}
