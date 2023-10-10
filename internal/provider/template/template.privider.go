package template

import (
	"bytes"
	html "html/template"
	text "text/template"
)

func Rend(data interface{}, temp string) (content string, err error) {
	content, err = RendHtml(data, temp)
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
