package util

import "strings"

func Join(j string, s ...string) string {
	return strings.Join(s, j)
}

// StatusIn 判定状态是否在列表中
func StatusIn(status int, list []int) (b bool) {
	for _, v := range list {
		if status == v {
			b = true
			break
		}
	}
	return
}
