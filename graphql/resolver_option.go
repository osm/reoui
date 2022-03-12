package graphql

import (
	"github.com/osm/reoui/config"
)

type ResolverOption func(*Resolver)

func WithDataDir(dataDir string) ResolverOption {
	return func(r *Resolver) {
		r.dataDir = dataDir
	}
}

func WithCameras(c []config.Camera) ResolverOption {
	return func(r *Resolver) {
		r.cameras = c
	}
}
