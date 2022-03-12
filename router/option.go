package router

import (
	"embed"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/osm/reoui/reolink"
)

type Option func(*Router)

func WithGraphql(graphql *handler.Server) Option {
	return func(r *Router) {
		r.graphql = graphql
	}
}

func WithFrontend(frontendFS embed.FS) Option {
	return func(r *Router) {
		r.frontendFS = frontendFS
	}
}

func WithDataDir(dataDir string) Option {
	return func(r *Router) {
		r.dataDir = dataDir
	}
}

func WithReolinks(c []*reolink.Client) Option {
	return func(r *Router) {
		r.reolinks = c
	}
}
