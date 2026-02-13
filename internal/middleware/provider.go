package middleware

import "github.com/google/wire"

type Middlewares struct {
	CorsMiddleware *CorsMiddleware
}

func NewMiddlewares(
	cors *CorsMiddleware,
) *Middlewares {
	return &Middlewares{
		CorsMiddleware: cors,
	}
}

var ProviderSet = wire.NewSet(
	NewMiddlewares,
	NewCorsMiddleware,
)
