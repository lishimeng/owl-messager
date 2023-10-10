package sms

import (
	"fmt"
	"sort"
	"strings"
)

func map2array(m map[string]any) (arr []string) {

	// {1}为您的登录验证码，请于{2}分钟内填写，如非本人操作，请忽略本短信。
	// 参数名为数字
	// key: 参数index
	// 参数不能超过10个

	var keys []string
	for key := range m {
		keys = append(keys, fmt.Sprintf("%0*s", 1, key))
	}

	sort.Strings(keys)

	for _, key := range keys {
		value := fmt.Sprintf("%+v", m[key])
		arr = append(arr, value)
	}
	return
}

func buildHuaweiTemplateParams(params []string) (s string) {
	var p []string
	for _, v := range params {
		p = append(p, fmt.Sprintf(`"%s"`, v))
	}
	tmp := strings.Join(p, ",")
	s = fmt.Sprintf("[%s]", tmp)
	return
}
