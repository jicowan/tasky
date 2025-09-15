package main

import (
	"net/http"

	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	controller "github.com/jeffthorne/tasky/controllers"
	"github.com/jeffthorne/tasky/middleware"
	"github.com/joho/godotenv"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func main() {
	godotenv.Overload()

	router := gin.Default()

	// Security middleware
	router.Use(secure.New(secure.Config{
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; font-src 'self' https://fonts.gstatic.com;",
		BrowserXssFilter:      true,
		ContentTypeNosniff:    true,
		SSLRedirect:           true,
		SSLHost:               "localhost:8080",
		STSSeconds:            31536000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
	}))

	router.LoadHTMLGlob("assets/*.html")
	router.Static("/assets", "./assets")

	// Public routes
	router.GET("/", index)
	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)

	// Protected routes
	protected := router.Group("")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/todos/:userid", controller.GetTodos)
		protected.GET("/todo/:id", controller.GetTodo)
		protected.POST("/todo/:userid", controller.AddTodo)
		protected.DELETE("/todo/:userid/:id", controller.DeleteTodo)
		protected.DELETE("/todos/:userid", controller.ClearAll)
		protected.PUT("/todo", controller.UpdateTodo)
		protected.GET("/todo", controller.Todo)
	}

	router.Run(":8080")
}
