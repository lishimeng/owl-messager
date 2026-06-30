package task

import "strings"

var activeStrategy = MemQueue

// Configure selects the message channel once at startup.
// channel "db" uses DB polling; anything else (including empty) uses in-memory queue.
func Configure(channel string, scanInterval int) []Option {
	if strings.EqualFold(strings.TrimSpace(channel), "db") {
		activeStrategy = Db
		interval := scanInterval
		if interval <= 0 {
			interval = DefaultDbScanInterval
		}
		return []Option{WithDb(interval)}
	}
	activeStrategy = MemQueue
	return []Option{WithQueue(DefaultQueueSize)}
}

func UseMemQueue() bool {
	return activeStrategy == MemQueue
}
