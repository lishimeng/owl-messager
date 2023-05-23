package handlebar

import "github.com/aymerick/raymond"

func Render(data interface{}, tpl string) (txt string, err error) {

	txt, err = raymond.Render(tpl, data)
	return
}
