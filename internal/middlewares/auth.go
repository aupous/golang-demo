package middlewares

import (
	"awesomeProject/internal/configs"
	responsehelper "awesomeProject/internal/helpers"
	"awesomeProject/internal/modules/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuth(config *configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenString == "" {
			responsehelper.ResponseWithUnauthorized(c, "Unauthorized")
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &auth.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})

		if err != nil {
			responsehelper.ResponseWithUnauthorized(c, "Token is not valid")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*auth.CustomClaims); ok && token.Valid {
			c.Set("User", claims)
			c.Next()
			return
		} else {
			responsehelper.ResponseWithUnauthorized(c, "Cannot parse token data")
			c.Abort()
			return
		}
	}
}
