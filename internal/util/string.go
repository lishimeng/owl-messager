package util

import "strings"

func Join(j string, s ...string) string {
	return strings.Join(s, j)
}
