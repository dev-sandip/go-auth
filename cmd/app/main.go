package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/styles"
	"github.com/dev-sandip/go-auth/cmd/db"
	model "github.com/dev-sandip/go-auth/cmd/models"
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
	db.InitMongoCollections(client)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Static("/static", "./static")
	// -----Views Routes-----
	router.GET("/", RenderHomePage)
	router.GET("/login", RenderLoginPage)
	router.GET("/register", RenderRegisterPage)
	router.GET("/home", RenderIndexPage)
	// markdown style serving
	router.GET("/css/syntax-highlighting.css", func(c *gin.Context) {
		c.Header("Content-Type", "text/css")
		style := styles.Get("dracula")
		formatter := html.New(html.WithLineNumbers(true))
		formatter.WriteCSS(c.Writer, style)
	})

	// Serve static files
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
	router.POST("/register", func(c *gin.Context) {
		var user model.User

		// First, bind the JSON body to the user struct
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON input"})
			return
		}

		// Validate required fields
		if user.Name == "" || user.Email == "" || user.Password == "" {
			c.JSON(400, gin.H{"error": "All fields (name, email, password) are required"})
			return
		}

		// Insert user into the DB
		_, err = db.CreateUser(&user)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user"})
			return
		}

		// Respond with success
		c.JSON(201, gin.H{"message": "User registered successfully"})
	})

	log.Println("Server starting on :8080...")
	router.Run(":8080")
}
