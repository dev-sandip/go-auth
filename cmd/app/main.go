package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dev-sandip/go-auth/templates/pages"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env loading
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	mongoURI := os.Getenv("MONGO_URI")

	fmt.Println("Mongo URI", mongoURI)
	client, err := ConnectMongo(mongoURI)
	if err != nil {
		log.Fatal("❌ Error connecting to MongoDB", err)
	}
	defer DisconnectMongo(client)
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
