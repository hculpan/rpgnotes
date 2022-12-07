package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"Title": "RPG Notes",
	})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.gohtml", gin.H{})
}
