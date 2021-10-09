package auth

import (
	"awesomeProject/internal/configs"
	"awesomeProject/internal/model"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine, config *configs.Config, userRepo model.UserRepository) {
	handler := NewHandler(userRepo, config)
	RegisterRoutes(g, handler)
}
