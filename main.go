package main

import (
	"fmt"
	"log"

	"github.com/a-h/templ"
	"github.com/dev-sandip/go-auth/templates/pages"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	return template.Render(c.Request.Context(), c.Writer)
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		component := pages.Index()
		render(c, 200, component)
	})

	router.GET("/login", func(c *gin.Context) {
		component := pages.Login()
		render(c, 200, component)
	})
	router.GET("/register", func(c *gin.Context) {
		component := pages.Register()
		render(c, 200, component)
	})
	router.POST("/login", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		if email != "" && password != "" {
			fmt.Println("Login successful", email, password)
			c.Redirect(302, "/")
			return
		}

		component := pages.Login()
		render(c, 200, component)
	})

	log.Println("Server starting on :8080...")
	router.Run(":8080")
}
