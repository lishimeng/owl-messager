package model

import "testing"

func TestSenderConfig(t *testing.T) {
	var err error
	content := "4r43423reqrr432rqref"
	s := SenderConfig(content)
	err = s.Encode()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("encode done")
	err = s.Decode()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("decode done")

	if content != string(s) {
		t.Fatalf("not match: %s<-->%s", content, string(s))
	}
	t.Log("match")
}
