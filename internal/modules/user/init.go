package user

import (
	"awesomeProject/internal/model"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine, userRepo model.UserRepository) {
	handler := NewHandler(userRepo)
	RegisterRoutes(g, handler)
}