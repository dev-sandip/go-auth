package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dev-sandip/go-auth/templates/pages"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		component := pages.Index()
		component.Render(context.Background(), c.Writer)
	})

	router.GET("/login", func(c *gin.Context) {
		component := pages.Login()
		component.Render(context.Background(), c.Writer)
	})

	router.POST("/login", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		if email != "" && password != "" {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}

		component := pages.Login()
		component.Render(context.Background(), c.Writer)
	})

	log.Println("Server starting on :8080...")
	router.Run(":8080")
}
