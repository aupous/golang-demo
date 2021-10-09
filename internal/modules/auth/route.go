package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.Engine, h *Handler) {
	g.POST("/auth/sign-in", h.SignIn)
}
