package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hculpan/rpgnotes/config"
	"github.com/hculpan/rpgnotes/controllers"
	"github.com/hculpan/rpgnotes/middleware"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
	config.LoadCateogires()
}

func main() {
	r := gin.Default()

	r.POST("/api/notes", middleware.RequireAuth, controllers.NoteCreate)
	r.PUT("/api/notes/:id", middleware.RequireAuth, controllers.NoteUpdate)

	r.GET("/api/notes", middleware.RequireAuth, controllers.NotesIndex)
	r.GET("/api/notes/:id", middleware.RequireAuth, controllers.NoteShow)

	r.DELETE("/api/notes/:id", middleware.RequireAuth, controllers.NoteDelete)

	r.POST("/api/signup", controllers.Signup)
	r.POST("/api/login", controllers.Login)
	r.POST("/api/logout", controllers.Logout)

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.gohtml")

	r.GET("/", middleware.RequireWebAuth, controllers.RootPage)
	r.GET("/login", controllers.LoginPage)

	r.Run()
}
