package template

import (
	"github.com/lishimeng/owl/internal/db/model"
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
