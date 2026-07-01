package templateApi

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNormalizeParamsShorthand(t *testing.T) {
	got, err := normalizeParams("userName, code")
	if err != nil {
		t.Fatal(err)
	}
	var m Params
	if err = json.Unmarshal([]byte(got), &m); err != nil {
		t.Fatal(err)
	}
	if len(m) != 2 {
		t.Fatalf("expected 2 keys, got %d", len(m))
	}
	if len(m["userName"].Attr) != 1 || m["userName"].Attr[0] != "userName" {
		t.Fatalf("unexpected userName mapping: %+v", m["userName"])
	}
}

func TestNormalizeParamsJSON(t *testing.T) {
	input := `{
		"userName": {
			"description": "用户名",
			"attr": ["UserName"]
		}
	}`
	got, err := normalizeParams(input)
	if err != nil {
		t.Fatal(err)
	}
	var m Params
	if err = json.Unmarshal([]byte(got), &m); err != nil {
		t.Fatal(err)
	}
	if m["userName"].Attr[0] != "UserName" {
		t.Fatalf("unexpected attr: %+v", m["userName"])
	}
}

func TestNormalizeParamsJSONAutoAttr(t *testing.T) {
	got, err := normalizeParams(`{"code": {"description": "验证码"}}`)
	if err != nil {
		t.Fatal(err)
	}
	var m Params
	if err = json.Unmarshal([]byte(got), &m); err != nil {
		t.Fatal(err)
	}
	if m["code"].Attr[0] != "code" {
		t.Fatalf("expected auto attr, got %+v", m["code"])
	}
}

func TestFormatParamsForDisplay(t *testing.T) {
	got, err := formatParamsForDisplay(`{"code":{"attr":["Code"]}}`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "\n") {
		t.Fatalf("expected pretty json, got %s", got)
	}
}
