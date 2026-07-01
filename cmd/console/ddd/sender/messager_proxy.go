package sender

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/lishimeng/owl-messager/internal/etc"
)

func defaultMessagerHost() string {
	return strings.TrimSpace(etc.Config.Messager.Host)
}

func normalizeMessagerHost(host string) string {
	host = strings.TrimSpace(host)
	host = strings.TrimRight(host, "/")
	if strings.HasSuffix(host, "/api") {
		host = strings.TrimSuffix(host, "/api")
		host = strings.TrimRight(host, "/")
	}
	return host
}

func shouldUseInternalSend(host string) bool {
	host = normalizeMessagerHost(host)
	return host == "" || strings.EqualFold(host, "internal")
}

func resolveMessagerHost(host string) string {
	host = normalizeMessagerHost(host)
	if host != "" {
		return host
	}
	return normalizeMessagerHost(defaultMessagerHost())
}

func probeMessager(host string) error {
	host = resolveMessagerHost(host)
	if host == "" {
		return errors.New("未配置 Messager 地址：请在测试表单填写，或在配置 [messager] host 中设置")
	}
	client := &http.Client{Timeout: 5 * time.Second}
	for _, path := range []string{"/messages/ping", "/api/messages/ping"} {
		resp, err := client.Get(host + path)
		if err != nil {
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			continue
		}
		var payload struct {
			Message string `json:"message"`
		}
		if json.Unmarshal(body, &payload) == nil && payload.Message == "owl-messager" {
			return nil
		}
	}
	return fmt.Errorf("地址 %s 未检测到 owl-messager（GET /messages/ping）；请确认该端口运行的是 messager 服务，而非 console（console 仅有 /api/*）", host)
}
