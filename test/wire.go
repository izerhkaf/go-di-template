//go:build wireinject
// +build wireinject

package test

import (
	"template/internal/config"
	"template/internal/handler"
	"template/internal/middleware"
	"template/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeServer() (*gin.Engine, error) {
	wire.Build(
		config.ProviderSetTest,
		// database.ProviderSet,
		// repository.ProviderSet,
		// service.ProviderSet,
		handler.ProviderSet,
		routes.ProviderSet,
		middleware.ProviderSet,
	)

	return &gin.Engine{}, nil
}
