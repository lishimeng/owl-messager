package mail

import (
	"errors"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/provider/template"
)

func buildMailBody(tpl model.MessageTemplate, params map[string]interface{}) (body string, err error) {

	body, err = template.Rend(params, tpl.Body)
	if err != nil {
		log.Info("template render failed")
		log.Info(err)
		return
	}
	if len(body) <= 0 {
		log.Info("mail body empty")
		err = errors.New("mail body empty")
		return
	}
	return
}
