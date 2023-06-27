package template

import (
	"github.com/lishimeng/owl-messager/internal/db/model"
	"testing"
)

func TestRender(t *testing.T) {
	data := make(map[string]interface{})
	tmp := `haha {{ .A }}--- {{ .C }}`
	data["A"] = "aaa"
	data["C"] = "bbb"
	c, err := Rend(data, tmp, model.MailTemplateCategoryHtml)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(c)
}

func TestMapLowerCase(t *testing.T) {
	data := make(map[string]interface{})
	tmp := `haha {{ .a }}--- {{ .c }}`
	data["a"] = "aaa"
	data["c"] = "bbb"
	c, err := Rend(data, tmp, model.MailTemplateCategoryHtml)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(c)
}

type Params struct {
	A string
	B string
}

func TestStruct(t *testing.T) {
	var data Params
	tmp := `haha {{ .A }}--- {{ .A }}`
	data.A = "aaa"
	data.B = "bbb"
	c, err := Rend(data, tmp, model.MailTemplateCategoryHtml)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(c)
}
