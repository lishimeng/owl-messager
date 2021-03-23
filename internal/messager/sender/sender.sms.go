package sender

import "github.com/lishimeng/owl/internal/db/model"

type Sms interface {
	Send(model.SmsMessageInfo) (err error)
}
