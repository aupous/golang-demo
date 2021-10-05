package user

import (
	"awesomeProject/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	UserRepo model.UserRepository
}

func NewHandler(userRepo model.UserRepository) *Handler {
	return &Handler{UserRepo: userRepo}
}

func (h *Handler) Create(c *gin.Context)  {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"status": http.StatusBadRequest,
		})
		return
	}
	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err := h.UserRepo.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"status": http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}