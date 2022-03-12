package clean

import (
	"time"
)

type Option func(*clean)

func WithDataDir(dataDir string) Option {
	return func(c *clean) {
		c.dataDir = dataDir
	}
}

func WithCleanFilesInterval(i time.Duration) Option {
	return func(c *clean) {
		c.cleanFilesInterval = i
	}
}
