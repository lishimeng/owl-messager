package tpl

import (
	"bytes"
	"fmt"
	"text/template"
)

func Render(data interface{}, temp string) {

	t, err := template.New("123").Parse(temp)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := new(bytes.Buffer)
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(w.String())
	return
}
