package user

import (
	"awesomeProject/internal/configs"
	"awesomeProject/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.Engine, config *configs.Config, h *Handler) {
	g.GET("/users", middlewares.JWTAuth(config), h.Find)
	g.POST("/users", h.Create)
	g.PUT("/users/:user-id", h.Update)
}
