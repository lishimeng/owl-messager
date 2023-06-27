package template

import (
	"bytes"
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	html "html/template"
	text "text/template"
)

func Rend(data interface{}, temp string, category int) (content string, err error) {
	switch category {
	case model.MailTemplateCategoryText:
		content, err = RendText(data, temp)
	case model.MailTemplateCategoryHtml:
		content, err = RendHtml(data, temp)
	default:
		log.Debug("unknown mail template category:%d", category)
		err = fmt.Errorf("unknown mail template category:%d", category)
		return
	}

	return
}

func RendText(data interface{}, temp string) (content string, err error) {
	t, err := text.New("_").Parse(temp)
	if err != nil {
		return
	}
	w := new(bytes.Buffer)
	err = t.Execute(w, data)
	if err != nil {
		return
	}
	content = w.String()
	return
}

func RendHtml(data interface{}, temp string) (content string, err error) {
	t, err := html.New("_").Parse(temp)
	if err != nil {
		return
	}
	w := new(bytes.Buffer)
	err = t.Execute(w, data)
	if err != nil {
		return
	}
	content = w.String()
	return
}
