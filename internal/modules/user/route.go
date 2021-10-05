package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.Engine ,h *Handler) {
	g.POST("/users", h.Create)
}