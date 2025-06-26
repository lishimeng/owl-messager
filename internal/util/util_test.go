package util

import "testing"

func TestJoin(t *testing.T) {
	s := Join("-", "a", "s", "d")
	if s != "a-s-d" {
		t.Fatal("not match")
	}
	t.Logf(s)
}

func TestFormRequest(t *testing.T) {

	host := "https://www.baidu.com"
	params := make(map[string]string)
	params["a"] = "a"
	params["b"] = "b"
	params["c"] = "c"
	code, body, err := New().Form(host, params, nil)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(code)
	t.Logf(body)
}
