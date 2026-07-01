package msg

import (
	"fmt"
	"strings"
)

// FromEmailAddress builds Tencent SES FromEmailAddress.
// Format: "alias <email@domain.com>" or "email@domain.com".
func (c TencentConfig) FromEmailAddress() (string, error) {
	email := strings.TrimSpace(c.SenderEmail)
	alias := strings.TrimSpace(c.SenderAlias)
	legacy := strings.TrimSpace(c.Sender)

	if email == "" && legacy != "" {
		if parsedAlias, parsedEmail, ok := parseAliasEmail(legacy); ok {
			email = parsedEmail
			if alias == "" {
				alias = parsedAlias
			}
		} else {
			email = legacy
		}
	}
	if email == "" {
		return "", fmt.Errorf("发信地址未配置：请在发件人配置中填写 senderEmail（腾讯云 SES 已验证地址）")
	}
	if alias == "" {
		return email, nil
	}
	if strings.Contains(alias, ":") {
		return "", fmt.Errorf("发件人别名不能包含冒号")
	}
	return alias + " <" + email + ">", nil
}

func parseAliasEmail(s string) (alias, email string, ok bool) {
	lt := strings.Index(s, "<")
	gt := strings.LastIndex(s, ">")
	if lt < 0 || gt <= lt {
		return "", "", false
	}
	alias = strings.TrimSpace(s[:lt])
	email = strings.TrimSpace(s[lt+1 : gt])
	if email == "" {
		return "", "", false
	}
	return alias, email, true
}
