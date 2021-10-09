package auth

import (
	"awesomeProject/internal/configs"
	responsehelper "awesomeProject/internal/helpers"
	"awesomeProject/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Handler struct {
	UserRepo model.UserRepository
	Config   *configs.Config
}

func NewHandler(userRepo model.UserRepository, config *configs.Config) *Handler {
	return &Handler{UserRepo: userRepo, Config: config}
}

func (h *Handler) SignIn(c *gin.Context) {
	var req SignInRequset
	if c.Bind(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Bad request",
			"status": http.StatusBadRequest,
		})
		return
	}
	/*
		Tìm thằng user có email gửi lên
		- Nếu không có -> trả lỗi luôn
		- Nếu có -> So sánh password
	*/
	user, err := h.UserRepo.FindByEmail(req.Email)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"request": req,
			"error":   err,
		}).Error("failed to find user by email")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Internal Server Error",
			"status": http.StatusInternalServerError,
		})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "User not exists",
			"status": http.StatusUnauthorized,
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "Password is not correct",
			"status": http.StatusUnauthorized,
		})
		return
	}
	token, err := CreateToken(user, h.Config.JWTSecret)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user":  user,
			"error": err,
		}).Error("failed to create user token")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Internal Server Error",
			"status": http.StatusInternalServerError,
		})
		return
	}
	responsehelper.Response(c, SignInResponse{
		User:  *user,
		Token: token,
	}, "Success")
}
