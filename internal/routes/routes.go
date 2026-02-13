package routes

import (
	"strings"
	"template/internal/config"
	"template/internal/handler"
	"template/internal/middleware"
	"template/internal/routes/routers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handlers, m *middleware.Middlewares, c *config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(m.CorsMiddleware.AllowCORS())

	origins := c.ORIGINS
	allowed_origins := strings.Split(origins, ",")
	config := cors.Config{
		AllowOrigins: allowed_origins,
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization", "X-Api-Key",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	v1 := r.Group("/v1")
	{
		routers.PingRoutes(v1, h)
	}

	r.OPTIONS("/*cors", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})

	return r
}
