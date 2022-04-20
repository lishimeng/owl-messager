package util

import "testing"

func TestJoin(t *testing.T) {
	s := Join("-", "a", "s", "d")
	if s != "a-s-d" {
		t.Fatal("not match")
	}
	t.Logf(s)
}
