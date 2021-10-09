package server

import (
	"awesomeProject/internal/configs"
	"awesomeProject/internal/db"
	"awesomeProject/internal/modules/auth"
	"awesomeProject/internal/modules/user"
	"awesomeProject/internal/repository"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Start() {
	r := gin.Default()
	config := configs.LoadEnv()
	pgDB := db.DB{}
	pgDB.Connect(config.PostgresConfig)

	userRepo := repository.NewUserRepository(&pgDB)
	user.Init(r, config, userRepo)
	auth.Init(r, config, userRepo)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(); err != nil {
		panic(err)
	}
}
