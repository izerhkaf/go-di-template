//go:build wireinject
// +build wireinject

package template

import (
	// "template/internal/config"      // Unused
	// "template/internal/database"    // Unused
	"template/internal/config"
	"template/internal/handler"
	"template/internal/middleware"

	// "template/internal/repository"  // Unused
	"template/internal/routes"
	// "template/internal/service"     // Unused

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeServer() (*gin.Engine, error) {
	wire.Build(
		config.ProviderSet,
		// database.ProviderSet,    // Commented out: not used in dependency chain
		// repository.ProviderSet,  // Commented out: not used in dependency chain
		// service.ProviderSet,     // Commented out: not used in dependency chain
		handler.ProviderSet,
		routes.ProviderSet,
		middleware.ProviderSet,
	)

	return &gin.Engine{}, nil
}
