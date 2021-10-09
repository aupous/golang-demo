package user

import (
	"awesomeProject/internal/constants"
	"awesomeProject/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Handler struct {
	UserRepo model.UserRepository
}

func NewHandler(userRepo model.UserRepository) *Handler {
	return &Handler{UserRepo: userRepo}
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		//responsehelper.ResponseWithValidationError(c, err)
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(password),
	}
	err = h.UserRepo.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Internal Server Error",
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
			"error":  "Internal Server Error",
			"status": http.StatusInternalServerError,
		})
	}
	var req UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Bad request",
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
			"error":  "Bad request",
			"status": http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})
	return
}

func (h *Handler) Find(c *gin.Context) {
	claims, _ := c.Get("User")
	logrus.Info(claims)
	var req model.FindUserRequest
	if c.Bind(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Bad Request",
		})
		return
	}

	// Nếu k truyền page, perPage lên thì phải gán giá trị mặc định
	if req.Page == 0 {
		req.Page = constants.FIRST_PAGE
	}
	if req.PerPage == 0 {
		req.PerPage = constants.PER_PAGE
	}

	users, total, err := h.UserRepo.Find(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"request": req,
			"error":   err,
		}).Error("failed when find users")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"data":  users,
			"total": total,
		},
	})
}
