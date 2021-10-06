package user

import (
	"awesomeProject/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *Handler) Update(c *gin.Context) {
	userID := c.Param("user-id")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"status": http.StatusInternalServerError,
		})
	}
	var req UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"status": http.StatusBadRequest,
		})
		return
	}
	user := model.User{
		ID:       userUUID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := h.UserRepo.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Bad request",
			"status": http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": user,
	})
	return
}

func (h *Handler) Find(c *gin.Context) {
	//search := c.Query("search")

	//users, err := h.UserRepo.Find(search)
}