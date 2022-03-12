package graphql

type Option func(*Server)

func WithResolver(resolver *Resolver) Option {
	return func(s *Server) {
		s.resolver = resolver
	}
}
