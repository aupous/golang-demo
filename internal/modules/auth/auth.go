package auth

import (
	"awesomeProject/internal/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Role  string `json:"role"` // Với system có role
	ID    string `json:"id"`
}

func CreateToken(user *model.User, secret string) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
		Email: user.Email,
		Role:  "member",
		ID:    user.ID.String(),
		//
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
