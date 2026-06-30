package mailattachment

import "time"

type Config struct {
	Dir             string
	MaxFileSize     int64
	MaxTotalSize    int64
	MaxCount        int
	StagingTTL      time.Duration
	RetainAfterSend time.Duration
}

func (c Config) withDefaults() Config {
	if c.Dir == "" {
		c.Dir = "data/mail-attachments"
	}
	if c.MaxFileSize <= 0 {
		c.MaxFileSize = 10 * 1024 * 1024
	}
	if c.MaxTotalSize <= 0 {
		c.MaxTotalSize = 20 * 1024 * 1024
	}
	if c.MaxCount <= 0 {
		c.MaxCount = 5
	}
	if c.StagingTTL <= 0 {
		c.StagingTTL = 24 * time.Hour
	}
	return c
}
