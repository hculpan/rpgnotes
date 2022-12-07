package middleware

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hculpan/rpgnotes/config"
	"github.com/hculpan/rpgnotes/models"
)

func RequireWebAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	location := url.URL{Path: "/login"}

	if err != nil {
		c.Redirect(http.StatusFound, location.RequestURI())
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("RPGNOTES_SECRET")), nil
	})
	if err != nil {
		c.Redirect(http.StatusFound, location.RequestURI())
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}

		var user models.User
		if lastUser.ID == claims["sub"] {
			user = lastUser
		} else {
			result := config.DB.First(&user, claims["sub"])
			if result.Error != nil {
				c.Redirect(http.StatusFound, location.RequestURI())
				return
			}
			lastUser = user
		}

		c.Set("user", user)
	} else {
		c.Redirect(http.StatusFound, location.RequestURI())
		return
	}

	c.Next()
}
