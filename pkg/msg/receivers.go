package msg

import "strings"

func SplitReceivers(raw string) []string {
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func JoinReceivers(receivers []string) string {
	if len(receivers) == 0 {
		return ""
	}
	return strings.Join(receivers, ",")
}
