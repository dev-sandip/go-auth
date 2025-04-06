package main

import (
	"fmt"
	"log"

	"github.com/dev-sandip/go-auth/templates/pages"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Static("/static", "./static")
	// -----Views Routes-----
	router.GET("/", RenderHomePage)
	router.GET("/login", RenderLoginPage)
	router.GET("/register", RenderRegisterPage)

	// -----API Routes-----
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
