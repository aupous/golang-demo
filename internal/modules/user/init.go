package user

import (
	"awesomeProject/internal/configs"
	"awesomeProject/internal/model"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine, config *configs.Config, userRepo model.UserRepository) {
	handler := NewHandler(userRepo)
	RegisterRoutes(g, config, handler)
}
