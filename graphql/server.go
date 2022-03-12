package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/osm/reoui/graphql/generated"
)

type Server struct {
	resolver *Resolver
}

var mb int64 = 1 << 20

func NewServer(opts ...Option) *handler.Server {
	s := &Server{}

	for _, opt := range opts {
		opt(s)
	}

	srv := handler.New(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: s.resolver},
		),
	)

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		return graphql.DefaultErrorPresenter(ctx, e)
	})

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return srv
}
