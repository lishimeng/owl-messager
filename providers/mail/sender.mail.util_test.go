package mail

import (
	"strings"
	"testing"

	"github.com/lishimeng/owl-messager/internal/db/model"
)

func TestBuildMailBody(t *testing.T) {
	tpl := model.MessageTemplate{
		Body: `<html><body><p>{{ .code }}</p></body></html>`,
	}
	body, err := buildMailBody(tpl, map[string]interface{}{"code": "123456"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(body, "123456") {
		t.Fatalf("expected rendered code in body, got: %s", body)
	}
}

func TestBuildMailBodyEmpty(t *testing.T) {
	tpl := model.MessageTemplate{Body: ""}
	_, err := buildMailBody(tpl, map[string]interface{}{"code": "1"})
	if err == nil {
		t.Fatal("expected error for empty template body")
	}
}
