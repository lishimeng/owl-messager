package ding

import (
	"net/url"
	"testing"
)

func TestName(t *testing.T) {
	var a = "https://www.ab.com"
	var p = "dfdf/fgds"
	t.Log(url.JoinPath(a, p))
}
