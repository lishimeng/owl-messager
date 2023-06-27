package sender

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func svsGetSenders(org int) (senders []model.SenderInfo, err error) {

	_, err = app.GetOrm().Context.
		QueryTable(new(model.SenderInfo)).
		Filter("").All(&senders)
	return
}
