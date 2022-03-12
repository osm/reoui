package sync

import (
	"time"

	"github.com/osm/reoui/reolink"
)

type Option func(*sync)

func WithDataDir(dataDir string) Option {
	return func(s *sync) {
		s.dataDir = dataDir
	}
}

func WithSyncInterval(i time.Duration) Option {
	return func(s *sync) {
		s.syncInterval = i
	}
}

func WithReolinks(c []*reolink.Client) Option {
	return func(s *sync) {
		s.reolinks = c
	}
}
