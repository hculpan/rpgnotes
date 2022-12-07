package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hculpan/rpgnotes/config"
	"github.com/hculpan/rpgnotes/models"
)

var lastUser models.User = models.User{}

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("RPGNOTES_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		if lastUser.ID == claims["sub"] {
			user = lastUser
		} else {
			result := config.DB.First(&user, claims["sub"])
			if result.Error != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			lastUser = user
		}

		c.Set("user", user)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
