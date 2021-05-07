package jade

import (
	"bytes"
	"github.com/iris-contrib/jade"
	"html/template"
)

func Render(data interface{}, tpl string) (txt string, err error) {
	htm, err := jade.Parse("_", []byte(tpl))
	if err != nil {
		return
	}
	t, err := template.New("_").Parse(htm)
	if err != nil {
		return
	}
	w := new(bytes.Buffer)
	err = t.Execute(w, data)

	if err != nil {
		return
	}
	txt = w.String()
	return
}
