package graphql

import (
	"github.com/osm/reoui/config"
)

type Resolver struct {
	dataDir string
	cameras []config.Camera
}

func NewResolver(opts ...ResolverOption) *Resolver {
	r := &Resolver{}

	for _, opt := range opts {
		opt(r)
	}

	return r
}
