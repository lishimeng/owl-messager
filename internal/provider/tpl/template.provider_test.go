package tpl

import "testing"

func TestRender(t *testing.T) {
	data := make(map[string]interface{})
	tmp := `haha {{ .A }}--- {{ .C }}`
	data["A"] = "aaa"
	data["B"] = "bbb"
	Render(data, tmp)
}
