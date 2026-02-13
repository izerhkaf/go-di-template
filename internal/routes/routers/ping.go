package routers

import (
	"template/internal/handler"

	"github.com/gin-gonic/gin"
)

func PingRoutes(group *gin.RouterGroup, h *handler.Handlers) {
	group.GET("/ping", h.PingHandler.Ping)
}
