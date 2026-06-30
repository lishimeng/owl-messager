package consoleauth

import (
	"crypto/subtle"
	"sync"

	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

var (
	mu    sync.RWMutex
	token string
)

// Load reads console management token from config table (code=console.token).
func Load() {
	settings := repo.LoadConsoleTokenSettings()
	mu.Lock()
	token = settings.Token
	mu.Unlock()
	if token == "" {
		log.Warn("console token not configured (config.code=console.token)")
	}
}

func Valid(raw string) bool {
	mu.RLock()
	want := token
	mu.RUnlock()
	if want == "" || raw == "" {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(raw), []byte(want)) == 1
}
